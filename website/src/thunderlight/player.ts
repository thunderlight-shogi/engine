export type Player = "sente" | "gote";

export function getEnemyOf(player: Player): Player {
    switch (player) {
        case 'gote':
            return 'sente';

        case 'sente':
            return 'gote';
    }
}