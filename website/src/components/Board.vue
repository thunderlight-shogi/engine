<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { generateUUIDv4 } from '../crypto/uuids';
import { byClass, closest, indexOfChild, locate, locateMouse, measure, move } from '../dom/dom';
import { firecracker } from '../particles/particles';
import { useBoard } from '../stores/board-store';
import { Coordinate } from '../thunderlight/coordinate';
import { EngineMode } from '../thunderlight/engine-mode';
import { jukebox } from '../utils/jukebox';
import { sleep } from '../utils/sleep';
import DraggablePiece from './DraggablePiece.vue';
import Inventory from './Inventory.vue';
import ModeSwitch from './ModeSwitch.vue';
import { useFetch } from '@vueuse/core';
import { BestMove } from '../thunderlight/best-move';
import BestMoveDisplay from './BestMoveDisplay.vue';
import { PAWN } from "../thunderlight/piece-type";

const hand = ref<HTMLElement | undefined>(undefined);
const board = useBoard();
const bestMove = ref<BestMove>(new BestMove(true, new Coordinate(0, 0), new Coordinate(0, 0), "travel", PAWN));
const mode = ref<EngineMode>('board');
const { data } = useFetch("http://localhost:5173/start/").post({
    id: 1,
});

onMounted(async () => {
    console.log(data.value);
})

function getCells(): HTMLElement[] {
    return byClass("cell");
}

function locateHand(): Coordinate {
    if (hand.value === undefined) {
        throw new Error("Cannot locate an empty hand");
    }

    const cell = hand.value.parentElement;

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
    hand.value = element;
    hand.value.style.position = 'absolute';

    console.info("piece.grab")
}

async function onPieceDrop(_: HTMLElement) {
    if (hand.value === undefined) {
        return;
    }

    console.info("piece.drop")

    const cells = getCells();
    const underlyingCell = closest(cells, hand.value);
    const source = locateHand();
    const destination = locateCell(underlyingCell);

    hand.value.style.position = '';
    hand.value.style.top = '';
    hand.value.style.left = '';

    const move = board.move(source, destination);

    switch (move) {
        case 'travel':
            jukebox.play("piece.drop", 0.3);
            break;

        case 'attack':
            hand.value = undefined;
            
            await sleep(100);
            jukebox.play("piece.attack", 0.8, 0.3);

            await sleep(20);
            firecracker.splash("piece.shred", locate(underlyingCell), 30, 25);
            break;

        case 'back':
            jukebox.play("piece.back", 0.8, 0.3);
            break;

        case 'prohibited':
            jukebox.play("piece.prohibited");
            break;
    }

    if (move != 'prohibited') {
        console.log("The move is allowed, sending move/player request")

        const { isFetching, error, data } = await useFetch("http://localhost:5173/move/player").post({
            old_pos: {
                file: source.x,
                rank: source.y,
            },

            new_pos: {
                file: destination.x,
                rank: destination.y,
            },
        });

        console.log(`Answer is received (= ${JSON.stringify(isFetching)}, ${JSON.stringify(error)}, ${JSON.stringify(data)})`);

        await sleep(2000);

        console.log("Requesting to move/help...")

        const resp = await useFetch("http://localhost:5173/move/help").json().post();

        await sleep(500);
        
        while(true) {
            console.log("Wait for it...")

            console.log(resp);
            await sleep(250);
        }
    } 
    
    fadeCells();
    hand.value = undefined;
}

function onPieceMove(event: MouseEvent): void {
    if(hand.value === undefined) {
        return;
    }

    const mouse = locateMouse(event).shift(measure(hand.value).shorten(3));
    const cells = getCells();
    const underlyingCell = closest(cells, hand.value);

    if(!underlyingCell.classList.contains('highlighted')) {
        fadeCells();
    }

    move(hand.value, mouse);
    highlight(underlyingCell);
}

</script>

<template>
    <div id="board-ui">
        <ModeSwitch v-model="mode"></ModeSwitch>

        <div id="board">
            <Inventory player="gote" v-model="hand"></Inventory>

            <div id="cells" @mousemove="onPieceMove">
                <div class="cell" v-for="piece of board.cells">
                    <DraggablePiece 
                        v-if="piece !== undefined"
                        :key="piece?.id ?? generateUUIDv4()"
                        :type="piece.type"
                        :grabbable="piece?.player === board.turn"
                        :enemy="piece?.player === 'gote'"
                        :state="piece?.state"
                        ref="pieces"
                        @grab="onPieceGrab"
                        @drop="onPieceDrop"
                    ></DraggablePiece>
                </div>
            </div>

            <Inventory player="sente" v-model="hand"></Inventory>
        </div>

        <BestMoveDisplay v-model="bestMove"></BestMoveDisplay>
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