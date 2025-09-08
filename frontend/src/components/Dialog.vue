<script setup lang="ts">
  import { watch, ref } from 'vue';

  type Props = {
    isOpen: boolean;
  };

  const props = defineProps<Props>();

  const emit = defineEmits(['close']);

  const dialogRef = ref();

  const handleClose = () => {
    emit('close');
  }

  watch(
    () => props.isOpen,
    newValue => {
      if (!dialogRef.value) return;
      if (newValue) {
        dialogRef.value.showModal();
      } else {
        dialogRef.value.close();
      }
    }
  );
</script>

<template>
    <dialog ref="dialogRef">
        <slot/>
        <button @click="handleClose"> Close </button>
    </dialog>
</template>