import { onMounted, ref } from "vue";
import { GetCards } from "../../wailsjs/go/application/Factory";
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
            const cards = await GetCards(deckId)
            return cards
        } catch (err) {
            console.error("Failed to load cards:", err)
            return []
        }
    }

    onMounted(() => {
        loadCards().then(
            res => cardItems.value = res.map((card: { ID: any; Front: any; Back: any; }) => ({
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
        // If it's a newly added card (id=0), just remove it locally (do not send delete to backend)
        if (id === 0) {
            cardItems.value = cardItems.value.filter((_, i) => i !== index)
            return
        }
        // Mark existing card as deleted
        cardItems.value[index].visible = false
        cardItems.value[index].action = "delete"
    }

    const onEditCard = (index: number) => {
        const item = cardItems.value[index]
        if (!item) return
        if (item.action === "delete") return
        if (item.id === 0 && item.action === "add") return
        item.action = "update"
    }

    const submitUpdateDeck = async () => {
        try {
            if (name.value == null ){
                console.log("deck name must not be empty")
                return
            } 
            console.log('start deck update')
            
            // Build update list, skipping delete actions for new cards (id=0)
            const updateCardList = await Promise.all(
                cardItems.value
                    .filter(card => !(card.action === "delete" && card.id === 0))
                    .map(card => NewUpdateCardInput(card.id, card.front, card.back, card.action))
            )
            const updateInput = await NewUpdateInput(deckId, name.value??deckName, updateCardList.length, updateCardList)
            console.log(updateInput)
            await UpdateDeck(updateInput)
            console.log('deck update done')

            router.back()
        }
        catch (error) {
            console.log(error)
        }
    }

    return { name, cardItems, createCardItem, deleteCardItem, submitUpdateDeck, onEditCard }
}