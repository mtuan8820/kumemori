<script setup lang="ts">

import SideBar from "./components/SideBar.vue";
import Header from "./components/Header.vue";

document.body.addEventListener("click", function (event) {
  event.preventDefault();
});
</script>

<template>
  <div class="flex flex-col h-screen">
    <!-- Header -->
    <header class="h-[50px]">
        <Header />
    </header>
    
    <!-- Page -->
    <div class="grid grid-cols-10 flex-1 overflow-y-auto overflow-x-hidden">
      <!-- sidebar -->
      <SideBar class="col-span-3 bg-light "/>

      <!-- view -->
      <div class="col-span-7 bg-light overflow-y-auto">
        <router-view v-slot="{ Component }">
          <div class="p-4">
            <component :is="Component" />
          </div>
        </router-view>
      </div>
    </div>
  </div>

</template>

<style lang="scss">
@import url("./assets/css/reset.css");
@import url("./assets/css/font.css");

html {
  width: 100%;
  height: 100%;
}
body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  font-family: "JetBrainsMono";
  background-color: transparent;
}

#app {
  position: relative;
  // width: 900px;
  // height: 520px;
  height: 100%;
}
</style>
