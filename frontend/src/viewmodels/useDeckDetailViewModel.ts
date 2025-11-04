import type {model} from "../../wailsjs/go/models";
import { onMounted, ref } from "vue";
import { GetCards } from "../../wailsjs/go/application/Factory";

export function useDeckDetailViewModel(deckId: number){
    const cards = ref<model.Card[]>([])

    async function loadCards(){
        try {
            cards.value = await GetCards(deckId)
        } catch (err) {
            console.error("Failed to load cards:", err)
        }
    }

    async function createCard(){
        // const card: Card = {
        //     front: "",
        //     back: ""
        // }
        // await AddCard(deckId, card)
    }

    onMounted(async () =>{
        await loadCards()
    })

    return { cards, createCard, loadCards}
}

