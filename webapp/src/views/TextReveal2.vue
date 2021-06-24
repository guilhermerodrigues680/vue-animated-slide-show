<template>
  <div class="container">
    <h1>Revelação de texto 2</h1>
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
  animation: fade 1ms 500ms backwards;
  overflow: hidden;
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
  animation: has-reveal 1s cubic-bezier(0.77, 0, 0.175, 1) forwards;
}

@keyframes has-reveal {
  to {
    transform: translateX(100%);
  }
}

@keyframes fade {
  from {
    color: transparent;
  }
}
</style>
