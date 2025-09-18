<script setup lang="ts">
    import {useDeckDetailViewModel} from '@/viewmodels/useDeckDetailViewModel'
    import { onMounted, ref } from 'vue';
    import CardSlider from '@/components/CardSlider.vue'
    import { PencilIcon } from '@heroicons/vue/24/outline';

    interface DeckViewProps{
        id: string
        name?: string
    }

    const props = defineProps<DeckViewProps>()

    const deckId = Number(props.id)
    const {cards, loadCards, createCard} = useDeckDetailViewModel(deckId)
    
    onMounted(async () => {
        await loadCards()
    })

    const handleCreateCard = async () => {
        await createCard()
        await loadCards()
    }

</script>

<template>
    <div class="flex items-center justify-between mb-3">

        <h2>{{ props.name }}</h2>
        <button class="flex items-center gap-1" @click="">
            <PencilIcon class="h-4 w-4"/> <span>Edit</span>
        </button>
    </div>
 
    <CardSlider :cards="cards" />
</template>

<style scoped>
    .card{
        @apply rounded-md border-border border w-full h-full
    }
</style>