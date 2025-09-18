<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'

const props = defineProps<{
  modelValue: string
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
}>()

const textarea = ref<HTMLTextAreaElement | null>(null)

const resize = () => {
  const el = textarea.value
  if (!el) return
  el.style.height = 'auto'
  el.style.height = el.scrollHeight + 'px'
}

// keep resizing when value changes
watch(() => props.modelValue, resize)

onMounted(resize)

const onInput = (e: Event) => {
  const value = (e.target as HTMLTextAreaElement).value
  emit('update:modelValue', value)
  resize()
}
</script>

<template>
  <textarea
    ref="textarea"
    :value="modelValue"
    @input="onInput"
    rows="1"
    class="w-full border rounded px-2 py-1 resize-none overflow-hidden"
  ></textarea>
</template>