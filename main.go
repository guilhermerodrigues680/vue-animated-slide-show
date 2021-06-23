package main

import (
	"embed"
	"flag"
	"fmt"
)

var (
	// o conteúdo contém o conteúdo do nosso servidor da web estático.
	//go:embed webapp/dist/*
	content embed.FS
	listen  = flag.String("listen", ":8000", "listen address")
)

func main() {
	flag.Parse()
	fmt.Println("")
}
