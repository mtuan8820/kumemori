<script setup>
import { RouterLink } from 'vue-router';
import { Squares2X2Icon, PuzzlePieceIcon, FolderIcon, FolderPlusIcon, Cog6ToothIcon } from '@heroicons/vue/24/outline'

</script>
<template>
<div class="pl-4 pt-2 flex flex-col border-r border-border">
    <div>
        <RouterLink class="menuItem" to="/"> 
            <Squares2X2Icon class="icon"/> 
            Dashboard
        </RouterLink>
        <RouterLink class="menuItem" to="/">
            <PuzzlePieceIcon class="icon"/>
            Create Quiz
        </RouterLink>
    </div>
    <div class="flex-1 overflow-auto">
        <h1>Recent Decks</h1>
        <div class="menuItem">
            <FolderIcon class="icon" />
            Deck example 1
        </div>
        <div class="menuItem">
            <FolderIcon class="icon" />
            Deck example 2
        </div>
        <div class="menuItem">
            <FolderIcon class="icon" />
            Deck example 3
        </div>
        <div class="menuItem">
            <FolderIcon class="icon" />
            Deck example 4
        </div>
        <div class="menuItem">
            <FolderIcon class="icon" />
            Deck example 5
        </div>
        <button class="menuItem">
            <FolderPlusIcon class="icon" />
            Create New Deck
        </button>
    </div>

    <div>
        <button class="menuItem">
            <Cog6ToothIcon class="icon" />
            Setting
        </button>
    </div>
</div>
</template>

<style scoped>
    .menuItem{
        @apply pr-4 pl-2 text-sm flex mb-1 gap-3 text-textmuted cursor-pointer
        hover:font-medium hover:text-text
    }
    

    .router-link-active{
        @apply font-medium text-base text-text
    }

    .icon{
        @apply h-5 w-5
    }
</style>