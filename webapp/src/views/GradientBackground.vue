<template>
  <div class="container">
    <h1>Animação Degradê</h1>
  </div>
</template>

<script>
export default {
  name: 'GradientBackground',
}
</script>

<style scoped>
.container {
  animation: gradient 10s ease-in-out infinite alternate;
  background: linear-gradient(135deg, #23a6d5, #23d5ab, #ee7752, #e73c7e) no-repeat;
  background-size: 300% 300%;
  background-position: 0 0;
  display: grid;
  min-height: 100vh;
  place-content: center;
}

h1 {
  color: white;
  font-family: Arial, Helvetica, sans-serif;
  font-size: 4rem;
  text-align: center;
}

@keyframes gradient {
  from {
    background-position: 0 50%;
  }
  to {
    background-position: 100% 50%;
  }
}
</style>
