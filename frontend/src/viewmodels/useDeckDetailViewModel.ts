import type { Card } from "@/models/Card";
import { ref } from "vue";
import { ReadCardsByDeck, CreateCard } from "../../wailsjs/go/service/CardService";

export function useDeckDetailViewModel(deckId: number){
    const cards = ref<Card[]>([])

    async function loadCards(){
        const cardEntities = await ReadCardsByDeck(deckId);
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
        await CreateCard("Front 5","Back", deckId)
    }

    return { cards, createCard, loadCards}
}

