<script setup lang="ts">
import { pick } from '../utils/arrays';
import { flipCoin } from '../utils/booleans';
import Piece from './Piece.vue';


let grabbingPiece: any = undefined;


function onPieceGrab(event: Event): void {
    const element = event.currentTarget;
    
    grabbingPiece = element;
    onMouseMove(event)
}

function onMouseMove(event: Event): void {
    if(grabbingPiece === undefined) {
        return;
    }

    const mouseX = event.clientX - 20;
    const mouseY = event.clientY - 20;

    grabbingPiece.style.position = 'absolute';
    grabbingPiece.style.left = `${mouseX}px`;
    grabbingPiece.style.top = `${mouseY}px`;
}

function onPieceRelease(event: Event) {
    grabbingPiece.style = '';
    grabbingPiece = undefined;
}
</script>

<template>
    <div id="board" @mousemove="onMouseMove">
        <div class="cell" v-for="_ in 9 ** 2">
            <Piece 
                :kanji="pick(['王', '飛', '龍', '角', '馬', '金', '銀', '全', '桂', '今', '香', '杏', '歩', 'と'])" 
                :promoted="flipCoin()" 
                :grabbable="flipCoin()"
                @mousedown="onPieceGrab"
                @mouseup="onPieceRelease"
                v-if="flipCoin(0.2)"
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