<template>
  <div id="app">
    <Interpolation v-if="currentPage === 0" />
    <GradientBackground v-if="currentPage === 1" />
  </div>
</template>

<script>
import Interpolation from '@/views/Interpolation.vue'
import GradientBackground from '@/views/GradientBackground.vue'

export default {
  name: 'App',

  components: {
    Interpolation,
    GradientBackground,
  },

  data: () => ({
    currentPage: 1,
  }),
}
</script>

<style>
#app {

}
</style>
