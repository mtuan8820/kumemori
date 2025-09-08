import { ref, onMounted } from "vue"

import type { Deck } from "@/models/Deck" 
import {GetDeck, ListDecks, CreateDeck, DeleteDeck, UpdateDeck} from '../../wailsjs/go/service/DeckService'

export function useDeckViewModel(){
    const decks = ref<Deck[]>([])
    const loading = ref(false)

    async function loadDecks() {
        loading.value = true
        let deckEntities = await ListDecks()
        decks.value = deckEntities.map(deck => ({
            id: deck.ID,
            name: deck.Name,
        }))
        loading.value = false
    }

    async function createDeck(name:string) {
        if (name==null || name==""){
            return
        }

        loading.value = true

        const newDeckEntity = await CreateDeck(name)
        if (newDeckEntity){
            const deck: Deck = {
                id:newDeckEntity.ID,
                name: newDeckEntity.Name
            }
            decks.value.push(deck)
        }
        else {
            console.log("Failed to create deck")
        }

        loading.value = false
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

    return {decks, loading, loadDecks, createDeck, deleteDeck, updateDeck}
}

