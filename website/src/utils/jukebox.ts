import { throws } from "./errors";
import { between, shake } from "./numbers";

type SoundTag = `${string}.${string}`

class Jukebox {
    private audios: Map<SoundTag, HTMLAudioElement>;

    constructor(...soundTags: SoundTag[]) {
        this.audios = new Map(); 
        
        for(const soundTag of soundTags) {
            const audio = new Audio(`/sfx/${soundTag}.mp3`);
            audio.onerror = throws(`The audio ${soundTag} loading failed. It's probably missing on the server.`);
            this.audios.set(soundTag, audio);
        }
    }

    public play(soundTag: SoundTag, pitchShake: number = 0.0, pitchShift: number = 0.0): void {
        const originalAudio = this.getOriginalAudioBySoundTag(soundTag);

        if (originalAudio === undefined) {
            throw new Error(`The sound with tag ${soundTag} does not exist.`);
        }

        if (!between(0, pitchShake, 1)) {
            throw new Error(`The sound pitch range = ${pitchShake} must be between 0 and 1.`);
        }

        const audio = originalAudio.cloneNode() as HTMLAudioElement;

        audio.preservesPitch = false;
        audio.playbackRate = shake(1 + pitchShift, pitchShake);
        audio.play();
    }

    private getOriginalAudioBySoundTag(soundTag: SoundTag): HTMLAudioElement | undefined {
        return this.audios.get(soundTag);
    }
}

export const jukebox: Jukebox = new Jukebox("piece.grab", "piece.drop", "piece.attack", "piece.back");
