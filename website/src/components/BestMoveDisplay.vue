<script setup lang="ts">
import { BestMove } from '../thunderlight/best-move';

const bestMove = defineModel<BestMove>();
</script>

<template>
<div class="best-move">
    <span class="best-move__suggestion">
        {{ bestMove?.toString() }}
    </span>
</div>
</template>

<style scoped lang="sass">
.best-move
    display: flex
    align-items: center
    justify-content: center
    color: white

.best-move__suggestion
    color: white
    text-align: center
</style>