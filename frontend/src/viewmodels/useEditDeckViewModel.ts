import { onMounted, ref } from "vue";
import { FindAllCards } from "../../wailsjs/go/service/DeckService";

interface CardItem {
    id: number
    front: string,
    back: string,
    frontErrorMsg: string,
    backErrorMsg: string,
}

export function useEditDeckViewModel(deckId: number) {

    const cardItems = ref<CardItem[]>([])
    const disableDeleteCardItem = ref(true)

    const loadCards = async () => {
        try {
            const cards = await FindAllCards(deckId)
            return cards
        } catch (err) {
            console.error("Failed to load cards:", err)
            return []
        }
    }

    onMounted(() => {
        loadCards().then(
            res => cardItems.value = res.map(card => ({
                id: card.ID,
                front: card.Front,
                back: card.Back,
                frontErrorMsg: "",
                backErrorMsg: ""
            }))
        ).catch(
            error => {
                console.error('Error:', error);
            })

    })

    const createCardItem = () => {
        cardItems.value = [...cardItems.value, {id: 0, front:"", back:"", frontErrorMsg: "", backErrorMsg: ""}]
        // can remove card item if current quantity >= 2
        disableDeleteCardItem.value = false
    }

    const deleteCardItem = (index:number) => {
        cardItems.value.splice(index, 1)
        // make sure Card Item at least 1
        if (cardItems.value.length <=1){
            disableDeleteCardItem.value = true
        }
    }

    return {cardItems, createCardItem, deleteCardItem}
}