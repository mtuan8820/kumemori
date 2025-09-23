import { onMounted, ref } from "vue";
import { FindAllCards, UpdateDeck } from "../../wailsjs/go/service/DeckService";
import type { model } from "../../wailsjs/go/models";

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
        cardItems.value = [...cardItems.value, { id: 0, front: "", back: "", frontErrorMsg: "", backErrorMsg: "" }]
        // can remove card item if current quantity >= 2
        disableDeleteCardItem.value = false
    }

    const deleteCardItem = (index: number) => {
        cardItems.value.splice(index, 1)
        // make sure Card Item at least 1
        if (cardItems.value.length <= 1) {
            disableDeleteCardItem.value = true
        }
    }

    const submitUpdateDeck = async (newName: string) => {
        newName = "new name";

        try {
            const deckCards: model.Card[] = cardItems.value.map(item => ({
                ID: 0,
                DeckID: deckId,
                Front: item.front,
                Back: item.back,
                CreatedAt: null,
                Repetitions: 0,
                Lapses: 0,
                EaseFactor: 0,
                Interval: 0,
                Due: null,
                LastReviewed: null,
                convertValues: () => { }
            }))

            await UpdateDeck(deckId, newName, deckCards)
        }
        catch (error) {
            console.log(error)
        }
    }

    return { cardItems, createCardItem, deleteCardItem, submitUpdateDeck }
}