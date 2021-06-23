<template>
  <div class="interpolation">
    <div class="plane"></div>
  </div>
</template>

<script>
export default {
  name: 'Interpolation',
}
</script>

<style scoped>
.interpolation {
  display: grid;
  align-items: center;
  min-height: 100vh;
  background-image: url('~@/assets/bg-game.jpg');
  background-repeat: no-repeat;
  background-size: cover;
  overflow: hidden;
}

.plane {
  width: 100%;
  height: 80px;
  animation: fly 2s ease-in-out infinite;
  background-image: url('~@/assets/plane-game.png');
  background-size: 120px 80px;
  background-repeat: no-repeat;
}

@keyframes fly {
  0% {
    transform: translate(-120px,0);
  }

  40% {
    transform: translate(calc(40% - (120px / 2)), 100px);
  }

  60% {
    transform: translate(calc(60% - (120px / 2)), -100px);
  }

  100% {
    transform: translate(calc(100%), 0);
  }
}
</style>
