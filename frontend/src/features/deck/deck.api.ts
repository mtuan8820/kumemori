import { CreateDeck, DeleteDeck, ListDecks } from '../../../wailsjs/go/service/DeckService'
import {Deck} from './deck.type'


export async function loadDecks(): Promise<Deck[]> {
  try {
    const result = await ListDecks()
    const decks = result.map(item => new Deck(item.ID, item.Name))
    return decks
  } catch (err) {
    console.error(err)
    return [] 
  }
}

export async function deleteDeck(id: number) {
    try {
        await DeleteDeck(id)
    }
    catch (err){
        console.error(err)
    }
}

export async function createDeck(name: string){
    try {
        let deck = await CreateDeck(name)
        return deck
    }
    catch (err){
        console.log(err)
    }

}