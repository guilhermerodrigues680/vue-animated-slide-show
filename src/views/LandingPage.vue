<template>
  <div class="page-container">
    <transition
      v-if="hasStarted"
      appear
      name="custom-classes-transition"
      enter-active-class="animate__animated animate__fadeInDown"
      :leave-active-class="transitionLeaveActiveClass"
      mode="out-in"
      @after-enter="afterEnter"
      @after-appear="afterEnter"
    >
      <ImageCard class="image-card"
        :imgurl="imageList[currentImage].imgUrl"
        :title="imageList[currentImage].title"
        :text="imageList[currentImage].text"
        :key="currentImage"
      />
    </transition>
    <transition
      v-else
      appear
      name="custom-classes-transition"
      enter-active-class="animate__animated animate__fadeInDownBig"
      mode="out-in"
    >
      <div  class="start-menu-container">
        <button class="start-button" @click="start()" :disabled="startButtonIsDisabled">Estou!</button>
      </div>
    </transition>

    <!-- Audio da página -->
    <audio ref="audio">
      <source :src="require('@/assets/ron-gelinas-chillout-lounge-where-will-i-go.min.mp3')" type="audio/mpeg">
      Your browser does not support the audio element.
    </audio>
  </div>
</template>

<script>
import ImageCard from "@/components/ImageCard.vue";
import Vue from 'vue';

export default {
  name: 'LandingPage',

  components: {
    ImageCard,
  },

  data: () => ({
    imageList: [
      { imgUrl: require('@/assets/demo-img-1.jpg'), title: 'demo-img-1', text: 'Imagem demo 1' },
      { imgUrl: require('@/assets/demo-img-2.jpg'), title: 'demo-img-2', text: 'Imagem demo 2' },
      { imgUrl: require('@/assets/demo-img-3.jpg'), title: 'demo-img-3', text: 'Imagem demo 3' },
      { imgUrl: require('@/assets/demo-img-4.jpg'), title: 'demo-img-4', text: 'Imagem demo 4' },
    ],
    currentImage: 0,
    intervalMs: 1000,
    transitionLeaveActiveClass: "animate__animated animate__fadeOut",
    hasStarted: false,
    startButtonIsDisabled: false,
  }),

  watch: {
    currentImage() {
      const transitions = [
        // "animate__fadeOut",
        "animate__fadeOutUp",
        "animate__fadeOutRight",
        "animate__fadeOutLeft",
        "animate__fadeOutDown",
        "animate__fadeOutTopLeft",
        "animate__fadeOutTopRight",
        "animate__fadeOutBottomRight",
        "animate__fadeOutBottomLeft",
      ];

      // intervalo [min,max) , `min` (inclusivo) e `max` (exclusivo)
      const randomParams = { max: transitions.length, min: 0 };
      const rIdx = Math.floor(Math.random() * (randomParams.max - randomParams.min)) + randomParams.min;
      this.transitionLeaveActiveClass =  `animate__animated ${transitions[rIdx]}`;
    },
  },

  methods: {
    async start() {
      /** @type{HTMLAudioElement}*/
      const audio = this.$refs.audio;

      try {
        await audio.play();
      } catch (error) {
        console.error(error);
        alert(error.message);
        return
      }

      // startButtonIsDisabled e Vue.nextTick, garantem que o botao seja desabilitado.
      this.startButtonIsDisabled = true;
      await Vue.nextTick();
      this.hasStarted = true;
    },
    afterEnter(/*el*/) {
      setTimeout(() => this.nextImage(), this.intervalMs);
    },
    nextImage() {
      const imageListLastIdx = this.imageList.length - 1;
      this.currentImage = this.currentImage + 1 > imageListLastIdx ? 0 : this.currentImage + 1;
    },
    previousImage() {
      const imageListLastIdx = this.imageList.length - 1;
      this.currentImage = this.currentImage - 1 < 0 ? imageListLastIdx : this.currentImage - 1;
    },
  }
}
</script>

<style scoped>
.page-container {
  display: grid;
  grid-template-rows: 100vh;
  place-items: center;
  background-color: #111111;
  background-image: url('~@/assets/night-sky-texture.jpg');
  /* background-image: linear-gradient(to top, #e0c3fc40, #8ec5fc40), url('~@/assets/night-sky-texture.jpg'); */
  color: white;
  font-family: sans-serif;
  position: relative;
  animation: night-sky-moving 30s linear infinite;
}

.page-container::before {
  content: "";
  z-index: 1;
  display: block;
  position: absolute;
  width: 100%;
  height: 100%;
  background: linear-gradient(to top, #e0c3fc30, #8ec5fc30);
  /* background: linear-gradient(to top, var(--from, #e0c3fc30), var(--to, #8ec5fc30)); */
}

.image-card,
.start-menu-container {
  z-index: 10;
}

.start-button {
  user-select: none;
  cursor: pointer;
  margin: 0;
  border: none;
  font-size: 2rem;
  padding: 1rem 2rem;
  border-radius: 0.5rem;
  color: white;
  font-weight: bold;
  /* background-color: #4158D010; */
  background-color: transparent;
  background-image: linear-gradient(43deg, #4159d080 0%, #C850C080 46%, #ffcd70c2 100%);
}

.start-button:hover {
  background-image: linear-gradient(43deg, #4159d0d0 0%, #C850C0d0 46%, #ffcd70e8 100%);
}

.start-button:focus,
.start-button:active {
  outline: 2px solid white;
}

.start-button:disabled {
  cursor: not-allowed;
}

@keyframes night-sky-moving {
  to {
    /* dimensões da textura */
    background-position: 600px 600px;
  }
}

/* ============ Vue Transition ============ */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 1.5s;
}

.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>
