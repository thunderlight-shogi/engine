<script setup lang="ts">
import Splash from './Splash.vue';
import { Proverb, getRandomProverb } from '../../thunderlight/proverbs';
import { ref } from 'vue';

const proverb = ref<Proverb>(getRandomProverb());
</script>

<template>
    <Splash class="proverb-splash">
        <div class="proverb-splash__text">
            <div class="proverb-splash__proverb-container">
                <span class="proverb-splash__proverb-number">Proverb #{{ proverb.number }}</span>
                <span class="proverb-splash__proverb-text">{{ proverb.text }}</span>
            </div>

            <div class="proverb-splash__connection-status">
                <div class="proverb-splash__connection-diode"></div>
                <span class="proverb-splash__connection-status-text">Trying to connect to a server THUNDRL#441...</span>
            </div>
        </div>

        <div class="proverb-splash__painting">
            <div class="proverb-splash__painting-filter"></div>
            <div class="proverb-splash__painting-image"></div>
        </div>
    </Splash>
</template>

<style scoped lang="sass">
.proverb-splash
    display: flex
    align-items: center
    justify-content: center

    & > div
        height: 100vh
        box-sizing: border-box

.proverb-splash__text
    flex-basis: 40%
    display: flex
    justify-content: space-between
    flex-direction: column
    padding-block: 4em
    padding-inline-start: 2em

.proverb-splash__painting
    flex-basis: 60%
    overflow: hidden
    position: relative

.proverb-splash__painting-image
    width: 100%
    height: 100%

.proverb-splash__painting-image
    opacity: 0.50
    background-image: url(/images/loading-painting.png)
    background-size: cover

.proverb-splash__painting-filter
    mix-blend-mode: color
    background: white
    position: absolute
    top: 0
    right: 0
    left: 0
    bottom: 0
    z-index: 100000

.proverb-splash__proverb-container
    display: flex
    flex-direction: column
    column-gap: 1em
    font-size: 1.5em

.proverb-splash__proverb-number
    font-size: 0.66em
    color: white
    text-transform: uppercase
    letter-spacing: 0.25em
    opacity: 0.33

.proverb-splash__proverb-text
    font-weight: 600
    color: white

.proverb-splash__connection-status
    display: flex
    align-items: center
    column-gap: 1em

.proverb-splash__connection-diode
    background: #00FF00
    width: 0.5em
    aspect-ratio: 1 / 1
    border-radius: 0.25rem

.proverb-splash__connection-status-text
    color: white
</style>