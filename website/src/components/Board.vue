<script setup lang="ts">
import { byClass, closest, heightOf, locate, locateMouse, measure, move, widthOf } from '../dom/dom';
import { firecracker } from '../particles/particles';
import { pick } from '../utils/arrays';
import { flipCoin } from '../utils/booleans';
import { Location2D, distanceBetween } from '../utils/geometry';
import { jukebox } from '../utils/jukebox';
import { translate } from '../utils/numbers';
import { sleep } from '../utils/sleep';
import Piece from './Piece.vue';


let hand: HTMLElement | undefined = undefined;

function getCells(): HTMLElement[] {
    return byClass("cell");
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
}

async function onPieceDrop(_: HTMLElement, mouseLocation: Location2D) {
    if (hand === undefined) {
        return;
    }

    const location = locate(hand);
    const cells = getCells();
    const underlyingCell = closest(cells, hand);

    hand.style.position = '';
    hand.style.top = '';
    hand.style.left = '';

    if(underlyingCell.children.length === 0) {
        jukebox.play("piece.drop", 0.3);
        underlyingCell.replaceChildren(hand);
    } else if(underlyingCell.firstChild === hand) {
        jukebox.play("piece.back", 0.3);
    } else {
        underlyingCell.replaceChildren(hand);
        hand = undefined;
        
        await sleep(100);
        jukebox.play("piece.attack", 0.5, 1);

        await sleep(20);
        firecracker.splash("piece.shred", locate(underlyingCell), 20, 25);
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

    if(!underlyingCell.classList.contains('highlighted')) {
        fadeCells();
    }

    move(hand, mouse);
    highlight(underlyingCell);
}

class PieceType {
    constructor(public readonly kanji: string, public readonly promoted: boolean = false) {}
}

class ActivePiece {
    constructor(public readonly type: PieceType,
                public readonly mine: boolean) {}
}

const pawn: PieceType = { kanji: '歩', promoted: false };
const king: PieceType = { kanji: '王', promoted: false };
const rook: PieceType = { kanji: '飛', promoted: false };
const dragon: PieceType = { kanji: '龍', promoted: false };
const bishop: PieceType = { kanji: '角', promoted: false };
const horse: PieceType = { kanji: '馬', promoted: false };
const gold: PieceType = { kanji: '金', promoted: false };
const silver: PieceType = { kanji: '銀', promoted: false };
const knight: PieceType = { kanji: '桂', promoted: false };
const goldenKnight: PieceType = { kanji: '今', promoted: false };
const lance: PieceType = { kanji: '香', promoted: false };
const goldenLance: PieceType = { kanji: '杏', promoted: false };
const tokin: PieceType = { kanji: 'と', promoted: false };

function e(type: PieceType) {
    return new ActivePiece(type, false);
}

function f(type: PieceType) {
    return new ActivePiece(type, true);
}

const board: (ActivePiece | null)[] = [
    e(lance), e(knight), e(silver), e(gold), e(king), e(gold), e(silver), e(knight), e(lance),
    null, e(rook), null, null, null, null, null, e(bishop), null,
    e(pawn), e(pawn), e(pawn), e(pawn), e(pawn), e(pawn), e(pawn), e(pawn), e(pawn),
    null, null, null, null, null, null, null, null, null,
    null, null, null, null, null, null, null, null, null,
    null, null, null, null, null, null, null, null, null,
    f(pawn), f(pawn), f(pawn), f(pawn), f(pawn), f(pawn), f(pawn), f(pawn), f(pawn),
    null, f(bishop), null, null, null, null, null, f(rook), null,
    f(lance), f(knight), f(silver), f(gold), f(king), f(gold), f(silver), f(knight), f(lance),
];
</script>

<template>
    <div id="board" @mousemove="onPieceMove">
        <div class="cell" v-for="piece of board">
            <Piece 
                :kanji="piece?.type.kanji"
                :promoted="piece?.type.promoted" 
                :grabbable="piece?.mine"
                v-if="piece !== null"
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