<script lang="ts" setup>
import AdaptiveTextarea from '@/components/AdaptiveTextarea.vue';
import {useCreateDeckViewModel} from '@/viewmodels/useCreateDeckViewModel'
import { TrashIcon } from "@heroicons/vue/24/outline";

const {cardItems, name, createCardItem, submitCreateDeck, deleteCardItem, disableDeleteCardItem} = useCreateDeckViewModel()


</script>
<template>
    <form >
        <div class="flex justify-between sticky top-0 z-10 bg-light py-2">
            <div>
                <h1>Create a new deck</h1>
                <input v-model="name" placeholder="name"/>
            </div>
            <div class="flex justify-end">
                <button @click="submitCreateDeck">Create</button>
            </div>
        </div>
        
        <div class="card" v-for="(card, index) in cardItems">
            <div class="flex justify-between mb-1">
                <div>
                    {{ index+1 }}
                </div>

                <button @click="deleteCardItem(index)" :disabled="disableDeleteCardItem" >
                    <TrashIcon  class="w-5 h-5 deleteicon "/>
                </button>
            </div>
            <div class="flex justify-between gap-4">
                <div class="flex flex-col flex-1">
                    <AdaptiveTextarea v-model="card.front"/>
                    Front
                </div> 
                <div class="flex flex-col flex-1">
                    <AdaptiveTextarea v-model="card.back"/>
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
    

</style>