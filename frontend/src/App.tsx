import './App.css';
import {ListDecks} from "../wailsjs/go/service/DeckService"
import { useEffect, useState } from 'react';
import { DeckList } from './features/deck/deckList';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import Quiz from './pages/Quiz';
import Test from './pages/Test';
import Sidebar from './components/Sidebar';
import { routes } from './routes';

function App() {

    return (
        <div id="App">
            <BrowserRouter>
                <Sidebar routes={routes}/>
                <Routes>
                    {routes.map((item)=><Route path={item.to} element={<item.element/>} />)}
                </Routes>
            </BrowserRouter>
        </div>
    )
}

export default App
