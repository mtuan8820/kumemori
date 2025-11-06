import type {model} from "../../wailsjs/go/models";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { CreateDeck, NewCreateDeckInput } from "../../wailsjs/go/application/Factory";

interface CardItem{
    front: string,
    back: string,
    frontErrorMsg: string,
    backErrorMsg:string,
}

export function useCreateDeckViewModel(){
    const router = useRouter()
    const cardItems = ref<CardItem[]> ([{front:"", back:"", frontErrorMsg: "", backErrorMsg: ""}]);
    const name = ref("")
    const loading = ref(false)
    const disableDeleteCardItem = ref(true)
    const nameErrorMsg = ref("")

    const validateForm = () => {
        let flag = true
        if (name.value==""){
            nameErrorMsg.value = "Please enter a title to create your deck."
            flag = false
        } 
        
        return flag
    }

    const createCardItem = () => {
        cardItems.value = [...cardItems.value, {front:"", back:"", frontErrorMsg: "", backErrorMsg: ""}]
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

        if (validateForm() == false) {
            return
        }
        try {
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

            const createInput = await NewCreateDeckInput(name.value, deckCards)
            await CreateDeck(createInput)
            router.back()

        } catch (err) {
            console.error(err)
        }
    
    }

    return {cardItems, name, createCardItem, loading, submitCreateDeck, deleteCardItem, disableDeleteCardItem, nameErrorMsg}
}