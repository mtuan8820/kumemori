<script setup lang="ts">
import { useI18n } from "vue-i18n";
import {SunIcon, MoonIcon} from '@heroicons/vue/24/outline'
import { onMounted, ref } from "vue";

const { t, availableLocales: languages, locale } = useI18n();

const onclickLanguageHandle = (item: string) => {
  item !== locale.value ? (locale.value = item) : false;
};

const isDark = ref(false)

const toggleThemeHandle = () => {
    document.body.classList.toggle("dark")
    isDark.value = !isDark.value
}


const onclickMinimise = () => {};

const onclickQuit = () => {};

onMounted(() => {
  isDark.value = document.body.classList.contains("dark")
})

</script>
<template>
<div class="header">
    <h1 class="content-center">Kumemori</h1>
    <!-- Menu -->
    <div class="menu flex items-center">
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
        <SunIcon v-if="!isDark" class="w-6 h-6 text-black" />
        <MoonIcon v-else class="w-6 h-6 text-white" />

      </button>
    </div>
</div>

</template>

<style >
    .header{
        @apply h-full w-full px-4 
            flex content-center justify-between
            bg-white border-b border-gray-900/10 
            dark:bg-black 
    }       
</style>