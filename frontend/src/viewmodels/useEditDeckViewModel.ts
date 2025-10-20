import { onMounted, ref } from "vue";
import { FindAllCards } from "../../wailsjs/go/service/DeckService";
import { NewUpdateCardInput, NewUpdateInput, UpdateDeck } from "../../wailsjs/go/application/Factory";

import router from "@/router";

interface CardItem {
    id: number
    front: string,
    back: string,
    frontErrorMsg: string,
    backErrorMsg: string,
    action: string,
    visible: boolean,
}

export function useEditDeckViewModel(deckId: number, deckName: string | null | undefined) {
    const name = ref(deckName)
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
                backErrorMsg: "", 
                action: "not changed",
                visible: true
            }))
        ).catch(
            error => {
                console.error('Error:', error);
            })
    })

    const createCardItem = () => {
        cardItems.value = [...cardItems.value, { id: 0, front: "", back: "", frontErrorMsg: "", backErrorMsg: "" , action: "add", visible: true}]
        // can remove card item if current quantity >= 2
        disableDeleteCardItem.value = false
    }

    const deleteCardItem = (index: number, id: number) => {
        // make sure Card Item at least 1
        if (cardItems.value.length <= 1) {
            disableDeleteCardItem.value = true
        }
        cardItems.value[index].visible = false
        cardItems.value[index].action = "delete"
    }

    const submitUpdateDeck = async () => {
        try {
            if (name.value == null ){
                console.log("deck name must not be empty")
                return
            } 

            const updateCardList = await Promise.all(cardItems.value.map(card=>NewUpdateCardInput(card.id, card.front, card.back, card.action)))
            const updateInput = await NewUpdateInput(deckId, deckName??'', updateCardList.length, updateCardList)

            await UpdateDeck(updateInput)

            router.back()
        }
        catch (error) {
            console.log(error)
        }
    }

    return { name, cardItems, createCardItem, deleteCardItem, submitUpdateDeck }
}