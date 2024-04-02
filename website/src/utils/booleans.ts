export function flipCoin(chanceOfTrue: number = 0.5): boolean {
    return Math.random() < chanceOfTrue;
}