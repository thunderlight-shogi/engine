import { BISHOP, GOLD, KING, KNIGHT, LANCE, PAWN, Piece, ROOK, SILVER } from "./piece-type";

export const DEFAULT_BOARD: (Piece | undefined)[] = [
    new Piece(LANCE, 'gote'), new Piece(KNIGHT, 'gote'), new Piece(SILVER, 'gote'), new Piece(GOLD, 'gote'), new Piece(KING, 'gote'), new Piece(GOLD, 'gote'), new Piece(SILVER, 'gote'), new Piece(KNIGHT, 'gote'), new Piece(LANCE, 'gote'),
    undefined, new Piece(ROOK, 'gote'), undefined, undefined, undefined, undefined, undefined, new Piece(BISHOP, 'gote'), undefined,
    new Piece(PAWN, 'gote'), new Piece(PAWN, 'gote'), new Piece(PAWN, 'gote'), new Piece(PAWN, 'gote'), new Piece(PAWN, 'gote'), new Piece(PAWN, 'gote'), new Piece(PAWN, 'gote'), new Piece(PAWN, 'gote'), new Piece(PAWN, 'gote'),
    undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined,
    undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined,
    undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined, undefined,
    new Piece(PAWN, 'sente'), new Piece(PAWN, 'sente'), new Piece(PAWN, 'sente'), new Piece(PAWN, 'sente'), new Piece(PAWN, 'sente'), new Piece(PAWN, 'sente'), new Piece(PAWN, 'sente'), new Piece(PAWN, 'sente'), new Piece(PAWN, 'sente'),
    undefined, new Piece(BISHOP, 'sente'), undefined, undefined, undefined, undefined, undefined, new Piece(ROOK, 'sente'), undefined,
    new Piece(LANCE, 'sente'), new Piece(KNIGHT, 'sente'), new Piece(SILVER, 'sente'), new Piece(GOLD, 'sente'), new Piece(KING, 'sente'), new Piece(GOLD, 'sente'), new Piece(SILVER, 'sente'), new Piece(KNIGHT, 'sente'), new Piece(LANCE, 'sente'),
];