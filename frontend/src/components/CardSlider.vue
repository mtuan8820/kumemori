<script setup lang="ts">
import {  ref } from 'vue';
import emblaCarouselVue from 'embla-carousel-vue';
import { ArrowRightIcon, ArrowLeftIcon } from '@heroicons/vue/24/outline';
import type { Card } from '@/models/Card';

interface CardSliderProps{
  cards: Card[]
}

const props = defineProps<CardSliderProps>()

const [emblaRef, emblaApi] = emblaCarouselVue({ loop: true, container: '.embla__container' }, );

const showFront = ref(true)

const prevSlide = () =>   emblaApi.value?.scrollPrev();
const nextSlide = () =>   emblaApi.value?.scrollNext();
const toggleFrontBack = () => showFront.value = !showFront.value

</script>

<template>
  <div class="embla" ref="emblaRef">
  <div class="embla__viewport">
      <div class="embla__container">

        <div class="embla__slide" v-for="card in props.cards" :key="card.id">
          <div class="embla__parallax">
            <div class="embla__parallax__layer">
              <div class="embla__card" @click="toggleFrontBack">
                <div v-if="showFront">{{ card.front }}</div>
                <div v-else>{{ card.back }}</div>
              </div>
            </div>
          </div>
        </div>

      </div>
    </div>
    <div class="flex justify-center gap-3 mt-3">
      <button @click="prevSlide" class="slider__button"><ArrowLeftIcon class="h-4 w-4"/></button>
      <button @click="nextSlide" class="slider__button"><ArrowRightIcon class="h-4 w-4"/></button>
    </div>
  </div>
</template>

<style>
.embla {
  max-width: 48rem;
  margin: auto;
  --slide-height: 19rem;
  --slide-spacing: 1rem;
  --slide-size: 100%;
}

.embla__viewport {
  overflow: hidden;
}

.embla__container {
  display: flex;
  touch-action: pan-y pinch-zoom;
  margin-left: calc(var(--slide-spacing) * -1);
}

.embla__slide {
  transform: translate3d(0, 0, 0);
  flex: 0 0 var(--slide-size);
  min-width: 0;
  padding-left: var(--slide-spacing);
}

.embla__parallax {
  @apply rounded-2xl h-full overflow-hidden
}
.embla__parallax__layer {
  @apply h-full w-full
    relative flex justify-center
}

.embla__card{
  @apply w-full  p-4 text-center h-[var(--slide-height)] 
  border border-border rounded-2xl
  bg-dark
}

.slider__button{
  @apply border border-border rounded-md px-4 py-2
}

</style>