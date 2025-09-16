<script setup lang="ts">
import { useI18n } from "vue-i18n";
import {SunIcon, MoonIcon} from '@heroicons/vue/24/outline'
import { onMounted, ref } from "vue";

import { OpenURL } from "../../wailsjs/go/main/App"  

// const { t, availableLocales: languages, locale } = useI18n();

// const onclickLanguageHandle = (item: string) => {
//   item !== locale.value ? (locale.value = item) : false;
// };

const isDark = ref(false)

const toggleThemeHandle = () => {
    document.body.classList.toggle("dark")
    isDark.value = !isDark.value
}

onMounted(() => {
  isDark.value = document.body.classList.contains("dark")
})


function openGithub() {
  OpenURL("https://github.com/mtuan8820/kumemori")
}

</script>
<template>
<div class="header ">
    <h1 class="content-center">Kumemori</h1>
    <!-- Menu -->
    <div class="menu flex items-center gap-3">
      <!-- <div class="language">
        <div
          v-for="item in languages"
          :key="item"
          :class="{ active: item === locale }"
          @click="onclickLanguageHandle(item)"
          class="lang-item"
        >
          {{ t("languages." + item) }}
        </div>
      </div> -->
      <!-- <div class="bar">
        <div class="bar-btn" @click="onclickMinimise">
          {{ t("topbar.minimise") }}
        </div>
        <div class="bar-btn" @click="onclickQuit">{{ t("topbar.quit") }}</div>
      </div> -->
    
      <button @click="toggleThemeHandle()">
        <SunIcon v-if="!isDark" class="w-6 h-6 text-text" />
        <MoonIcon v-else class="w-6 h-6 text-text" />
      </button>
      <button @click="openGithub">

        <img src="/github-mark.svg" alt="Github icon" class="w-6 h-6 text-main" />
      </button>
    </div>
</div>

</template>

<style>
  .header{
    @apply 
        h-full w-full px-4 bg-light
        flex content-center justify-between
        border-b border-border
  }
</style>