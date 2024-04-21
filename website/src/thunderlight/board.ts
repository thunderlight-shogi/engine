import { ThunderlightEngine } from "../api/engine";
import { Inventory, MoveType } from "../stores/board-store";
import { between } from "../utils/numbers";
import { Coordinate } from "./coordinate";
import { Piece, PieceType, PieceTypes } from "./piece-type";
import { Player, getEnemyOf } from "./player";

export class Board {
    cells: (Piece | undefined)[] = [];
    pieceTypes: PieceTypes = new PieceTypes();
    turn: Player = 'sente';
    player: Player = 'sente';
    inventories: Inventory[] = [];
    engine: ThunderlightEngine;

    constructor(engine: ThunderlightEngine) {
        this.engine = engine;
        this.cells = [];
        for (let i = 0; i < 81; i++) {
            this.cells.push(undefined);
        }
        this.turn = 'sente';
        this.player = 'sente';
    }

    async init() {
        await this.engine.start();
        this.pieceTypes = await this.engine.getPieceTypes();
        this.cells = await this.engine.getStartingPosition();
        this.inventories = [
            new Inventory('sente', this.pieceTypes),
            new Inventory('gote', this.pieceTypes)
        ];
    }

    ensureOccupied(coordinate: Coordinate): void {
        if (this.isVacant(coordinate)) {
            throw Error(`There's no piece at ${coordinate}.`);
        }
    }

    isEnemyCamp(coordinate: Coordinate): boolean {
        return (
            (between(6, coordinate.y, 8) && this.turn === 'sente') ||
            (between(0, coordinate.y, 2) && this.turn === 'gote')
        );
    }

    put(piece: Piece | undefined, coordinate: Coordinate): void {
        this.cells.splice(coordinate.absolute, 1, piece);
    }

    pickUp(coordinate: Coordinate): Piece {
        const piece = this.at(coordinate);

        if (piece === undefined) {
            throw Error("Cannot pick up a piece from a vacant cell.");
        }

        this.put(undefined, coordinate);
        return piece;
    }

    at(coordinate: Coordinate): Piece | undefined {
        return this.cells[coordinate.absolute];
    }

    isVacant(coordinate: Coordinate): boolean {
        return this.at(coordinate) === undefined;
    }

    isEnemy(coordinate: Coordinate): boolean {
        this.ensureOccupied(coordinate);
        return this.at(coordinate)?.player !== this.turn;
    }

    getMoveType(source: Coordinate, destination: Coordinate): MoveType {
        this.ensureOccupied(source);

        if (source.equals(destination)) {
            return 'back';
        }

        if (this.isEnemy(source)) {
            return 'prohibited';
        }

        if (this.isVacant(destination)) {
            return 'travel';
        }

        if (this.isEnemy(destination)) {
            return 'attack';
        }

        return 'prohibited';
    }

    move(source: Coordinate, destination: Coordinate, moveType: MoveType = this.getMoveType(source, destination)): MoveType {
        if (moveType === 'attack') {
            const victim = this.at(destination);

            if (victim) {
                this.getInventoryOf(this.turn).add(victim.type);
            }
        }

        if (moveType !== 'prohibited' && moveType !== 'back') {
            const piece = this.pickUp(source);
            this.put(piece, destination);
            this.switchPlayer();

            if (this.isEnemyCamp(destination) && piece.type.promotable && !piece.type.promoted) {
                piece.promote();
            }

            piece.state = 'drop';
        }

        console.info(`Perfomed a move: ${source} ---${moveType}--> ${destination}`);
        return moveType;
    }

    drop(destination: Coordinate, type: PieceType) {
        this.put(new Piece(type, this.turn, 'drop'), destination);
        this.getInventoryOf(this.turn).take(type);
        this.switchPlayer();
    }

    switchPlayer(): void {
        this.turn = getEnemyOf(this.turn);
    }

    getInventoryOf(player: Player): Inventory {
        for (const inventory of this.inventories) {
            if (inventory.isOwnedBy(player)) {
                return inventory;
            }
        }

        throw Error(`There's no inventory for ${player}.`);
    }
}

export default Board;