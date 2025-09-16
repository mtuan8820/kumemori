<script lang="ts" setup>
import {useCreateDeckViewModel} from '@/viewmodels/useCreateDeckViewModel'

const {cardItems, name, createCardItem} = useCreateDeckViewModel()

</script>
<template>
    <form>
        <h1>Create a new deck</h1>
        <input v-text="name" placeholder="name"/>
        
        <div v-for="(card, index) in cardItems">
            <div>
                {{ index+1 }}
            </div>
            Front: <input v-text="" placeholder=""/>
            Back :<input v-text="" placeholder=""/>
        </div>
        <button @click="createCardItem">Add new card</button>
    </form>


</template>