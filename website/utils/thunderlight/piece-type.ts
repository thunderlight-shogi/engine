import { Player } from "./player";

export type PieceState = "idle" | "drop";

export class PieceTypes {
    public readonly list: PieceType[];

    constructor() {
        this.list = [];
    }

    public add(id: number, kanji: string): void {
        this.list.push(new PieceType(id, kanji, this));
    }

    public find(id: number): PieceType {
        const pieceType: PieceType | undefined = this.list.find(pieceType => pieceType.id === id);

        if (pieceType === undefined) {
            throw new Error(`There is no piece type with id = ${id}.`);
        }

        return pieceType;
    }

    public addDemotion(pieceId: number, demotionId: number): void {
        this.find(pieceId)._demotion = this.find(demotionId);
    }
}

export class PieceType {
    constructor(public readonly id: number,
                public readonly kanji: string, 
                private readonly pieceTypes: PieceTypes,
                public _demotion: PieceType | undefined = undefined) {}

    get demotion(): PieceType {
        return this._demotion ?? this;
    }

    get promotions(): PieceType[] {
        return this.pieceTypes.list.filter(type => this.equals(type._demotion));
    }

    get promotable(): boolean {
        return this.promotions.length === 1;
    }

    get promotion(): PieceType {
        if (!this.promotable) {
            throw Error(`To promote a piece, there should be exactly one promotion type, not ${this.promotions.length}.`)
        }

        return this.promotions[0];
    }

    get promoted(): boolean {
        return this._demotion !== undefined;
    }

    public equals(other: PieceType | undefined): boolean {
        return this.kanji === other?.kanji;
    }

    public toString(): string {
        return `${this.kanji} ${this}`
    }
}

export class Piece {
    constructor(public type: PieceType,
                public readonly player: Player,
                public state: PieceState = "idle") {}

    public promote(): void {
        if (this.type.promoted) {
            throw Error(`The ${this.type.kanji} is unpromotable.`);
        }

        this.type = this.type.promotion;
    }
}