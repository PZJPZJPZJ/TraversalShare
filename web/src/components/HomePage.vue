<template>
  <div class="container">

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
  align-items: center;
  min-height: 100vh;
  background-color: #F7F7F7;
}
</style>
