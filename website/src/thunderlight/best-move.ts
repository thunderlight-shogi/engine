import { MoveType } from "../stores/board-store";
import { Coordinate } from "./coordinate";
import { PieceType } from "./piece-type";

export class BestMove {
    constructor(
        public pending: boolean,
        public sourcePosition: Coordinate,
        public destinationPosition: Coordinate,
        public moveType: MoveType,
        public pieceType: PieceType,
    ) {
    }

    public toString(): string {
        if (this.pending) {
            return `The engine is trying to figure out the best move. Please, wait...`;
        }

        return `You need to ${this.moveType} your ${this.pieceType.kanji} from ${this.sourcePosition.toString()} to ${this.destinationPosition.toString()}.`;
    }
}