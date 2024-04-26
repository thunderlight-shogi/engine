<script setup lang="ts">
const metric = defineModel<string>();
</script>

<template>
    <pre class="metrics-display">{{ metric?.trim() }}</pre>
</template>

<style scoped lang="sass">
.metrics-display
    color: $primary
    font-family: "Fira Mono", "DejaVu Sans Mono", Menlo, Consolas, "Liberation Mono", Monaco, "Lucida Console", monospace
    height: 100%
    width: 12em
</style>