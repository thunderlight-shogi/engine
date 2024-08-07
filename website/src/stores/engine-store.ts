import { ref, Ref } from "vue";
import { defineStore } from "pinia";
import { Coordinate } from "../thunderlight/coordinate";
import { Piece, PieceTypes } from "../thunderlight/piece-type";
import { Player } from "../thunderlight/player";
import { RestAPI } from "../utils/requests";

const ENGINE_API_DOMAIN = "http://localhost:5174";

export const useEngine = defineStore("engine", () => {
  const api = new RestAPI(ENGINE_API_DOMAIN);
  const started = ref(false);
  const startingPositionId = ref(1);
  const pieces = ref(new Array<Piece | undefined>(81).fill(undefined)) as Ref<(Piece | undefined)[]>;
  const pieceTypes = new PieceTypes();

  async function start() {
    if (started.value) {
      throw new Error(`The Thunderlight Engine is already started. Please, do not start it more than once.`);
    }

    await api.post("start", { id: startingPositionId.value });
    started.value = true;
  }

  async function getPieceTypes() {
    ensureStarted();

    const response = await api.get("piece/list");

    if (!Array.isArray(response)) {
      throw Error(`The API didn't respond with a piece list array = ${JSON.stringify(response)}`);
    }

    for (const pieceType of response) {
      pieceTypes.add(pieceType.id, String.fromCharCode(pieceType.kanji));
    }

    for (const firstPieceType of pieceTypes.list) {
      for (const secondPieceType of response) {
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

    return pieceTypes;
  }

  async function getStartingPosition() {
    await getPieceTypes();

    const startingPosition = await api.post("preset/get", { id: startingPositionId.value });

    for (const { rank, file, player: playerId, piece_type: { id: pieceTypeId } } of startingPosition.pieces) {
      const player = getPlayerName(playerId);
      const pieceType = pieceTypes.find(pieceTypeId);

      pieces.value[getFlatCoordinates(rank, file)] = new Piece(pieceType, player);
    }
  }

  function ensureStarted() {
    if (!started.value) {
      throw new Error(`Just before using the Thunderlight Engine, start it via \`await engine.start()\`.`);
    }
  }

  function translateAPICoordinateComponent(coordinateComponent: number) {
    return coordinateComponent - 1;
  }

  function getCoordinate(file: number, rank: number): Coordinate {
    return new Coordinate(
      translateAPICoordinateComponent(rank),
      translateAPICoordinateComponent(file),
    );
  }

  function getFlatCoordinates(file: number, rank: number): number {
    return getCoordinate(file, rank).absolute;
  }

  function getPlayerName(playerId: number): Player {
    switch (playerId) {
      case 0:
        return "sente";
      case 1:
        return "gote";
      default:
        throw new Error(`There is no player with id = ${playerId}.`);
    }
  }

  return {
    api,
    started,
    startingPositionId,
    pieces,
    pieceTypes,
    start,
    getPieceTypes,
    getStartingPosition,
    ensureStarted,
    translateAPICoordinateComponent,
    getCoordinate,
    getFlatCoordinates,
    getPlayerName
  };
});
