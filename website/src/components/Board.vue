<script setup lang="ts">
import { ref } from 'vue';
import { generateUUIDv4 } from '../crypto/uuids';
import { byClass, closest, indexOfChild, locate, locateMouse, measure, move } from '../dom/dom';
import { firecracker } from '../particles/particles';
import { MoveType, useBoard } from '../stores/board-store';
import { Coordinate } from '../thunderlight/coordinate';
import { EngineMode } from '../thunderlight/engine-mode';
import { jukebox } from '../utils/jukebox';
import { sleep } from '../utils/sleep';
import DraggablePiece from './DraggablePiece.vue';
import Inventory from './Inventory.vue';
import ModeSwitch from './ModeSwitch.vue';


let hand: HTMLElement | undefined = undefined;
const board = useBoard();
const mode = ref<EngineMode>('board');

function getCells(): HTMLElement[] {
    return byClass("cell");
}

function locateHand(): Coordinate {
    if (hand === undefined) {
        throw new Error("Cannot locate an empty hand");
    }

    const cell = hand.parentElement;

    if (cell === null) {
        throw new Error("Cannot locate a hand with no parent")
    }

    return locateCell(cell);
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
    jukebox.play("piece.grab", 0.2);
    hand = element;
    hand.style.position = 'absolute';
}

async function onPieceDrop(piece: HTMLElement) {
    if (hand === undefined) {
        return;
    }

    const cells = getCells();
    const underlyingCell = closest(cells, hand);
    const source = locateHand();
    const destination = locateCell(underlyingCell);

    hand.style.position = '';
    hand.style.top = '';
    hand.style.left = '';

    const move = board.move(source, destination);

    switch (move) {
        case MoveType.TRAVEL:
            jukebox.play("piece.drop", 0.3);
            break;

        case MoveType.ATTACK:
            hand = undefined;
            
            await sleep(100);
            jukebox.play("piece.attack", 0.8, 0.3);

            await sleep(20);
            firecracker.splash("piece.shred", locate(underlyingCell), 30, 25);
            break;

        case MoveType.BACK:
            jukebox.play("piece.back", 0.8, 0.3);
            break;

        case MoveType.PROHIBITED:
            jukebox.play("piece.prohibited");
            break;
    }
    
    fadeCells();
    hand = undefined;
}

function onPieceMove(event: MouseEvent): void {
    if(hand === undefined) {
        return;
    }

    const mouse = locateMouse(event).shift(measure(hand).shorten(3));
    const cells = getCells();
    const underlyingCell = closest(cells, hand);

    console.info(`${mouse}`, hand);

    if(!underlyingCell.classList.contains('highlighted')) {
        fadeCells();
    }

    move(hand, mouse);
    highlight(underlyingCell);
}

</script>

<template>
    <div id="board-ui">
        <ModeSwitch v-model="mode"></ModeSwitch>

        <div id="board">
            <Inventory player="gote"></Inventory>

            <div id="cells" @mousemove="onPieceMove">
                <div class="cell" v-for="piece of board.cells">
                    <DraggablePiece 
                        v-if="piece !== undefined"
                        :key="piece?.id ?? generateUUIDv4()"
                        :kanji="piece.type.kanji"
                        :promoted="piece?.type.promoted" 
                        :grabbable="piece?.player === board.player"
                        :enemy="piece?.player === 'gote'"
                        :state="piece?.state"
                        ref="pieces"
                        @grab="onPieceGrab"
                        @drop="onPieceDrop"
                    ></DraggablePiece>
                </div>
            </div>

            <Inventory player="sente"></Inventory>
        </div>
    </div>
    
</template> 

<style scoped lang="sass">
$cell-size: 4em
$cell-gap: 0.66em

@keyframes cell-highlight
    0%
        transform: scale(1.00)
    100%
        transform: scale(0.95)

#board-ui
    display: flex
    align-items: center
    flex-direction: column
    row-gap: 1em

    & > *
        flex: 1

#board
    display: flex
    gap: 2em

#cells
    display: grid
    aspect-ratio: 1 / 1
    grid-template-rows: repeat(9, $cell-size)
    grid-template-columns: repeat(9, $cell-size)
    gap: $cell-gap
    user-select: none

.cell
    background: $gray
    display: flex
    align-items: center
    justify-content: center
    box-sizing: border-box
    transition: 200ms ease-in-out
    &.highlighted
        background: $primary
        border: 1px cyan solid
        box-shadow: 0 0 30px $primary

ion-icon
    color: $primary
</style>