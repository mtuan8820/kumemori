<script setup>
    import { useDeckViewModel } from '@/viewmodels/useDeckViewModel'
    const { decks, loadDecks, loading } = useDeckViewModel()
</script>

<template>
    <div>
        <h1>
            Decks
        </h1>
        <ul>
            <li v-for="deck in decks" :key="deck.id">
                {{ deck.name }}
            </li>
        </ul>
    </div>
</template>