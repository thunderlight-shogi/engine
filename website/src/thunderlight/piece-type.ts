import { generateUUIDv4 } from "../crypto/uuids";
import { Player } from "./player";

export type PieceState = "idle" | "drop";

export class PieceType {
    constructor(public readonly kanji: string, 
                public readonly source: PieceType | undefined = undefined) {}

    get promoted(): boolean {
        return this.source !== undefined;
    }

    public equals(other: PieceType | undefined): boolean {
        return this.kanji === other?.kanji;
    }
}

export class Piece {
    public readonly id: string = generateUUIDv4();

    constructor(public type: PieceType,
                public readonly player: Player,
                public state: PieceState = "idle") {}

    public promote(): void {
        if (this.type.promoted) {
            throw Error(`The ${this.type.kanji} is unpromotable.`);
        }

        const promotionTypes = TYPES.filter(type => this.type.equals(type.source)); 

        if (promotionTypes.length !== 1) {
            throw Error(`To promote a piece, there should be exactly one promotion type, not ${promotionTypes.length}.`)
        }

        this.type = promotionTypes[0];
    }
}

export const PAWN = new PieceType('歩');
export const KING = new PieceType('王');
export const ROOK = new PieceType('飛');
export const DRAGON = new PieceType('龍', ROOK);
export const BISHOP = new PieceType('角');
export const HORSE = new PieceType('馬', BISHOP);
export const GOLD = new PieceType('金');
export const SILVER = new PieceType('銀');
export const GOLDEN_SILVER = new PieceType('全', SILVER);
export const KNIGHT = new PieceType('桂');
export const GOLDEN_KNIGHT = new PieceType('今', KNIGHT);
export const LANCE = new PieceType('香');
export const GOLDEN_LANCE = new PieceType('杏', LANCE);
export const TOKIN = new PieceType('と', PAWN);

export const TYPES = [
    PAWN, KING, ROOK, DRAGON,
    BISHOP, HORSE, GOLD, SILVER,
    GOLDEN_SILVER, KNIGHT,
    GOLDEN_KNIGHT, LANCE,
    GOLDEN_LANCE, TOKIN
];