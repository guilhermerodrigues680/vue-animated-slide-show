<template>
  <div class="wrapper">
    <img class="photo" :src="imgurl" alt="foto">
    <h3>{{title}}</h3>
    <p>{{text}}</p>
  </div>
</template>

<script>
export default {
  name: 'ImageCard',

  props: {
    imgurl: String,
    title: String,
    text: String,
  },
}
</script>

<style scoped>
h3, p {
  margin: 0;
}

.wrapper {
  display: grid;
  grid-template-columns: repeat(1, minmax(0, 1fr));
  grid-template-rows: 4fr 1fr 1fr;
  place-items: center;
  height: 40%;
  width: 50%;
}

.photo {
  height: 100%;
  width: auto;
  object-fit: contain;
  border-radius: 0.5rem;
  vertical-align: middle;
}
</style>
