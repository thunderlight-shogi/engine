import { addClass, move, resize } from "../dom/dom";
import { Location2D, Vector2D, randomVector } from "../utils/geometry";
import { randomInt, shake } from "../utils/numbers";
import { currentTime, isPast } from "../utils/time";

type ParticleTag = `${string}.${string}`;
type ParticleSet = { tag: ParticleTag, variants: number };

class Environment extends Vector2D {
    public static readonly WEIGHTLESSNESS = new Environment(0, 0);
    public static readonly COMMON = new Environment(0.6, 0.025);

    constructor(gravity: number, wind: number) {
        super(wind, gravity);
    }
}

class Particle {
	private static readonly FPS: number = 60;
	private static readonly LIFETIME_MS: number = 15000;
    private readonly despawnTime: number;
    private readonly element: HTMLImageElement;
    private readonly location: Location2D;
    private readonly velocity: Vector2D;
    private readonly environment: Environment;
    private interval: NodeJS.Timeout | undefined;

    constructor(element: HTMLImageElement, location: Location2D, force: number, environment: Environment) {
        this.interval = undefined;
        this.despawnTime = currentTime() + Particle.LIFETIME_MS;
        this.element = element;
        this.location = location.copy();
        this.velocity = randomVector(force);
        this.environment = environment;
    }

    private needsDespawn(): boolean {
        return isPast(this.despawnTime);
    }

    private update() {
        this.velocity.shift(this.environment);
        this.location.shift(this.velocity);

        move(this.element, this.location);

        if (this.needsDespawn()) {
            this.despawn();
        }
    }

    spawn() {
        document.body.appendChild(this.element);
        this.interval = setInterval(() => this.update(), 1000 / Particle.FPS);
    }

    despawn() {
        this.element.remove();
        clearInterval(this.interval);
    }
}

class Firecracker {
    private particles: Map<ParticleTag, number>;
	private static readonly SIZE: number = 12;
	private static readonly SIZE_SHAKE: number = 12;

	constructor(...particles: ParticleSet[]) {
        this.particles = new Map();

        for (const particle of particles) {
            const tag: ParticleTag = particle.tag;
            const variants: number = particle.variants;

            this.particles.set(tag, variants);
        }
    }

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

	private spawn(particleTag: ParticleTag, position: Location2D, force = 5): void {
		const particleImage = this.createImageByParticleTag(particleTag);
        const environment = Environment.COMMON;
        const particle = new Particle(particleImage, position, shake(force, force), environment);

		particle.spawn();
	}

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