<script setup lang="ts">
defineProps<{ 
    viewbox: string,
}>()
</script>

<template>
    <svg class="icon" :viewBox="viewbox">
        <slot></slot>
    </svg>
</template>

<style scoped lang="sass">
.icon
    width: 1em
    aspect-ratio: 1 / 1
</style>