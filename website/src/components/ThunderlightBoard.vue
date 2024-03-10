<script setup lang="ts">

</script>

<template>
    <div id="board">
        <div class="cell" v-for="_ in 9 ** 2">
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