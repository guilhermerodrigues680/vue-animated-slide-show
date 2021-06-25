<template>
  <div class="page-container">
    <ImageCard
      :imgurl="imageList[currentImage].imgUrl"
      :title="imageList[currentImage].title"
      :text="imageList[currentImage].text"
    />
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
  }),
}
</script>

<style scoped>
.page-container {
  display: grid;
  grid-template-rows: 100vh;
  place-items: center;
  background: #157bad url('~@/assets/night-sky.jpg') no-repeat;
  background-size: cover;
  color: white;
  font-family: sans-serif;
}
</style>
