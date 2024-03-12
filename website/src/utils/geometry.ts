import { Comparator } from "./arrays";
import { random, randomUpTo } from "./numbers";

export class Vector2D {
    constructor(public x: number, public y: number) {}

    public copy(): Vector2D {
        return new Vector2D(this.x, this.y);
    }

    public rotatedBy(radians: number): Vector2D {
        const sin = Math.sin(radians);
        const cos = Math.cos(radians);
        const x = cos * this.x - sin * this.y;
        const y = sin * this.x + cos * this.y;

        return new Vector2D(x, y);
    }

    public shorten(divisor: number): Vector2D {
        if (divisor === 0) {
            throw new Error("Cannot shorten vector by divisor of zero.");
        }

        this.x /= divisor;
        this.y /= divisor;

        return this;
    }

    public shift(term: Vector2D): Vector2D {
        this.x += term.x;
        this.y += term.y;

        return this;
    }

    public toString(): string {
        return `Vector2D(${this.x}, ${this.y})`;
    }

    get angle(): number {
        return Math.atan2(this.y, this.x);
    }
}

export function randomVector(length: number) {
    return new Vector2D(0, length).rotatedBy(randomAngle());
}

export class Location2D {
    constructor(public x: number, public y: number) {}

    public copy(): Location2D {
        return new Location2D(this.x, this.y);
    }

    public shift(shift: Vector2D): Location2D {
        this.x += shift.x;
        this.y += shift.y;

        return this;
    }

    public toString(): string {
        return `Location2D(${this.x}, ${this.y})`;
    }
}

function randomAngle(): number {
    return random(-Math.PI, Math.PI);
}

export function centerLocation(): Location2D {
    return new Location2D(0, 0);
}

export function randomLocation(): Location2D {
    const screenWidth: number = window.innerWidth;
    const screenHeight: number = window.innerHeight;
    const x = randomUpTo(screenWidth);
    const y = randomUpTo(screenHeight);

    return new Location2D(x, y);
}

export function distanceBetween(first: Location2D, second: Location2D): number {
    return Math.sqrt((first.x - second.x) ** 2 + (first.y - second.y) ** 2);
}

export function byDistanceTo(location: Location2D): Comparator<Location2D> {
    return (first, second) => distanceBetween(first, location) - distanceBetween(second, location);
}

export function angleBetween(first: Location2D, second: Location2D): number {
    return Math.atan2(second.y - first.y, second.x - first.x);
}