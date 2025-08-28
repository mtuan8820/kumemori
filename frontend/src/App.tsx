import './App.css';
import {ListDecks} from "../wailsjs/go/service/DeckService"
import { useEffect, useState } from 'react';
import loadDecks from './features/deck/deckList';

class Deck{
    id: number;
    name: string;

    constructor(id: number, name: string){
        this.name = name;
        this.id = id
    }
}


function App() {
    const [decks, setDecks] = useState<Deck[]>([])

    useEffect(()=>{
        async function fetchDecks(){
            const data = await loadDecks()
            setDecks(data)
        }
        fetchDecks()
    }, [])

    return (
        <div id="App">
            {decks.map((deck, index) => (
                <div key={index}>
                    <p>{deck.name}</p>
                </div>
            ))}
        </div>
    )
}

export default App
