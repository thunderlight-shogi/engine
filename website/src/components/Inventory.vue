<script setup lang="ts">
import { generateUUIDv4 } from '../crypto/uuids';
import { byClass, closest, indexOfChild } from '../dom/dom';
import { useBoard } from '../stores/board-store';
import { Coordinate } from '../thunderlight/coordinate';
import { PieceType } from '../thunderlight/piece-type';
import { Player } from '../thunderlight/player';
import { jukebox } from '../utils/jukebox';
import DraggablePiece from './DraggablePiece.vue';

defineProps<{ 
    player: Player,
}>();

const hand = defineModel<HTMLElement | undefined>(undefined);
const board = useBoard();

function getCells(): HTMLElement[] {
    return byClass("cell");
}

function locateHand(): Coordinate {
    return new Coordinate(4, 4);

    // if (hand.value === undefined) {
    //     throw new Error("Cannot locate an empty hand");
    // }

    // const cell = hand.value.parentElement;

    // if (cell === null) {
    //     console.error("The hand is ", hand.value, hand.value.parentElement)
    //     throw new Error("Cannot locate a hand with no parent")
    // }

    // return locateCell(cell);
}

function locateCell(cell: HTMLElement): Coordinate {
    const childIndex = indexOfChild(cell);
    const x = childIndex % 9;
    const y = Math.floor(childIndex / 9); 

    return new Coordinate(x, y);
}

function fadeCells() {
    getCells().forEach(cell => {
        cell.classList.remove('highlighted')
    });
}

function highlight(cell: HTMLElement) {
    cell.classList.add('highlighted');
}

function onPieceGrab(element: HTMLElement) {
    console.info(element, hand.value);

    jukebox.play("piece.grab", 0.2);
    hand.value = element;
    hand.value.style.position = 'absolute';
}

async function onPieceDrop(_: HTMLElement, type: PieceType) {
    if (hand.value === undefined) {
        return;
    }

    const cells = getCells();
    const underlyingCell = closest(cells, hand.value);
    const destination = locateCell(underlyingCell);

    hand.value.style.position = '';
    hand.value.style.top = '';
    hand.value.style.left = '';

    board.drop(destination, type);
    jukebox.play("piece.drop", 0.3);
    
    fadeCells();
    hand.value = undefined;
}
</script>

<template>
<div class="inventory">
    <div class="inventory-slot" v-for="slot of board.getInventoryOf(player).slots">
        <DraggablePiece 
            :key="generateUUIDv4()"
            :type="slot.type" 
            :grabbable="player === board.player && slot.count !== 0"
            :enemy="player === 'gote'"
            @drop="onPieceDrop"
            @grab="onPieceGrab"
            state="idle"
        ></DraggablePiece>
        <span class="inventory-slot-count">{{ slot.count }}</span>
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

.inventory-slot
    display: flex
    align-items: center
    justify-content: center

.inventory-slot-count
    color: white
    background: $background
    width: 1.5em
    aspect-ratio: 1 / 1
    text-align: center
    border-radius: 3em
    margin-top: auto
    margin-left: -1em
    z-index: 1
</style>