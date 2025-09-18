import { ref, onMounted } from "vue"

import type { model } from "../../wailsjs/go/models" 
import { GetDecks,  Delete,} from '../../wailsjs/go/service/DeckService'

export function useDeckViewModel(){
    const decks = ref<model.Deck[]>([])

    async function loadDecks() {
        decks.value = await GetDecks()
    }

    async function createDeck(name:string) {

    }

    async function deleteDeck(id: number){
        await DeleteDeck(id)
    }

    async function updateDeck(id: number, name: string){
        await UpdateDeck(id, name)
    }

    onMounted(() => {
        loadDecks()
    })

    return {decks, loadDecks, createDeck, deleteDeck, updateDeck}
}

