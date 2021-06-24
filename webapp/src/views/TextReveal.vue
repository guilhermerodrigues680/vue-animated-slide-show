<template>
  <div class="container">
    <h1>Revelação de texto</h1>
  </div>
</template>

<script>
export default {
  name: 'TextReveal',
}
</script>

<style scoped>
.container {
  display: grid;
  place-items: center;
  min-height: 100vh;
  font-family: Arial, Helvetica, sans-serif;
  background-color: #FBAB7E;
  background-image: linear-gradient(62deg, #FBAB7E 0%, #F7CE68 100%);
}

h1 {
  position: relative;
  color: #444444;
  animation: has-reveal-text 750ms linear;
}

h1::after {
  content: "";
  display: block;
  background-color: white;
  position: absolute;
  top: 0;
  right: 0;
  width: 100%;
  height: 100%;
  transform-origin: center right;
  animation: has-reveal 750ms linear forwards;
}

@keyframes has-reveal {
  to {
    transform: scaleX(0);
  }
}

@keyframes has-reveal-text {
  from {
    color: transparent;
  }
  to {
    color: transparent;
  }
}
</style>
