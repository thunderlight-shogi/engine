<script setup lang="ts">
import { byClass, closest, indexOfChild } from '../dom/dom';
import Board from '../thunderlight/board';
import { Coordinate } from '../thunderlight/coordinate';
import { PieceType } from '../thunderlight/piece-type';
import { Player } from '../thunderlight/player';
import { jukebox } from '../utils/jukebox';

defineProps<{ 
    player: Player,
    board: Board,
}>();

const hand = defineModel<HTMLElement | undefined>(undefined);

function getCells(): HTMLElement[] {
    return byClass("cell");
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
    <!--div class="inventory-slot" v-for="slot of board.getInventoryOf(player).slots">
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
    </!--div-->
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