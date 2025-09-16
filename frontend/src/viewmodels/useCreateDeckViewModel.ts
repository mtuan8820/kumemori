import type { Card } from "@/models/Card";
import { ref } from "vue";

interface CardItem{
    front: string,
    back: string
}

export function useCreateDeckViewModel(){
    const cardItems = ref<CardItem[]> ([{front:"", back:""}]);
    const name = ref("")

    const createCardItem = () => cardItems.value = [...cardItems.value, {front:"", back:""}]

    return {cardItems, name, createCardItem}
}