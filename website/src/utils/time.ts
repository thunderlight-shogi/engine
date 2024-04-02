export function currentTime(): number {
    return Date.now();
}

export function isPast(time: number): boolean {
    return currentTime() > time;
}