<template>
  <div class="container">
    <div class="header">
      <div class="nav-blur"></div>
      <div class="nav-bg"></div>
    </div>
    <div class="main">

    </div>
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
  background: #F7F7F7 url("https://api.bimg.cc/random?w=1920&h=1080&mkt=zh-CN") no-repeat fixed;
  background-size: 100%;
}

.nav-blur {
  mask-image: linear-gradient(black 40%, transparent);
  backdrop-filter: blur(20px) saturate(180%);
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 70px;
}

.nav-bg {
  background: linear-gradient(to bottom, rgba(255, 255, 255, 25%), transparent);
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 70px;
  opacity: 70%;
}
.main{
  background-color: rgba(0,0,0,0.2);
  width: 80vw;
  margin: 70px 0 0 0;
}
</style>
