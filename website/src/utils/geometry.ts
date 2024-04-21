import { Comparator } from "./arrays";
import { random, randomUpTo } from "./numbers";

/**
 * The Vector2D class represents a 2D vector
 * with methods for copying, rotating, shortening, shifting, 
 * and getting the angle of the vector. 
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export class Vector2D {
    constructor(public x: number, public y: number) {}

    /**
     * The `copy` function creates and returns a new Vector2D object with the same x and
     * y values as the original vector.
     * @returns A new instance of the Vector2D class with the same x and y values as the current
     * instance.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    public copy(): Vector2D {
        return new Vector2D(this.x, this.y);
    }

    /**
     * The function `rotatedBy` rotates a 2D vector by a specified angle in radians.
     * @param {number} radians - The `radians` parameter represents the angle by which the vector
     * should be rotated in radians.
     * @returns A new `Vector2D` object with its coordinates rotated by the specified number of radians
     * is being returned.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    public rotatedBy(radians: number): Vector2D {
        const sin = Math.sin(radians);
        const cos = Math.cos(radians);
        const x = cos * this.x - sin * this.y;
        const y = sin * this.x + cos * this.y;

        return new Vector2D(x, y);
    }

    /**
     * The `shorten` function divides the `x` and `y` components of a Vector2D by a specified divisor.
     * @param {number} divisor - The `divisor` parameter is a number that is used to divide the `x` and
     * `y` components of a 2D vector in the `shorten` method. If the `divisor` is 0, an error is thrown
     * to prevent division by zero.
     * @returns The `shorten` method is returning the modified `Vector2D` object after dividing its `x`
     * and `y` components by the specified `divisor`.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    public shorten(divisor: number): Vector2D {
        if (divisor === 0) {
            throw new Error("Cannot shorten vector by divisor of zero.");
        }

        this.x /= divisor;
        this.y /= divisor;

        return this;
    }

    /**
     * The `shift` function adds the x and y components of a given Vector2D term to the
     * current Vector2D object and returns the updated object.
     * @param {Vector2D} term - The `term` parameter is a `Vector2D` object that represents the amount
     * by which the current `Vector2D` object should be shifted in both the x and y directions.
     * @returns The `Vector2D` object itself is being returned after the `shift` operation is
     * performed.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    public shift(term: Vector2D): Vector2D {
        this.x += term.x;
        this.y += term.y;

        return this;
    }

    /**
     * The `toString` returns a string representation of a Vector2D object.
     * @returns The `toString` method is being overridden to return a string representation of a
     * Vector2D object with its x and y coordinates. The returned string will be in the format
     * "Vector2D(x, y)", where x and y are the coordinates of the Vector2D object.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    public toString(): string {
        return `Vector2D(${this.x}, ${this.y})`;
    }

    /**
     * The `angle` function returns the arctangent of the ratio of the y-coordinate to the
     * x-coordinate.
     * @returns The `get angle()` method is returning the angle in radians between the positive x-axis
     * and the point (this.x, this.y) using the `Math.atan2()` function.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    get angle(): number {
        return Math.atan2(this.y, this.x);
    }
}

/**
 * The function `randomVector` generates a random 2D vector of a specified length with a random angle.
 * @param {number} length - The `length` parameter in the `randomVector` function represents the
 * magnitude or size of the vector that will be generated. It determines how long the vector will be in
 * a 2D space.
 * @returns A new Vector2D object with a length of the specified value and rotated by a random angle.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function randomVector(length: number) {
    return new Vector2D(0, length).rotatedBy(randomAngle());
}

/** 
 * The `Location2D` class represents a 2D location 
 * with x and y coordinates and provides methods for 
 * copying the location, shifting it by a vector, 
 * and converting it to a string representation. 
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export class Location2D {
    constructor(public x: number, public y: number) {}

    /**
     * The `copy` function creates and returns a new `Location2D` object with the same
     * `x` and `y` coordinates as the original object.
     * @returns A new instance of the `Location2D` class with the same `x` and `y` coordinates as the
     * current instance is being returned.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    public copy(): Location2D {
        return new Location2D(this.x, this.y);
    }

    /**
     * The `shift` function updates the location coordinates by adding the values of a
     * given vector.
     * @param {Vector2D} shift - The `shift` parameter in the `shift` method is of type `Vector2D`,
     * which represents a 2-dimensional vector with `x` and `y` components. This method is used to
     * shift the current `Location2D` object by the specified `Vector2D` amount in
     * @returns The `Location2D` object itself is being returned after the `shift` operation is
     * performed.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    public shift(shift: Vector2D): Location2D {
        this.x += shift.x;
        this.y += shift.y;

        return this;
    }

    /**
     * The `toString` function returns a string representation of a Location2D object with its x and y
     * coordinates.
     * @returns The `toString` method is returning a string representation of the `Location2D` object
     * with its `x` and `y` coordinates. The format of the returned string is "Location2D(x, y)", where
     * `x` and `y` are the values of the `x` and `y` properties of the object.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    public toString(): string {
        return `Location2D(${this.x}, ${this.y})`;
    }
}

/**
 * The function `randomAngle` returns a random angle value within the range of -π to π.
 * @returns The function `randomAngle` returns a random angle value between -π and π.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
function randomAngle(): number {
    return random(-Math.PI, Math.PI);
}

/**
 * The function `centerLocation` returns a new `Location2D` object with coordinates (0, 0).
 * @returns A new `Location2D` object with coordinates (0, 0) is being returned.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function centerLocation(): Location2D {
    return new Location2D(0, 0);
}

/**
 * The `randomLocation` generates a random 2D location within the dimensions of the window.
 * @returns A Location2D object with random x and y coordinates within the screen dimensions.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function randomLocation(): Location2D {
    const screenWidth: number = window.innerWidth;
    const screenHeight: number = window.innerHeight;
    const x = randomUpTo(screenWidth);
    const y = randomUpTo(screenHeight);

    return new Location2D(x, y);
}

/**
 * The function calculates the Euclidean distance between two 2D locations.
 * @param {Location2D} first - The first location.
 * @param {Location2D} second - The second location.
 * @returns The function `distanceBetween` calculates the Euclidean distance between two 2D locations
 * (`first` and `second`) and returns the distance as a number.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function distanceBetween(first: Location2D, second: Location2D): number {
    return Math.sqrt((first.x - second.x) ** 2 + (first.y - second.y) ** 2);
}

/**
 * The function `byDistanceTo` returns a comparator function that sorts locations based on their
 * distance to a specified location.
 * @param {Location2D} location - The `location` parameter in the `byDistanceTo` function is of type
 * `Location2D`. It represents a two-dimensional location on a map or grid, typically defined by its x
 * and y coordinates.
 * @returns A comparator function that compares two `Location2D` objects based on their distance to a
 * specified `location`.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function byDistanceTo(location: Location2D): Comparator<Location2D> {
    return (first, second) => distanceBetween(first, location) - distanceBetween(second, location);
}

/**
 * The function calculates the angle between two 2D locations.
 * @param {Location2D} first - The `first` parameter represents the first location in 2D space with
 * coordinates (x, y).
 * @param {Location2D} second - The `second` parameter in the `angleBetween` function represents a 2D
 * location in space. It likely has `x` and `y` coordinates that define its position relative to the
 * origin (0,0).
 * @returns The function `angleBetween` calculates the angle in radians between two 2D locations
 * (`first` and `second`) and returns the result as a number.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function angleBetween(first: Location2D, second: Location2D): number {
    return Math.atan2(second.y - first.y, second.x - first.x);
}