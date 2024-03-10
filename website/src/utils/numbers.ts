import { flipCoin } from "./booleans";

function randomSign(): number {
    return flipCoin() ? 1 : -1;
}

function randomUpTo(threshold: number) {
    return Math.random() * threshold;
}

export function between(min: number, number: number, max: number): boolean {
    return min <= number && number <= max;
}

export function shake(number: number, shakeAmount: number) {
    return number + randomSign() * randomUpTo(shakeAmount);
}