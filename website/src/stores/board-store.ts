import { defineStore } from "pinia";
import { reactive } from "vue";
import { DEFAULT_BOARD } from "../thunderlight/board";
import { Player } from "../thunderlight/player";
import { Coordinate } from "../thunderlight/coordinate";
import { between } from "../utils/numbers";
import { Piece } from "../thunderlight/piece-type";


export enum MoveType {
    TRAVEL, ATTACK, BACK, PROHIBITED
}

export const useBoard = defineStore('board', () => {
    const cells = reactive(DEFAULT_BOARD);
    let turn: Player = 'sente';
    let player: Player = 'sente';

    function ensureOccupied(coordinate: Coordinate) {
        if (isVacant(coordinate)) {
            throw Error(`There's no piece at ${coordinate}.`);
        }
    }

    function isEnemyCamp(coordinate: Coordinate): boolean {
        return between(6, coordinate.y, 8) && player === 'sente' ||
               between(0, coordinate.y, 2) && player === 'gote';
    }

    function put(piece: Piece | undefined, coordinate: Coordinate): void {
        cells.splice(coordinate.absolute, 1, piece);
    }

    function pickUp(coordinate: Coordinate): Piece {
        const piece = at(coordinate);

        if(piece === undefined) {
            throw Error("Cannot pick up a piece from a vacant cell.");
        }

        put(undefined, coordinate);

        return piece;
    }

    function at(coordinate: Coordinate): Piece | undefined {
        return cells[coordinate.absolute];
    }

    function isVacant(coordinate: Coordinate): boolean {
        return at(coordinate) === undefined;
    }

    function isFriend(coordinate: Coordinate): boolean {
        ensureOccupied(coordinate);
        return at(coordinate)?.player === player;
    }

    function isEnemy(coordinate: Coordinate): boolean {
        ensureOccupied(coordinate);
        return at(coordinate)?.player !== player;
    }

    function getMoveType(source: Coordinate, destination: Coordinate): MoveType {
        ensureOccupied(source);

        if(source.equals(destination)) {
            return MoveType.BACK;
        }

        if(isEnemy(source)) {
            return MoveType.PROHIBITED;
        }

        if(isVacant(destination)) {
            return MoveType.TRAVEL;
        }

        if(isEnemy(destination)) {
            return MoveType.ATTACK;
        }

        return MoveType.PROHIBITED;
    }

    function move(source: Coordinate, destination: Coordinate): MoveType {
        const moveType = getMoveType(source, destination);

        if(moveType !== MoveType.PROHIBITED && moveType !== MoveType.BACK) {
            const piece = pickUp(source);
            put(piece, destination);
            switchPlayer();

            if (isEnemyCamp(destination) && !piece.type.promoted) {
                piece.promote();
            }

            piece.state = 'drop';
        }

        return moveType;
    }

    function switchPlayer(): void {
        switch (player) {
            case 'gote':
                player = 'sente';
                break;

            case 'sente':
                player = 'gote';
                break;
        }
    }

    return {
        player,
        turn,
        cells,

        move
    };
});