<template>
  <div class="page-container">
    <transition
      appear
      name="fade"
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
  </div>
</template>

<script>
import ImageCard from "@/components/ImageCard.vue";

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
  }),

  methods: {
    afterEnter(/*el*/) {
      console.debug('afterEnter()');
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

.image-card {
  z-index: 10;
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
