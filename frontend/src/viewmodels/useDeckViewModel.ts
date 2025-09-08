import { ref } from "vue"
import { getDecks } from "@/services/deckService"
import type { Deck } from "@/models/Deck" 

export function useDeckViewModel(){
    const decks = ref<Deck[]>([])
    const loading = ref(false)

    async function loadDecks() {
        loading.value = true
        decks.value = await getDecks()
        loading.value = false
    }

    return {decks, loadDecks, loading}
}

