<script setup lang="ts">
import { locateMouse } from '../dom/dom';
import { shake } from '../utils/numbers';

const emit = defineEmits([
    'grab',
    'drop',
]);

const props = defineProps<{ 
    kanji: string, 
    promoted: boolean,
    grabbable: boolean,
}>();

function onPieceGrab(event: MouseEvent): void {
    // if (!props.grabbable) {
    //     return;
    // }

    const element: HTMLElement = event.currentTarget as HTMLElement;
    element.style.transition = '';
    element.style.zIndex = '99';

    emit("grab", element, locateMouse(event));
}

function onPieceDrop(event: MouseEvent) {
    // if (!props.grabbable) {
    //     return;
    // }

    const element: HTMLElement = event.currentTarget as HTMLElement;

    element.style.zIndex = '0';
    element.style.transition = '200ms ease-in-out';
    element.style.transform = `rotate(${shake(0, 15) + (props.grabbable ? 0 : 180)}deg)`;

    emit("drop", element, locateMouse(event));
}
</script>

<template>
<div 
    class="piece" 
    :class="{grabbable, promoted}"
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
    position: absolute
    display: grid
    width: 3em
    height: 3em
    place-items: center
    user-select: none
    filter: hue-rotate(180deg)
    opacity: 0.7
    cursor: not-allowed
    transform-style: preserve-3d
    transform: rotate(180deg)
    animation: piece-drop 200ms ease-in
    transition: 75ms
    &.grabbable
        opacity: 1
        cursor: grab
        transform: rotate(0) 
        filter: grayscale(0)

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