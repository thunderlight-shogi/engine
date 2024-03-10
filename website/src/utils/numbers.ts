import { flipCoin } from "./booleans";

function randomSign(): number {
    return flipCoin() ? 1 : -1;
}

function randomUpTo(threshold: number): number {
    return Math.random() * threshold;
}

export function between(min: number, number: number, max: number): boolean {
    return min <= number && number <= max;
}

export function shake(number: number, shakeAmount: number): number {
    return number + randomSign() * randomUpTo(shakeAmount);
}

export function mean(first: number, second: number): number {
    return (first + second) / 2;
}