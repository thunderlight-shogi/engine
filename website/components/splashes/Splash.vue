<script setup lang="ts">
</script>

<template>
<div class="splash">
    <slot></slot>
</div>
</template>

<style scoped lang="sass">
.splash
    background: $background
    position: absolute
    top: 0
    left: 0
    right: 0
    bottom: 0
    width: 100%
    height: 100%
    z-index: 1000
</style>