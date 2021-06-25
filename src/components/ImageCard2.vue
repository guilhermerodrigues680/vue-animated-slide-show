<template>
  <div class="wrapper">
    <img class="photo" :src="imgurl" alt="foto">
    <h3>{{title}}</h3>
    <p>{{text}}</p>

    <svg class="mask">
      <clipPath id="60da142bimagecard2mask0" clipPathUnits="objectBoundingBox">
        <path v-if="mask === 'heart'" d="M0.852,0.031 c-0.13,-0.066,-0.295,0,-0.352,0.141 C0.442,0.023,0.27,-0.044,0.148,0.031 C0.011,0.106,-0.054,0.288,0.054,0.511 C0.133,0.669,0.27,0.785,0.5,1 c0.237,-0.215,0.374,-0.331,0.446,-0.489 C1,0.296,0.989,0.106,0.852,0.031" fill="#C4C4C4" />
        <path v-if="mask === 'star'" d="M0.5,0.846 l-0.309,0.154,0.044,-0.357 L0,0.382 l0.336,-0.067 L0.5,0 l0.164,0.315,0.336,0.067,-0.235,0.261,0.044,0.357" fill="#C4C4C4" />
      </clipPath>
    </svg>
  </div>
</template>

<script>
export default {
  name: 'ImageCard2',

  props: {
    imgurl: String,
    title: String,
    text: String,
  },

  data: () => ({
    mask: 'star',
    // mask: 'heart',
  }),
}
</script>

<style scoped>
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

.photo {
  height: 200px;
  width: 200px;
  object-fit: cover;
  clip-path: url(#60da142bimagecard2mask0);
}

.mask {
  width: 0px;
  height: 0px;
  position: absolute;
}
</style>
