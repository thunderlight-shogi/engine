<script setup lang="ts">
import { byClass, closest } from '../dom/dom';
import { pick } from '../utils/arrays';
import { flipCoin } from '../utils/booleans';
import { jukebox } from '../utils/jukebox';
import Piece from './Piece.vue';


let hand: HTMLElement | undefined = undefined;

function getCells(): HTMLElement[] {
    return byClass("cell");
}

function fadeCells() {
    getCells().forEach(cell => cell.classList.remove('highlighted'));
}

function highlight(cell: HTMLElement) {
    cell.classList.add('highlighted');
}

function onPieceGrab(element: HTMLElement) {
    jukebox.play("piece.grab", 0.2);
    hand = element;
}

function onPieceDrop(_: HTMLElement) {
    if (hand === undefined) {
        return;
    }

    const cells = getCells();
    const underlyingCell = closest(cells, hand);

    hand.style.position = '';
    hand.style.top = '';
    hand.style.left = '';
    underlyingCell.appendChild(hand);
    jukebox.play("piece.drop", 0.3);
    fadeCells();

    hand = undefined;
}

function onPieceMove(event: MouseEvent): void {
    if(hand === undefined) {
        return;
    }

    const mouseX = event.pageX - 20;
    const mouseY = event.pageY - 20;

    hand.style.left = `${mouseX}px`;
    hand.style.top = `${mouseY}px`;

    const cells = getCells();
    const underlyingCell = closest(cells, hand);

    if(!underlyingCell.classList.contains('highlighted')) {
        fadeCells();
    }

    highlight(underlyingCell);
}
</script>

<template>
    <div id="board" @mousemove="onPieceMove">
        <div class="cell" v-for="_ in 9 ** 2">
            <Piece 
                :kanji="pick(['王', '飛', '龍', '角', '馬', '金', '銀', '全', '桂', '今', '香', '杏', '歩', 'と'])" 
                :promoted="flipCoin(0.1)" 
                :grabbable="flipCoin(0.3)"
                v-if="flipCoin(0.2)"
                @grab="onPieceGrab"
                @drop="onPieceDrop"
            ></Piece>
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

#board
    display: grid
    aspect-ratio: 1 / 1
    grid-template-rows: repeat(9, $cell-size)
    grid-template-columns: repeat(9, $cell-size)
    gap: $cell-gap

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
        box-shadow: 0 0 5px $primary

ion-icon
    color: $primary
</style>