<script setup lang="ts">
import { Ref, ref } from 'vue';
import { shake } from '../utils/numbers';
import { PieceState, PieceType } from '../thunderlight/piece-type';

const emit = defineEmits([
    'grab',
    'drop',
]);

const props = defineProps<{ 
    type: PieceType,
    grabbable: boolean,
    enemy: boolean,
    state: PieceState
}>();

const piece: Ref<HTMLElement | null> = ref(null);

function onPieceGrab(_: MouseEvent): void {
    if (!piece.value) {
        return;
    }

    piece.value.style.transition = '';
    piece.value.style.zIndex = '99';

    emit("grab", piece.value, props.type);
}

function onPieceDrop(_: MouseEvent) {
    if (!piece.value) {
        return;
    }

    piece.value.style.zIndex = '0';
    piece.value.style.transition = '200ms ease-in-out';
    piece.value.style.transform = `rotate(${shake(0, 15) + (props.enemy ? 180 : 0)}deg)`;

    emit("drop", piece.value, props.type);
}
</script>

<template>
<div 
    class="piece" 
    ref="piece"
    :class="{grabbable, promoted: type.promoted, enemy, [state]: true }"
    :title="enemy ? `That's not this player's turn!` : `Grab this piece to make a move.`"
    @mousedown="onPieceGrab"
    @mouseup="onPieceDrop"
>
    <svg class="piece-svg" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 59.97 62.8">
        <polygon class="piece-body" points="54.67 14.25 29.99 0 29.99 0 29.99 0 5.3 14.25 0 62.8 29.99 62.8 59.97 62.8 54.67 14.25"/>
    </svg>

    <span class="piece-kanji">{{ type.kanji }}</span>
</div>

</template>

<style scoped lang="sass">
@keyframes piece-hover
    0%
        scale: 1.00
    100%
        scale: 0.93

@keyframes piece-promotion
    0%
        transform: rotateY(180deg)
    50%
        transform: rotateY(90deg) scale(1.6)
    100%
        transform: rotateY(0deg)

@keyframes piece-drop
    0%
        scale: 1.50
    50%
        scale: 3.00
        filter: none
    60%
        scale: 0.90
        filter: drop-shadow(0 0 15px white) drop-shadow(0 0 45px white) drop-shadow(0 0 100px white)
    100%
        scale: 1.00
        filter: none

.piece
    display: grid
    width: 3em
    height: 3em
    place-items: center
    user-select: none
    cursor: not-allowed
    transform-style: preserve-3d
    transition: 75ms
    opacity: 0.3
    
    &.drop
        // animation: piece-drop 200ms ease-in

    &.enemy
        transform: rotate(180deg)

    &.grabbable
        opacity: 1
        cursor: grab
        filter: hue-rotate(0deg)

        &:active
            scale: 1.5
            cursor: grabbing
            filter: drop-shadow(10px 10px 5px transparentize($background, 0.60))

    & > *
        grid-row: 1
        grid-column: 1

.piece-body
    fill: #f3efe0
    stroke-width: 0px

.piece-kanji
    color: $background
    font-family: "Stick", sans-serif
    font-weight: 1000
    font-size: 2em
    z-index: 1

.piece.promoted
    .piece-kanji
        color: $primary

    .piece-body
        stroke: $primary
        stroke-width: 3px
        stroke-linejoin: miter
</style>