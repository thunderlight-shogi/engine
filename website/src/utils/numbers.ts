import { flipCoin } from "./booleans";

function randomSign(): number {
    return flipCoin() ? 1 : -1;
}

export function randomUpTo(threshold: number): number {
    return random(0, threshold);
}

export function random(lowerThreshold: number, upperThreshold: number): number {
    return Math.random() * (upperThreshold - lowerThreshold) + lowerThreshold;
}

export function randomInt(lowerThreshold: number, upperThreshold: number): number {
    return Math.floor(random(lowerThreshold, upperThreshold));
}

export function between(min: number, number: number, max: number): boolean {
    return min <= number && number <= max;
}

export function shake(number: number, shakeAmount: number): number {
    return number + randomSign() * randomUpTo(shakeAmount);
}

export function distance(first: number, second: number) {
    return Math.abs(first - second);
}

export function mean(first: number, second: number): number {
    return (first + second) / 2;
}

export function translate(translationMin: number, translationMax: number, number: number, originalMin: number, originalMax: number) {
    const originalRange = distance(originalMin, originalMax);
    const translationRange = distance(translationMin, translationMax);
    const numberRate = number / originalRange;

    return numberRate * translationRange + translationMin;
}