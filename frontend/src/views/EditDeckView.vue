<script lang="ts" setup>

import AdaptiveTextarea from '@/components/AdaptiveTextarea.vue';
import { useEditDeckViewModel } from '@/viewmodels/useEditDeckViewModel';
import { TrashIcon } from '@heroicons/vue/24/outline';


interface EditDeckViewProps {
    id: string
    name?: string
}

const props = defineProps<EditDeckViewProps>()

const deckId = Number(props.id)

const { cardItems, createCardItem, deleteCardItem, submitUpdateDeck } = useEditDeckViewModel(deckId)



</script>
<template>
    <form>
        <div class="flex justify-between sticky top-0 z-10 bg-light py-2">
            <div>
                <h1>Edit deck</h1>

            </div>
            <div class="flex justify-end gap-2">
                <button @click="$router.back()">Cancel</button>

                <button @click="submitUpdateDeck('')">Save</button>
            </div>
        </div>

        <div class="card" v-for="(card, index) in cardItems" :key="card.id">
            <div class="flex justify-between mb-1">
                <div>
                    {{ index + 1 }}
                </div>

                <button @click="deleteCardItem(index)">
                    <TrashIcon class="w-5 h-5 deleteicon " />
                </button>
            </div>
            <div class="flex justify-between gap-4">
                <div class="flex flex-col flex-1">
                    <AdaptiveTextarea :max-length=1000 v-model="card.front" />
                    Front
                </div>
                <div class="flex flex-col flex-1">
                    <AdaptiveTextarea :max-length=1000 v-model="card.back" />
                    Back
                </div>
            </div>

        </div>
        <div class="flex justify-center">
            <button @click="createCardItem">Add new card</button>
        </div>
    </form>
</template>

<style scoped>
    .card{
        @apply mb-4 px-4 py-2
        border border-border rounded-md
    }

    textarea{
        @apply border border-border rounded-md
    }


    button:disabled .deleteicon{
        @apply stroke-gray-400
    }
    
    .errorMsg{
        @apply text-sm text-red-600
    }

</style>