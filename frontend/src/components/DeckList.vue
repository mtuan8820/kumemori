<script setup>
    import { useDeckViewModel } from '@/viewmodels/useDeckViewModel'
    import { onMounted, ref } from 'vue'
    import Dialog from './Dialog.vue'

    const { decks, loading, loadDecks,  createDeck, deleteDeck, updateDeck } = useDeckViewModel()

    const deckName = ref("")

    const dialogVisible = ref(false)
    const selectedDeckId = ref(-1)
    const deckUpdateName = ref("")

    async function handleCreateDeck(){
        await createDeck(deckName.value)
        deckName.value = ""
        await loadDecks()
    }

    async function handleDeleteDeck(id){
        await deleteDeck(id)
        await loadDecks()
    }

    function openRenameDialog(id, name){
        console.log(id, name)
        selectedDeckId.value = id
        deckUpdateName.value = name
        dialogVisible.value = true
    }

    async function handleRenameDeck(){
        if (deckUpdateName.value == "" || selectedDeckId.value == -1){
            dialogVisible.value = false
            return
        }
        await updateDeck(selectedDeckId.value, deckUpdateName.value)
        deckUpdateName.value = ""
        selectedDeckId.value = -1
        dialogVisible.value = false
        await loadDecks()
        return
    }


</script>

<template>
        <div>
        <h1>
            Decks
        </h1>
        <form>
            <p>New deck name:</p>
            <input v-model.trim="deckName" placeholder=""/>
            <button @click="handleCreateDeck">Create</button>
        </form>
        <Dialog :is-open="dialogVisible" @close="dialogVisible=false">
            <form>
                <p>New deck name:</p>
                <input v-model.trim = "deckUpdateName"/>
                <button @click="handleRenameDeck">Confirm</button>
            </form>
        </Dialog>

        <ul>
            <li v-for="deck in decks" :key="deck.id">
                {{ deck.name }}
                <button @click="handleDeleteDeck(deck.id)">x</button>
                {{" "}}
                <button @click="openRenameDialog(deck.id, deck.name)">Edit</button>
            </li>
        </ul>
    </div>
</template>