<template>
  <div class="container">
    <svg viewBox="4 4 16 16">
      <path fill="currentColor" d="M11,4H13V16L18.5,10.5L19.92,11.92L12,19.84L4.08,11.92L5.5,10.5L11,16V4Z" />
    </svg>
  </div>
</template>

<script>
export default {
  name: 'DownloadArrow',
}
</script>

<style scoped>
.container {
  display: grid;
  place-content: center;
  background-color: #D9AFD9;
  background-image: linear-gradient(0deg, #D9AFD9 0%, #97D9E1 100%);
  min-height: 100vh;
}

svg {
  height: 100px;
  color: #06517d;
  animation: is-bouncing 800ms cubic-bezier(0.68, -0.55, 0.265, 1.55) infinite alternate;
}

@keyframes is-bouncing {
  to {
    transform: translateY(10%);
  }
}
</style>
