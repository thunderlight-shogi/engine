<script setup lang="ts">
import { pick } from '../utils/arrays';
import { flipCoin } from '../utils/booleans';
import { jukebox } from '../utils/jukebox';
import Piece from './Piece.vue';


let hand: HTMLElement | undefined = undefined;

function onPieceGrab(element: HTMLElement) {
    jukebox.play("piece.grab", 0.2);
    hand = element;
}

function onPieceDrop(_: HTMLElement) {
    jukebox.play("piece.drop", 0.3);
    hand = undefined;
}

function onPieceMove(event: MouseEvent): void {
    if(hand === undefined) {
        return;
    }

    const mouseX = event.clientX - 20;
    const mouseY = event.clientY - 20;

    hand.style.position = 'absolute';
    hand.style.left = `${mouseX}px`;
    hand.style.top = `${mouseY}px`;
}


</script>

<template>
    <div id="board" @mousemove="onPieceMove">
        <div class="cell" v-for="_ in 9 ** 2">
            <Piece 
                :kanji="pick(['王', '飛', '龍', '角', '馬', '金', '銀', '全', '桂', '今', '香', '杏', '歩', 'と'])" 
                :promoted="flipCoin()" 
                :grabbable="flipCoin()"
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

ion-icon
    color: $primary
</style>