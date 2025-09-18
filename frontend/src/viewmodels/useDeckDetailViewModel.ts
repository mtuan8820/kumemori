import type {model} from "../../wailsjs/go/models";
import { ref } from "vue";
import { FindAllCards } from "../../wailsjs/go/service/DeckService";

export function useDeckDetailViewModel(deckId: number){
    const cards = ref<model.Card[]>([])

    async function loadCards(){
        try {
            cards.value = await FindAllCards(deckId)
            return cards
        } catch (err) {
            console.error("Failed to load cards:", err)
            return []
        }
    }

    async function createCard(){
        // const card: Card = {
        //     front: "",
        //     back: ""
        // }
        // await AddCard(deckId, card)
    }

    return { cards, createCard, loadCards}
}

