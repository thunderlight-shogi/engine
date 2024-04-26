import { throws } from "./errors";
import { between, shake } from "./numbers";

type SoundTag = `${string}.${string}`

/** 
 * The Jukebox class is designed to manage and play audio files 
 * with specified sound tags, allowing for pitch manipulation during playback. 
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
class Jukebox {
    private audios: Map<SoundTag, HTMLAudioElement>;

    /**
     * Creates a new jukebox that can play sounds by tags provided.
     * @param soundTags - The array of sound tags present to play by this jukebox.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    constructor(...soundTags: SoundTag[]) {
        this.audios = new Map(); 
        
        for(const soundTag of soundTags) {
            const audio = new Audio(`/sfx/${soundTag}.mp3`);
            audio.onerror = throws(`The audio ${soundTag} loading failed. It's probably missing on the server.`);
            this.audios.set(soundTag, audio);
        }
    }

    /**
     * The function `play` plays a sound with optional pitch modifications based on the provided
     * parameters.
     * @param {SoundTag} soundTag - The `soundTag` parameter is used to specify the tag of the sound
     * that you want to play. It is of type `SoundTag`.
     * @param {number} [pitchShake=0.0] - The `pitchShake` parameter in the `play` function is used to
     * control the pitch variation of the sound being played. It should be a number between 0 and 1,
     * where 0 means no pitch variation and 1 means maximum pitch variation.
     * @param {number} [pitchShift=0.0] - The `pitchShift` parameter in the `play` function allows you
     * to adjust the pitch of the audio being played. A positive value will increase the pitch, while a
     * negative value will decrease it.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
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

    /**
     * The function `getOriginalAudioBySoundTag` returns the original HTML audio element associated
     * with a given sound tag, or undefined if not found.
     * @param {SoundTag} soundTag - The `soundTag` parameter is a unique identifier or key used to
     * retrieve an HTMLAudioElement from a Map called `audios`.
     * @returns An `HTMLAudioElement` or `undefined` is being returned.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    private getOriginalAudioBySoundTag(soundTag: SoundTag): HTMLAudioElement | undefined {
        return this.audios.get(soundTag);
    }
}

export const jukebox: Jukebox = new Jukebox(
    "piece.grab", 
    "piece.drop", 
    "piece.attack", 
    "piece.back", 
    "piece.prohibited",
    
    "mode.switch",
);
