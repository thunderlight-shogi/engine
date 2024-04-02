<script setup lang="ts">
import { Ref, ref } from 'vue';
import { shake } from '../utils/numbers';
import { PieceState } from '../thunderlight/piece-type';

const emit = defineEmits([
    'grab',
    'drop',
]);

const props = defineProps<{ 
    kanji: string,
    grabbable: boolean,
    enemy: boolean,
    promoted: boolean,
    state: PieceState
}>();

const piece: Ref<HTMLElement | null> = ref(null);

function onPieceGrab(event: MouseEvent): void {
    // if (!props.grabbable) {
    //     return;
    // }

    if (!piece.value) {
        return;
    }

    // const element: HTMLElement = event.currentTarget as HTMLElement;
    piece.value.style.transition = '';
    piece.value.style.zIndex = '99';

    emit("grab", piece.value);
}

function onPieceDrop(event: MouseEvent) {
    // if (!props.grabbable) {
    //     return;
    // }

    // const element: HTMLElement = event.currentTarget as HTMLElement;

    if (!piece.value) {
        return;
    }

    piece.value.style.zIndex = '0';
    piece.value.style.transition = '200ms ease-in-out';
    piece.value.style.transform = `rotate(${shake(0, 15) + (props.enemy ? 180 : 0)}deg)`;

    emit("drop", piece.value);
}
</script>

<template>
<div 
    class="piece" 
    ref="piece"
    :class="{grabbable, promoted, enemy, [state]: true }"
    :title="grabbable ? `😈⚔️ Grab this ${kanji} to move!` : `⛔🤚 This is not your ${kanji}!`"
    @mousedown="onPieceGrab"
    @mouseup="onPieceDrop"
>
    <svg class="piece-svg" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 59.97 62.8">
        <polygon class="piece-body" points="54.67 14.25 29.99 0 29.99 0 29.99 0 5.3 14.25 0 62.8 29.99 62.8 59.97 62.8 54.67 14.25"/>
    </svg>

    <span class="piece-kanji">{{ kanji }}</span>
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
    filter: hue-rotate(0deg)
    
    &.drop
        animation: piece-drop 200ms ease-in

    &.enemy
        transform: rotate(180deg)

    &.grabbable
        opacity: 1
        cursor: grab
        filter: hue-rotate(0)

        &:active
            scale: 1.5
            cursor: grabbing
            filter: drop-shadow(10px 10px 5px transparentize($background, 0.60))

    &:not(.grabbable)
        opacity: 0.3

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