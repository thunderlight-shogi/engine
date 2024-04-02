<script setup lang="ts">
import { generateUUIDv4 } from '../crypto/uuids';
import { useBoard } from '../stores/board-store';
import { TYPES } from '../thunderlight/piece-type';
import { Player } from '../thunderlight/player';
import DraggablePiece from './DraggablePiece.vue';

defineProps<{ 
    player: Player,
}>();

const board = useBoard();
</script>

<template>
<div class="inventory">
    <div class="inventory-cell" v-for="type of TYPES.filter(type => !type.promoted)">
        <DraggablePiece 
            :key="generateUUIDv4()"
            :kanji="type.kanji"
            :promoted="type.promoted" 
            :grabbable="player === board.player"
            :enemy="player === 'gote'"
            state="idle"
        ></DraggablePiece>
    </div>
</div>
</template>

<style scoped lang="sass">
.inventory
    display: flex
    align-items: center
    justify-content: center
    width: 4em
    background: $gray
    flex-direction: column
    gap: 1em
</style>