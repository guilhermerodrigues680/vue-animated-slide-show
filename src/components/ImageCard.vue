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
  grid-template-columns: 1fr;
  grid-template-rows: 4fr 1fr 1fr;
  place-items: center;
  height: 40%;
  width: 50%;
  position: relative;
  user-select: none;
}

/* disable select */
.wrapper::before {
  content: "";
  display: block;
  z-index: 1;
  position: absolute;
  width: 100%;
  height: 100%;
}

.photo {
  height: 100%;
  width: 100%;
  object-fit: contain;
  filter: drop-shadow(0px 0px 1px rgb(224, 195, 252, 0.3)) drop-shadow(0px 0px 5px rgb(224, 195, 252, 0.3));
}
</style>
