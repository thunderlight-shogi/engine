import { defineStore } from "pinia";
import { Ref, reactive, ref } from "vue";
import { DEFAULT_BOARD } from "../thunderlight/board";
import { Player, getEnemyOf } from "../thunderlight/player";
import { Coordinate } from "../thunderlight/coordinate";
import { between } from "../utils/numbers";
import { Piece, PieceType, TYPES } from "../thunderlight/piece-type";


export type MoveType = "travel" | "attack" | "back" | "prohibited";
export class Inventory {
    public readonly slots: InventorySlot[];

    constructor(protected readonly owner: Player, types: PieceType[] = TYPES.filter(type => !type.promoted)) {
        this.slots = [];

        for (const type of types) {
            this.slots.push(new InventorySlot(type, 0));
        }
    }

    public isOwnedBy(player: Player): boolean {
        return this.owner === player;
    }

    private getSlotOf(type: PieceType): InventorySlot {
        const slot: InventorySlot | undefined = this.slots.find(slot => slot.type.equals(type.demotion));

        if (!slot) {
            throw new Error(`The inventory slot of ${type.kanji} does not exist.`);
        }

        return slot;
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

export const useBoard = defineStore('board', () => {
    const cells = reactive(DEFAULT_BOARD);
    let turn = ref<Player>('sente');
    let player = ref<Player>('sente');
    let inventories = ref([
        new Inventory('sente'),
        new Inventory('gote')
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
        const inventory = inventories.value.find(inventory => inventory.isOwnedBy(player));

        if (!inventory) {
            throw Error(`There's no inventory for ${player}.`);
        }

        return inventory;
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