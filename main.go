package main

import (
	"context"
	"embed"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	// o conteúdo contém o conteúdo do nosso servidor da web estático.
	//go:embed webapp/dist/*
	embeddedContent embed.FS
	listen          = flag.String("listen", ":8000", "listen address")
	logDebug        *log.Logger
	logWebServer    *log.Logger
	logInfo         *log.Logger
	logWarn         *log.Logger
	logError        *log.Logger
)

type logServer struct {
	hdl http.Handler
}

func (l *logServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lw := NewLoggingResponseWriter(w)
	l.hdl.ServeHTTP(lw, r)
	remoteIp, _, _ := net.SplitHostPort(r.RemoteAddr)
	logWebServer.Printf("%s '%s %s' %d", remoteIp, r.Method, r.URL, lw.statusCode)
}

// loggingResponseWriter faz o log do código HTTP da resposta
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	// WriteHeader(int) não é chamado se nossa resposta retornar implicitamente 200 OK, então
	// configura-se por default este status code.
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lw *loggingResponseWriter) WriteHeader(code int) {
	lw.statusCode = code
	lw.ResponseWriter.WriteHeader(code)
}

func normalizeUrl(url string) string {
	normalizedUrl := url
	if strings.HasPrefix(normalizedUrl, ":") {
		normalizedUrl = "http://localhost" + normalizedUrl
	}

	if !strings.HasPrefix(normalizedUrl, "http") {
		normalizedUrl = "http://" + normalizedUrl
	}

	return normalizedUrl
}

func openbrowser(url string) error {
	err := error(nil)
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}

func init() {
	flag.Parse()

	log.SetOutput(os.Stdout)
	log.SetPrefix("\033[34m[INFO]  \033[0m")
	logDebug = log.New(os.Stdout, "\033[35m[DEBUG] \033[0m", log.Ldate|log.Ltime|log.Lshortfile)
	logWebServer = log.New(os.Stdout, "\033[36m[WEB]   \033[0m", log.Ldate|log.Ltime)
	logInfo = log.Default()
	logWarn = log.New(os.Stdout, "\033[33m[WARN]  \033[0m", log.LstdFlags)
	logError = log.New(os.Stderr, "\033[31m[ERROR] \033[0m", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	log.Println("Iniciando execução do servidor web. PID:", os.Getpid())
	webappDist, err := fs.Sub(embeddedContent, "webapp/dist")
	if err != nil {
		logError.Fatalln(err)
	}

	httpServerExitDone := &sync.WaitGroup{}
	httpServerExitDone.Add(1)

	http.Handle("/", http.FileServer(http.FS(webappDist)))
	http.HandleFunc("/__/api/v1/teste-shutdown", func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		rw.WriteHeader(http.StatusOK)
	})
	server := &http.Server{Addr: *listen, Handler: &logServer{http.DefaultServeMux}}
	serverErrChan := make(chan error)
	go func() {
		defer httpServerExitDone.Done()
		log.Println("Servidor web escutando o endereco:", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				close(serverErrChan)
			} else {
				logError.Println("Serve Error:", err)
				serverErrChan <- err
			}
		}
	}()

	normalizedUrl := normalizeUrl(*listen)
	if err := openbrowser(normalizedUrl); err == nil {
		log.Println("Aberto no browser a página", normalizedUrl)
	} else {
		logWarn.Println("Não foi possivel abrir o browser. Acesse manualmente a página", normalizedUrl)
	}

	notifyContext, cancel := signal.NotifyContext(context.Background(), syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	select {
	case <-notifyContext.Done():
		log.Println("O servidor recebeu um sinal para ser encerrado. Finalizando tarefas...")
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				logError.Fatalln("TIMEOUT! Encerrada execução do servidor web por timeout. PID:", os.Getpid())
			} else {
				logError.Fatalln(err)
			}
		}
		httpServerExitDone.Wait()
	case err := <-serverErrChan:
		if err != nil {
			logError.Fatalln(err)
		}
	}

	log.Println("OK! Encerrada execução do servidor web. PID:", os.Getpid())
}
