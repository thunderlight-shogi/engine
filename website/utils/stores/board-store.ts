import { defineStore } from "pinia";
import { Ref, ref } from "vue";
import { Player, getEnemyOf } from "../thunderlight/player";
import { Coordinate } from "../thunderlight/coordinate";
import { between } from "../utils/numbers";
import { Piece, PieceType, PieceTypes } from "../thunderlight/piece-type";


export type MoveType = "travel+" | "attack+" | "travel" | "attack" | "back" | "prohibited" | "drop" | "resign";
export class Inventory {
    public slots: InventorySlot[];

    constructor(public owner: Player, types: PieceTypes) {
        this.slots = [];

        for (const type of types.list) {
            if (type.promoted) {
                continue;
            }

            this.slots.push(new InventorySlot(type, 0));
        }
    }

    public isOwnedBy(player: Player): boolean {
        return this.owner === player;
    }

    public getSlotOf(type: PieceType): InventorySlot {
        for (const slot of this.slots) {
            if (slot.type.demotion.kanji === type.demotion.kanji) {
                return slot;
            }
        }

        throw new Error(`The inventory slot of ${type.kanji} does not exist.`);
    }

    public add(type: PieceType) {
        this.getSlotOf(type).add();
    }

    public take(type: PieceType) {
        this.getSlotOf(type).take();
    }
}

export class InventorySlot {
    constructor(public readonly type: PieceType,
                public count: number = 0) {}

    get empty(): boolean {
        return this.count === 0;
    }

    public take(): void {
        if (this.empty) {
            throw Error(`Can't take a ${this.type.kanji} from an empty inventory.`);
        }

        this.count -= 1;
    }

    public add(): void {
        this.count += 1;
    }
}

export const useBoard = defineStore('board', async () => {
    const cells = ref() as Ref<(Piece | undefined)[]>;
    const pieceTypes = ref() as Ref<PieceTypes>;
    let turn = ref<Player>('sente');
    let player = ref<Player>('sente');
    let inventories = ref([
        new Inventory('sente', pieceTypes.value),
        new Inventory('gote', pieceTypes.value)
    ]) as Ref<Inventory[]>;

    function ensureOccupied(coordinate: Coordinate) {
        if (isVacant(coordinate)) {
            throw Error(`There's no piece at ${coordinate}.`);
        }
    }

    function isEnemyCamp(coordinate: Coordinate): boolean {
        return between(6, coordinate.y, 8) && turn.value === 'sente' ||
               between(0, coordinate.y, 2) && turn.value === 'gote';
    }

    function put(piece: Piece | undefined, coordinate: Coordinate): void {
        cells.value.splice(coordinate.absolute, 1, piece);
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
        return cells.value[coordinate.absolute];
    }

    function isVacant(coordinate: Coordinate): boolean {
        return at(coordinate) === undefined;
    }

    function isEnemy(coordinate: Coordinate): boolean {
        ensureOccupied(coordinate);
        return at(coordinate)?.player !== turn.value;
    }

    function getMoveType(source: Coordinate, destination: Coordinate): MoveType {
        ensureOccupied(source);

        if (source.equals(destination)) {
            return 'back';
        }

        if (isEnemy(source)) {
            return 'prohibited';
        }

        if (isVacant(destination)) {
            return 'travel';
        }

        if (isEnemy(destination)) {
            return 'attack';
        }

        return 'prohibited';
    }

    function move(source: Coordinate, destination: Coordinate): MoveType {
        const moveType = getMoveType(source, destination);

        if (moveType === 'attack') {
            const victim = at(destination);

            if (victim) {
                getInventoryOf(turn.value).add(victim.type);
            }
        }

        if (moveType !== 'prohibited' && moveType !== 'back') {
            const piece = pickUp(source);
            put(piece, destination);
            switchPlayer();

            if (isEnemyCamp(destination) && piece.type.promotable && !piece.type.promoted) {
                piece.promote();
            }

            piece.state = 'drop';
        }

        console.info(`Perfomed a move: ${source} ---${moveType}--> ${destination}`);
        return moveType;
    }

    function drop(destination: Coordinate, type: PieceType) {
        put(new Piece(type, turn.value, 'drop'), destination);
        getInventoryOf(turn.value).take(type);
        switchPlayer();
    }

    function switchPlayer(): void {
        turn.value = getEnemyOf(turn.value);
    }

    function getInventoryOf(player: Player): Inventory {
        for (const inventory of inventories.value) {
            if (inventory.isOwnedBy(player)) {
                return inventory;
            }
        }

        throw Error(`There's no inventory for ${player}.`);
    }

    return {
        player,
        turn,
        cells,

        getInventoryOf,
        move,
        drop
    };
});