import type {model} from "../../wailsjs/go/models";
import { ref } from "vue";
import { CreateDeck  } from "../../wailsjs/go/service/DeckService";

interface CardItem{
    front: string,
    back: string
}

export function useCreateDeckViewModel(){
    const cardItems = ref<CardItem[]> ([{front:"", back:""}]);
    const name = ref("")
    const loading = ref(false)
    const disableDeleteCardItem = ref(true)

    const createCardItem = () => {
        cardItems.value = [...cardItems.value, {front:"", back:""}]
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

    const submitCreateDeck = async () => {
        loading.value = true

        

        const deckCards: model.Card[] = cardItems.value.map(item => ({
            ID: 0,
            DeckID: 0,
            Front: item.front,
            Back: item.back,
            CreatedAt: null,
            Repetitions: 0,
            Lapses: 0,
            EaseFactor: 0,
            Interval: 0,
            Due: null,
            LastReviewed: null,
            convertValues: () => {}
        }))


        await CreateDeck(name.value, deckCards)

        loading.value = false
    }

    return {cardItems, name, createCardItem, loading, submitCreateDeck, deleteCardItem, disableDeleteCardItem}
}