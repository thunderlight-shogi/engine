import { addClass, move, resize } from "../dom/dom";
import { Location2D, Vector2D, randomVector } from "../utils/geometry";
import { randomInt, shake } from "../utils/numbers";
import { currentTime, isPast } from "../utils/time";

type ParticleTag = `${string}.${string}`;
type ParticleSet = { tag: ParticleTag, variants: number };

/**
 * The environment affects the particles with gravity and wind.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
class Environment extends Vector2D {
    public static readonly WEIGHTLESSNESS = new Environment(0, 0);
    public static readonly COMMON = new Environment(0.6, 0.025);

    constructor(gravity: number, wind: number) {
        super(wind, gravity);
    }
}

/**
 * The particle draws above all the elements on a screen and falls.
 * To spawn a particle, use a `Firecracker` class.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
class Particle {
	private static readonly FPS: number = 60;
	private static readonly LIFETIME_MS: number = 4000;
    private readonly despawnTime: number;
    private readonly element: HTMLImageElement;
    private readonly location: Location2D;
    private readonly velocity: Vector2D;
    private readonly environment: Environment;
    private interval: NodeJS.Timeout | undefined;

    /**
     * Creates a new particle.
     * 
     * @param element - The HTML element for the particle.
     * @param location - The initial location of the particle on the screen.
     * @param force - The initial force of the particle.
     * @param environment - The environment affecting the particle movement.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    constructor(element: HTMLImageElement, location: Location2D, force: number, environment: Environment) {
        this.interval = undefined;
        this.despawnTime = currentTime() + Particle.LIFETIME_MS;
        this.element = element;
        this.location = location.copy();
        this.velocity = randomVector(force);
        this.environment = environment;
    }

    /**
     * Does the particle need despawning?
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    private needsDespawn(): boolean {
        return isPast(this.despawnTime);
    }

    /**
     * Moves the particle by velocity and environment.
     * Despawns the particle if needed.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    private update() {
        this.velocity.shift(this.environment);
        this.location.shift(this.velocity);

        move(this.element, this.location);

        if (this.needsDespawn()) {
            this.despawn();
        }
    }

    /**
     * Spawns the particle on the screen.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    spawn() {
        document.body.appendChild(this.element);
        this.interval = setInterval(() => this.update(), 1000 / Particle.FPS);
    }

    /**
     * Removes the particle from the screen and stops its updates.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    despawn() {
        this.element.remove();
        clearInterval(this.interval);
    }
}

/**
 * The firecracker helps with particle spawning.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
class Firecracker {
    private particles: Map<ParticleTag, number>;
	private static readonly SIZE: number = 12;
	private static readonly SIZE_SHAKE: number = 12;

    /**
     * Creates a new firecracker.
     * @param particles - The particle set for firecracker to use.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
	constructor(...particles: ParticleSet[]) {
        this.particles = new Map();

        for (const particle of particles) {
            const tag: ParticleTag = particle.tag;
            const variants: number = particle.variants;

            this.particles.set(tag, variants);
        }
    }

    /**
     * Creates a random particle variation image.
     * @param particleTag - A particle tag to create an image by.
     * @returns A particle image.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    private createImageByParticleTag(particleTag: ParticleTag): HTMLImageElement {
        const variations: number | undefined = this.particles.get(particleTag);

        if (variations === undefined) {
            throw Error(`The particle with tag ${particleTag} does not exist.`);
        }

        const variation = randomInt(0, variations);
        const particleImage = new Image();
        const size = shake(Firecracker.SIZE, Firecracker.SIZE_SHAKE);
        particleImage.src = `/particles/${particleTag}.${variation}.svg`;
        resize(particleImage, size);
        addClass(particleImage, "particle");

        return particleImage;
    }

    /**
     * Spawns a particle on the screen.
     * @param particleTag - A particle tag for the spawning particle.
     * @param position - The initial particle position.
     * @param force - The initial force.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
	private spawn(particleTag: ParticleTag, position: Location2D, force = 5): void {
		const particleImage = this.createImageByParticleTag(particleTag);
        const environment = Environment.COMMON;
        const particle = new Particle(particleImage, position, shake(force, force), environment);

		particle.spawn();
	}

    /**
     * Splashes multiple particles on the screen.
     * @param particleTag - A particle tag for the spawning particles.
     * @param position - The initial particles position.
     * @param number - The number of particles to spawn.
     * @param force - The initial force.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
	splash(particleTag: ParticleTag, location: Location2D, count: number = 5, force: number = 5): void {
		for (let i = 0; i < count; i++) {
			this.spawn(particleTag, location, force);
		}
	}
}

export const firecracker: Firecracker = new Firecracker(
    {
        tag: "piece.shred",
        variants: 4
    }
);