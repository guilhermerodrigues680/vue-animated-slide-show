<template>
  <div id="app">
    <div class="navigation">
      <span @click="previousPage()">&lt;</span>
      <span @click="nextPage()">&gt;</span>
    </div>
    <Interpolation v-if="currentPage === 0" />
    <GradientBackground v-if="currentPage === 1" />
    <DownloadArrow v-if="currentPage === 2" />
    <TextReveal v-if="currentPage === 3" />
  </div>
</template>

<script>
import Interpolation from '@/views/Interpolation.vue'
import GradientBackground from '@/views/GradientBackground.vue'
import DownloadArrow from '@/views/DownloadArrow.vue'
import TextReveal from '@/views/TextReveal.vue'

const totalPages = 4;

export default {
  name: 'App',

  components: {
    Interpolation,
    GradientBackground,
    DownloadArrow,
    TextReveal,
  },

  data: () => ({
    currentPage: 0,
  }),

  methods: {
    nextPage() {
      this.currentPage = this.currentPage + 1 < totalPages ? this.currentPage + 1 : 0;
    },

    previousPage() {
      this.currentPage = this.currentPage - 1 >= 0 ? this.currentPage - 1 : totalPages-1;
    }
  }
}
</script>

<style>
/* #app {} */
.navigation {
  display: flex;
  top: 10px;
  left: 10px;
  background: #FFFFFFAA;
  border-radius: 0.5rem;
  position: fixed;
  user-select: none;
}

.navigation span {
  cursor: pointer;
  font-size: 1.5rem;
  font-weight: bold;
  padding: 0.25rem 0.25rem;
}

.navigation span:hover {
  padding: 0 0.25rem;
  font-size: 1.75rem;
}

.navigation span:active {
  color: deeppink;
}
</style>
