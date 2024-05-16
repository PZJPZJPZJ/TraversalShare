<template>
  <div class="container">
    <div class="app-left">
      <div class="menu-content">

      </div>
    </div>
    <div class="app-middle">
      <div class="list-content">

      </div>
    </div>
    <div class="app-right">
      <div class="chat-content">

      </div>
    </div>
  </div>
</template>

<script setup>

</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: rgb(39, 42, 55);
}
.app-left{
  background-color: rgba(0,0,0,0.2);
  width: 10vw;
  height: 100vh;
}
.app-middle{
  background-color: rgba(0,0,0,0.3);
  width: 20vw;
  height: 100vh;
}
.app-right{
  background-color: rgba(0,0,0,0.2);
  width: 70vw;
  height: 100vh;
}
.list-content{
  background-color: rgba(255,255,255,0.2);
  align-items: center;
  justify-content: center;
  height: 20px;
}
</style>