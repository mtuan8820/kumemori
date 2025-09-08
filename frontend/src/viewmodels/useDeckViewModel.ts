import { ref, onMounted } from "vue"
import { 
    getDecks as getDecksService,
    createDeck as createDeckService,
    deleteDeck as deleteDeckService,
    updateDeck as updateDeckService 
} from "@/services/deckService"

import type { Deck } from "@/models/Deck" 

export function useDeckViewModel(){
    const decks = ref<Deck[]>([])
    const loading = ref(false)

    async function loadDecks() {
        loading.value = true
        decks.value = await getDecksService()
        loading.value = false
    }

    async function createDeck(name:string) {
        if (name==null || name==""){
            return
        }

        loading.value = true

        const deck = await createDeckService(name)
        if (deck){
            decks.value.push(deck)
        }
        else {
            console.log("Failed to create deck")
        }

        loading.value = false
    }

    async function deleteDeck(id: number){
        await deleteDeckService(id)
    }

    async function updateDeck(id: number, name: string){
        await updateDeckService(id, name)
    }

    onMounted(() => {
        loadDecks()
    })

    return {decks, loading, loadDecks, createDeck, deleteDeck, updateDeck}
}


