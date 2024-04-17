import { MoveType } from "../stores/board-store";
import { BOARD_SIZE, Coordinate } from "../thunderlight/coordinate";
import { Piece, PieceType, PieceTypes } from "../thunderlight/piece-type";
import { Player } from "../thunderlight/player";
import { RestAPI } from "../utils/requests";

const ENGINE_API_DOMAIN = "http://localhost:5174";

export class Move {
    constructor(
        public source: Coordinate,
        public destination: Coordinate,
        public moveType: MoveType,
        public pieceType: PieceType,
    ) {}

    public toString(): string {
        return `${this.source.toString()} --${this.moveType}--> ${this.destination.toString()}`;
    }
}

export class ThunderlightEngine {
    private readonly api: RestAPI;
    private started: boolean;
    private startingPositionId: number;
    private moveTypes: MoveType[] = ["attack+", "travel+", "attack", "travel", "drop", "resign"]

    constructor(startingPositionId: number = 1) {
        this.api = new RestAPI(ENGINE_API_DOMAIN);
        this.started = false;
        this.startingPositionId = startingPositionId;
    }

    public async start(): Promise<void> {
        if(this.started) {
            throw new Error(`The Thunderlight Engine is already started. Please, do not start it more than once.`);
        }

        await this.api.post("start", {
            id: this.startingPositionId,
        });
        this.started = true;
    }

    public async getPieceTypes(): Promise<PieceTypes> {
        this.ensureStarted();

        const pieceTypes = new PieceTypes();
        const response = await this.api.get("piece/list");

        if(!Array.isArray(response)) {
            throw Error(`The API didn't response with a piece list array = ${JSON.stringify(response)}`);
        }

        for(const pieceType of response) {
            pieceTypes.add(pieceType.id, String.fromCharCode(pieceType.kanji));
        }

        for(const firstPieceType of pieceTypes.list) {
            for(const secondPieceType of response) {
                const secondPromotionPieceType = secondPieceType.promote_piece;

                if (!secondPromotionPieceType) {
                    continue;
                }

                const secondPieceTypeKanji = String.fromCharCode(secondPieceType.promote_piece.kanji);
                if (firstPieceType.kanji === secondPieceTypeKanji) {
                    pieceTypes.addDemotion(firstPieceType.id, secondPieceType.id);
                }
            }
        }

        console.log(pieceTypes);
        return pieceTypes;
    }

    public async getStartingPosition(): Promise<(Piece | undefined)[]> {
        const pieces: (Piece | undefined)[] = new Array<Piece | undefined>(81).fill(undefined);
        const pieceTypes = await this.getPieceTypes();
        const startingPosition = await this.api.post("preset/get", {
            id: this.startingPositionId,
        })

        console.log(startingPosition);

        for (const {rank, file, player: playerId, piece_type: { id: pieceTypeId }} of startingPosition.pieces) {
            const player = this.getPlayerName(playerId);
            const pieceType = pieceTypes.find(pieceTypeId);

            pieces[this.getFlatCoordinates(rank, file)] = new Piece(pieceType, player);
        }

        console.log(pieces);
        return pieces;
    } 

    public ensureStarted() {
        if (!this.started) {
            throw new Error(`Just before using the Thunderlight Engine, start it via \`await engine.start()\`.`);
        }
    }

    public async sendMove({source, destination, moveType, pieceType}: Move) {
        console.log({
            old_pos: {
                file: source.x,
                rank: BOARD_SIZE - source.y - 1,
            },

            new_pos: {
                file: destination.x,
                rank: BOARD_SIZE - destination.y, 
            },

            move_type: this.getMoveTypeId(moveType),

            piece_type: {
                id: pieceType.id,
            }
        })

        await this.api.post("move/player", {
            old_pos: {
                file: BOARD_SIZE - source.x - 1,
                rank: source.y,
            },

            new_pos: {
                file: BOARD_SIZE - destination.x - 1,
                rank: destination.y,
            },

            move_type: this.getMoveTypeId(moveType),

            piece_type: {
                id: pieceType.id,
            }
        });
    }

    public async makeBestMove(pieceTypes: PieceTypes) {
        const {
            old_pos: source,
            new_pos: destination,
            move_type: moveTypeId,
            piece_type: {id: pieceTypeId}
        } = await this.api.get("move/engine");

        return new Move(
            new Coordinate(source.file, source.rank),
            new Coordinate(destination.file, destination.rank),
            this.getMoveType(moveTypeId),
            pieceTypes.find(pieceTypeId),
        );
    }

    public async getBestMove(pieceTypes: PieceTypes) {
        const {
            old_pos: source,
            new_pos: destination,
            move_type: moveTypeId,
            piece_type: {id: pieceTypeId}
        } = await this.api.get("move/help");

        return new Move(
            new Coordinate(source.file, source.rank),
            new Coordinate(destination.file, destination.rank),
            this.getMoveType(moveTypeId),
            pieceTypes.find(pieceTypeId),
        );
    }

    private translateRankToX(rank: number) {
        return BOARD_SIZE - rank;
    }

    private translateFileToY(file: number) {
        return file - 1;
    }

    private translateXToRank(x: number) {
        return BOARD_SIZE - x;
    }

    private translateYToFile(y: number) {
        return y + 1;
    }

    private getMoveType(moveId: number): MoveType {
        switch (moveId) {
            case 0:
                return "attack+"

            case 1:
                return "travel+"

            case 2:
                return "attack"

            case 3:
                return "travel"

            case 4:
                return "drop"

            case 5:
                return "resign"

            default:
                throw new Error(`There is no move type defined with id = ${moveId}.`);
        }
    }

    private getMoveTypeId(moveType: MoveType) {
        return this.moveTypes.findIndex(element => element === moveType);
    }

    private getCoordinate(file: number, rank: number): Coordinate {
        return new Coordinate(
            this.translateRankToX(rank),
            this.translateFileToY(file),
        );
    }

    private getFlatCoordinates(file: number, rank: number): number {
        return this.getCoordinate(file, rank).absolute;
    }

    private getPlayerName(playerId: number): Player {
        switch(playerId) {
            case 0: 
                return "sente";

            case 1:
                return "gote";

            default:
                throw new Error(`There is no player with id = ${playerId}.`);
        }
    }
}