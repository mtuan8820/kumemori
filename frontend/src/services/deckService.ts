import type { Deck } from '@/models/Deck'; 
import {GetDeck, ListDecks, CreateDeck, DeleteDeck, UpdateDeck} from '../../wailsjs/go/service/DeckService'
import type { Card } from '@/models/Card';

export async function getDecks(): Promise<Deck[]> {
  let decks = await ListDecks()
  let result: Deck[] = []
  for (let deck of decks) {
     var cards: Card[] = []
     var d: Deck ={
      id: deck.ID,
      name: deck.Name,
     }
     result = [d,...result]
    }
  return result
  
}

export async function getDeck(id: number): Promise<Deck | null> {
  let deck = await GetDeck(id)
  if (deck != null){
    return {
      id: deck.ID,
      name: deck.Name
    }
  }
  return null
}

export async function createDeck(name: string): Promise<Deck|null> {
  let newDeck = await CreateDeck(name)
  if (newDeck){
    const deck: Deck={
      id: newDeck.ID,
      name: newDeck.Name
    }
    return deck
  }
  return null
}

export async function deleteDeck(id: number) {
  await DeleteDeck(id)
}

export async function updateDeck(id: number, name: string){
  await UpdateDeck(id, name)
}