<script setup>
import { useDeckViewModel } from '@/viewmodels/useDeckViewModel'
import { ref } from 'vue'
import Dialog from './Dialog.vue'
import { FolderIcon, MinusIcon, PencilIcon, PlusIcon, MagnifyingGlassIcon } from '@heroicons/vue/24/outline'
import { RouterLink } from 'vue-router'

const { decks, loading, loadDecks, createDeck, deleteDeck, updateDeck } = useDeckViewModel()

const deckName = ref("")

const dialogVisible = ref(false)
const selectedDeckId = ref(-1)
const deckUpdateName = ref("")

async function handleCreateDeck() {
    // await createDeck(deckName.value)
    // deckName.value = ""
    // await loadDecks()
    console.log("create")
}

async function handleDeleteDeck(id) {
    await deleteDeck(id)
    await loadDecks()
}

function openRenameDialog(id, name) {
    console.log(id, name)
    selectedDeckId.value = id
    deckUpdateName.value = name
    dialogVisible.value = true
}

async function handleRenameDeck() {
    if (deckUpdateName.value == "" || selectedDeckId.value == -1) {
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
        <h1 class="mb-3 text-lg">
            Decks
        </h1>
        <div class="flex justify-between mb-3 items-center">
            <div></div>
            <div class="flex items-center gap-3">
                <RouterLink to="/create-deck"><PlusIcon class="w-5 h-5"/></RouterLink>
                
                <div class="search">
                    <input type="text" class="search__input" placeholder="Search flashcards"/>
                    <MagnifyingGlassIcon class="w-5 h-5"/>
                </div>

            </div>
        </div>
        <Dialog :is-open="dialogVisible" @close="dialogVisible = false">
            <form>
                <p>New deck name:</p>
                <input v-model.trim="deckUpdateName" />
                <button @click="handleRenameDeck">Confirm</button>
            </form>
        </Dialog>

        <ul>
            <li v-for="deck in decks" :key="deck.ID" class="3">
                <FolderIcon class="h-7 w-7 mr-2"/>
                <div class="">
                        <div @click="$router.push(`/deck/${deck.ID}/${deck.Name}`)">{{ deck.Name }}</div>
                        <div class="text-xs">
                            <span>Last Studied: Sep 14, 25</span>  •  
                            <span>Progress: 15/30</span>
                        </div>
                </div>
                <div class="flex items-center gap-3">
                    <button @click="openRenameDialog(deck.ID, deck.Name)"><PencilIcon class="h-5 w-5"/></button>
                    <button @click="handleDeleteDeck(deck.ID)"><MinusIcon class="h-5 w-5"/></button>
                </div>
            </li>
        </ul>
    </div>
</template>

<style scoped>
li{
    @apply grid grid-cols-[auto,1fr,auto] gap-1 w-full items-center 
    p-3 rounded-lg
    hover:bg-dark
}

.search{
    @apply bg-light border border-black px-2 py-1 rounded-md 
    flex items-center gap-1
    focus-within:border-blue-500 
    focus:border-none
}

.search__input{
    @apply flex-1 border-none outline-none 
    dark:bg-light
}

</style>