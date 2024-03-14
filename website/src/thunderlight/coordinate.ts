export class Coordinate {
    constructor(public readonly x: number, public readonly y: number) {}

    public equals(other: Coordinate): boolean {
        return this.x === other.x && this.y === other.y;
    }

    public toString(): string {
        return `(${this.x}, ${this.y})`;
    }

    get absolute(): number {
        return this.x + 9 * this.y;
    }
}