import './App.css';
import {ListDecks} from "../wailsjs/go/service/DeckService"
import { useEffect, useState } from 'react';
import { DeckList } from './features/deck/deckList';

function App() {


    return (
        <div id="App">
            <DeckList />
        </div>
    )
}

export default App
