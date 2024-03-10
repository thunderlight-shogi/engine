<script setup lang="ts">

defineProps<{ 
    kanji: string, 
    promoted: boolean,
    grabbable: boolean,
}>()

</script>

<template>
<div 
    class="piece" 
    :class="{grabbable, promoted}"
    :title="grabbable ? `😈⚔️ Grab this ${kanji} to move!` : `⛔🤚 This is not your ${kanji}!`"
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
        transform: scale(1.00)
    100%
        transform: scale(0.93)

@keyframes piece-promotion
    0%
        transform: rotateY(180deg)
    50%
        transform: rotateY(90deg) scale(1.6)
    100%
        transform: rotateY(0deg)

.piece
    display: grid
    width: 3em
    height: 3em
    place-items: center
    user-select: none
    opacity: 0.5
    cursor: not-allowed
    transform-style: preserve-3d
    &.grabbable
        opacity: 1
        cursor: grab
        &:hover
            animation: piece-hover 500ms ease-in-out infinite alternate


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
    animation: piece-promotion 300ms ease-in-out

    .piece-kanji
        color: $primary

    .piece-body
        stroke: $primary
        stroke-width: 3px
        stroke-linejoin: miter

.piece:active
    scale: 2
    cursor: grabbing
</style>