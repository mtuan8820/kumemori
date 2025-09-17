import type { Card } from "@/models/Card";
import { ref } from "vue";
import { FindAllCards } from "../../wailsjs/go/service/DeckService";

export function useDeckDetailViewModel(deckId: number){
    const cards = ref<Card[]>([])

    async function loadCards(){
        const cardEntities = await FindAllCards(deckId);
        console.log('cards', cardEntities)
        if (cardEntities){
            cards.value = cardEntities.map(card => ({
                id: card.ID,
                deckId: deckId,
                front: card.Front,
                back: card.Back
            }))
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

