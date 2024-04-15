<script setup lang="ts">
import { EngineMode } from '../thunderlight/engine-mode';
import { jukebox } from '../utils/jukebox';

interface EngineModeInfo { 
    mode: EngineMode,
    name: string, 
    icon: string 
} 

const mode = defineModel<EngineMode>();
const modes: EngineModeInfo[] = [
    {
        mode: 'board',
        name: 'Static board',
        icon: 'albums',
    },
    {
        mode: 'analysis',
        name: 'Analysis',
        icon: 'flash',
    },
    {
        mode: 'play',
        name: 'Play',
        icon: 'game-controller',
    },
];

function switchMode(newMode: EngineMode) {
    mode.value = newMode;
    jukebox.play("mode.switch");
    console.log(newMode);
}

</script>

<template>
<div id="mode-switch">
    <button 
        v-for="buttonMode of modes" 
        class="mode" 
        :class="{selected: mode === buttonMode.mode}" 
        @click="() => switchMode(buttonMode.mode)"
    ><ion-icon :name="buttonMode.icon + (mode === buttonMode.mode ? '' : '-outline')"></ion-icon> {{ buttonMode.name }}</button>
</div>
</template>

<style scoped lang="sass">
#mode-switch
    display: flex
    align-items: center
    justify-content: space-between
    width: 100%
    gap: 0.5em

.mode
    all: unset
    flex: 1
    background: $gray
    color: white
    text-align: center
    padding: 0.5em
    transition: 50ms ease-out
    user-select: none
    display: flex
    align-items: center
    justify-content: center
    gap: 1em

    &.selected
        background: $primary
</style>