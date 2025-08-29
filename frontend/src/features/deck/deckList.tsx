import { useEffect, useState } from "react"
import { Deck } from "./deck.type"
import { loadDecks } from "./deck.api"
import { createDeck, deleteDeck } from "./deck.api"

type CreateDeckFormProps = {
  onCreate: (deck: Deck) => void;
};

function CreateDeckForm({ onCreate }: CreateDeckFormProps) {
    async function handleSubmit(e: React.FormEvent<HTMLFormElement>){
        e.preventDefault(); 
        const formData = new FormData(e.currentTarget);
        const name = formData.get('name')
        if (typeof name !== 'string') {
            throw new Error('Name is required');
        }
        let deck = await createDeck(name)
        if (deck != null){
            onCreate(deck)
        }

        e.currentTarget.reset();
    }

    return (
        <form onSubmit={handleSubmit}>
            <label>
                Name: <input name="name" />
            </label>
            <button type="submit">Create</button>
        </form>
    )
}

type DeckItemProps = {
  name: string;
  id: number;
  onDelete?: (id: number) => void;
  callbackToUpdateState: (id: number) => void 
};

function DeckItem({ name, id, onDelete, callbackToUpdateState }: DeckItemProps) {
  return (
    <div>
      {name}{" "}
      <button onClick={() => {onDelete?.(id); callbackToUpdateState(id)}}>
        X
      </button>
    </div>
  );
}



export function DeckList(){
    
    const [decks, setDecks] = useState<Deck[]>([])

    useEffect(()=>{
        async function fetchDecks(){
            const data = await loadDecks()
            setDecks(data)
        }
        fetchDecks()
    }, [])

    function addDeck(deck: Deck){
        setDecks(prev => [...prev, deck])
    }

    function removeDeck(id: number){
        const updatedDecks = decks.filter((item)=>{
            return item.ID !== id
        })
        setDecks(updatedDecks)
    }

    return (
        <div>
            <CreateDeckForm onCreate={addDeck}/>
            <div>
                {decks.map(item=><DeckItem name={item.Name} id={item.ID} onDelete={deleteDeck} callbackToUpdateState={removeDeck} />)}
            </div>
        </div>
    )
}