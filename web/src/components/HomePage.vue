<template>
  <div class="container">
    <div class="navbar"></div>
  </div>
</template>

<script setup>
import {onMounted} from "vue";

const queryParams = new URLSearchParams(window.location.search);
onMounted(() => {
  queryParams.get('test')
})
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  min-height: 100vh;
  background: #F7F7F7 url("https://t.mwm.moe/pc/") no-repeat fixed;
  background-size: 100%;
}

.navbar {
  background: linear-gradient(to bottom, rgba(255, 255, 255, 0.3), transparent);
  mask-image: linear-gradient(black 40%, transparent);
  backdrop-filter: blur(20px) saturate(180%);
  position: sticky;
  top: 0;
  width: 100vw;
  height: 70px;
}
</style>
