import { ref, onMounted } from "vue"

import type { model } from "../../wailsjs/go/models" 
import { GetAllDecks } from "../../wailsjs/go/application/Factory"

export function useDeckViewModel(){
    const decks = ref<model.Deck[]>([])

    async function loadDecks() {
        decks.value = await GetAllDecks()
    }
    
    async function createDeck(name:string) {

    }

    async function deleteDeck(id: number){
    }

    async function updateDeck(id: number, name: string){
        // await Save(id, name)
    }

    onMounted(() => {
        loadDecks()
    })

    return {decks, loadDecks, createDeck, deleteDeck, updateDeck}
}
