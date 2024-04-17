import { Coordinate } from "../thunderlight/coordinate";
import { Piece, PieceTypes } from "../thunderlight/piece-type";
import { Player } from "../thunderlight/player";
import { RestAPI } from "../utils/requests";

const ENGINE_API_DOMAIN = "http://localhost:5174";

export class ThunderlightEngine {
    private readonly api: RestAPI;
    private started: boolean;
    private startingPositionId: number;

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
            pieceTypes.add(pieceType.id, pieceType.kanji);
        }

        for(const firstPieceType of pieceTypes.list) {
            for(const secondPieceType of response) {
                const secondPieceTypeKanji = String.fromCharCode(secondPieceType.promote_piece.kanji);
                if (firstPieceType.kanji === secondPieceTypeKanji) {
                    pieceTypes.addDemotion(firstPieceType.id, secondPieceType.id);
                }
            }
        }

        return pieceTypes;
    }

    public async getStartingPosition(): Promise<(Piece | undefined)[]> {
        const pieces: (Piece | undefined)[] = new Array<Piece | undefined>(81).fill(undefined);
        const pieceTypes = await this.getPieceTypes();
        const startingPosition = await this.api.post("preset/get", {
            id: this.startingPositionId,
        })

        for (const {rank, file, player: playerId, piece_type: { id: pieceTypeId }} of startingPosition.pieces) {
            const player = this.getPlayerName(playerId);
            const pieceType = pieceTypes.find(pieceTypeId);

            pieces[this.getFlatCoordinates(rank, file)] = new Piece(pieceType, player);
        }

        return pieces;
    } 

    public ensureStarted() {
        if (!this.started) {
            throw new Error(`Just before using the Thunderlight Engine, start it via \`await engine.start()\`.`);
        }
    }

    private translateAPICoordinateComponent(coordinateComponent: number) {
        return coordinateComponent - 1;
    }

    private getCoordinate(file: number, rank: number): Coordinate {
        return new Coordinate(
            this.translateAPICoordinateComponent(rank),
            this.translateAPICoordinateComponent(file),
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