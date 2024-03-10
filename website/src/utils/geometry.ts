import { Comparator } from "./arrays";

export class Location2D {
    constructor(readonly x: number, readonly y: number) {}
}

export function distanceBetween(first: Location2D, second: Location2D): number {
    return Math.sqrt((first.x - second.x) ** 2 + (first.y - second.y) ** 2);
}

export function byDistanceTo(location: Location2D): Comparator<Location2D> {
    return (first, second) => distanceBetween(first, location) - distanceBetween(second, location);
}