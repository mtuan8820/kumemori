<script setup lang="ts">
import { useI18n } from "vue-i18n";

import DeckList from "@/components/DeckList.vue";

const { t } = useI18n();
</script>

<template>
  <div class="home">
    <DeckList/>
  </div>
</template>


