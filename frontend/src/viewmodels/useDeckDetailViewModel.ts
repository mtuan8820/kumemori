import type {model} from "../../wailsjs/go/models";
import { ref } from "vue";
import { ReadCardsByDeck, CreateCard } from "../../wailsjs/go/service/CardService";

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
        await CreateCard("Front 5","Back", deckId)
    }

    return { cards, createCard, loadCards}
}

