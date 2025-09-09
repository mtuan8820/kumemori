import type { Card } from "@/models/Card"
import { ref, onMounted } from "vue"
import {GetCard, ListCards, CreateCard, DeleteCard} from '../../wailsjs/go/service/CardService'
export function useCardViewModel(){
    const cards = ref<Card[]>([])
    
    async function loadCards(){
        const cardEntities = await loadCards() 
    }
}