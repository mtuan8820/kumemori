import type { Deck } from '../models/Deck'; 
import {GetDeck, ListDecks, CreateDeck} from '../../wailsjs/go/service/DeckService'
import type { Card } from '@/models/Card';

// export interface DeckService{
//   getDecks(): Promise<Deck[]>
//   getDeck(id: number): Promise<Deck>
//   createDeck(deck: Deck): Promise<Deck>
// }

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

export async function createDeck(deck: Deck): Promise<Deck|null> {
  let newDeck = await CreateDeck(deck.name)
  if (newDeck != null){
    return {
      id: newDeck.ID,
      name: newDeck.Name
    }
  }
  return null
}