<script setup lang="ts">
import Icon from '../icons/Icon.vue';
import Splash from './Splash.vue';
import { sleep } from '../../utils/sleep';
import { ref, onMounted } from 'vue';

const DISPLAY_TIME_MILLISECONDS = 4450 * 0 + 6000;
const isDisplayed = ref<boolean>(true);

onMounted(async () => {
    await sleep(DISPLAY_TIME_MILLISECONDS);
    isDisplayed.value = false;
})

</script>

<template>
<Splash class="loading-splash" :class="{ 'hidden': !isDisplayed }">
    <div class="loading-splash__logo">
        <Icon viewbox="0 0 176.05 302.4">
            <path class="loading-splash__logo-image" d="M173.37,74.88l-94.3,77.24-2.55-52.76,88.5,110.13,11.03,13.72-9.51,8.45-79.56,70.72,54.73-91.3,1.52,22.18C124.29,220.61,16.3,148.32,0,137.42c0,0,23.27-35.48,23.27-35.48L90.11,0l83.26,74.88h0Z"/>
        </Icon>
        <span class="loading-splash__logo-text">Thunderlight</span>
    </div>
</Splash>
</template>

<style scoped lang="sass">
@keyframes loading-splash__fade-out
    0%
        opacity: 1

    100%
        opacity: 0

.loading-splash
    display: flex
    align-items: center
    justify-content: center
    user-select: none
    animation: loading-splash__fade-out 300ms ease-out 4100ms forwards

    &.hidden
        display: none

.loading-splash__logo
    display: flex
    align-items: center
    justify-content: center
    column-gap: 1em
    font-size: 4em
    filter: drop-shadow(0 0 0.4em #00000066)

.loading-splash__logo-image
    fill: white
    width: 1em


@keyframes loading-splash__logo-gradient
    0%
        background-position: 0% 50%
    100%
        background-position: 100% 50%

.loading-splash__logo-text
    text-transform: uppercase
    font-weight: 1000
    font-family: "Golos Text", sans-serif
    color: white
    background: linear-gradient(135deg, $primary 0%, $primary 48%, white 49%, white 51%, $primary 52%, $primary 100%)
    background-size: 400% 400%
    animation: loading-splash__logo-gradient 3s ease-in-out infinite
    -webkit-background-clip: text
    -webkit-text-fill-color: transparent
    -webkit-box-decoration-break: clone
    box-decoration-break: clone 


@keyframes loading-splash__logo-idle
    0%
        scale: 0.97
        opacity: 0.90
    100%
        scale: 1.03
        opacity: 1

@keyframes loading-splash__logo-fade-in
    0%
        scale: 1
        opacity: 0
    100%
        scale: 0.97
        opacity: 0.95

@keyframes loading-splash__logo-fade-out
    0%
        scale: 0.97
        opacity: 0.95
    90%
        scale: 1.05
        opacity: 1
    100%
        scale: 0.8
        opacity: 0

.loading-splash__logo
    animation: loading-splash__logo-fade-in 200ms ease-out 500ms, loading-splash__logo-idle 1.5s ease-in-out 700ms 2 alternate, loading-splash__logo-fade-out 500ms forwards 3700ms ease-in
    opacity: 0

</style>