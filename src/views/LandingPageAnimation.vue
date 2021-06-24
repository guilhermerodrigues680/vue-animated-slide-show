<template>
  <div class="landing-page-container">
    <div class="mountains"></div>
    <p class="has-shown">Landing Page</p>
    <h1 class="has-shown">Animações CSS</h1>
    <button class="button has-shown" style="animation-delay: 300ms;">Clique aqui</button>
  </div>
</template>

<script>
export default {
  name: 'LandingPageAnimation',
}
</script>

<style scoped>
.landing-page-container {
  display: grid;
  justify-items: center;
  align-content: center;
  font-family: Arial, Helvetica, sans-serif;
  position: relative;
  background-image: url('~@/assets/night-sky.jpg');
  background-size: cover;
  min-height: 100vh;
  animation: is-moving 4s ease-in-out infinite alternate;
}

@keyframes is-moving {
  to {
    background-position-x: 100%;
  }
}

.landing-page-container::before {
  content: "";
  display: block;
  position: absolute;
  z-index: 1;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: #8ec5fc10;
  background-image: linear-gradient(0deg, #e0c3fc50 0%, #8ec5fc10 80%);
}

.mountains {
  display: block;
  position: absolute;
  z-index: 10;
  bottom: 0;
  left: 0;
  height: 20vh;
  width: 100%;
  background-image: url('~@/assets/snow-cap-mountains.png');
  background-repeat: no-repeat;
  background-size: cover;
}

p, h1 {
  margin: 0;
  color: white;
  text-align: center;
}

p {
  font-size: 1.25rem;
}

h1 {
  font-size: 2.5rem;
}

.button {
  margin: 2rem 0;
  z-index: 10;
  padding: 0.75rem 1.75rem;
  font-size: 1.25rem;
  background-color: #746086;
  border: 0;
  color: white;
  border-radius: 0.125rem;
  text-transform: uppercase;
  cursor: pointer;
  position: relative;
}

.button::before {
  content: "";
  z-index: -1;
  display: block;
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  background-color: #746086;
  border-radius: 0.125rem;
  animation: is-rippling 1.6s cubic-bezier(0.23, 1, 0.320, 1) infinite;
}

.button:hover {
  background-color: #60506e;
}

.button:active {
  transform: scale(1.05);
}

.button:active::before {
  animation: none;
  display: none;
}

@keyframes is-rippling {
  to {
    opacity: 0;
    transform: scale(1.2, 1.6);
  }
}

.has-shown {
  animation: has-shown 1.2s cubic-bezier(0.645, 0.045, 0.355, 1) backwards;
}

@keyframes has-shown {
  from {
    opacity: 0;
    transform: translateY(50px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
</style>
