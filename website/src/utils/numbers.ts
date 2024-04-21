import { flipCoin } from "./booleans";

/**
 * The function `randomSign` returns either 1 or -1 randomly.
 * @returns The `randomSign` function returns either 1 or -1 randomly.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
function randomSign(): number {
    return flipCoin() ? 1 : -1;
}

/**
 * The function `randomUpTo` generates a random number between 0 and a specified threshold.
 * @param {number} threshold - The `threshold` parameter in the `randomUpTo` function represents the
 * maximum value that the random number generated should not exceed.
 * @returns The function `randomUpTo` returns a random number between 0 and the specified `threshold`
 * value.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function randomUpTo(threshold: number): number {
    return random(0, threshold);
}

/**
 * The `random` function generates a random number within a specified range.
 * @param {number} lowerThreshold - The `lowerThreshold` parameter represents the lower end of the
 * range from which you want to generate a random number.
 * @param {number} upperThreshold - The `upperThreshold` parameter represents the upper limit of the
 * range within which you want to generate a random number.
 * @returns The function `random` returns a random number within the range specified by the
 * `lowerThreshold` and `upperThreshold` parameters.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function random(lowerThreshold: number, upperThreshold: number): number {
    return Math.random() * (upperThreshold - lowerThreshold) + lowerThreshold;
}

/**
 * The function `randomInt` generates a random integer within a specified range.
 * @param {number} lowerThreshold - The `lowerThreshold` parameter represents the lower end of the
 * range from which you want to generate a random integer.
 * @param {number} upperThreshold - The `upperThreshold` parameter represents the maximum value that
 * the random integer can take.
 * @returns The function `randomInt` returns a random integer between the `lowerThreshold` and
 * `upperThreshold` values.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function randomInt(lowerThreshold: number, upperThreshold: number): number {
    return Math.floor(random(lowerThreshold, upperThreshold));
}

/**
 * The function "between" checks if a number is within a specified range.
 * @param {number} min - The `min` parameter represents the minimum value in the range you want to
 * check.
 * @param {number} number - The `number` parameter in the `between` function represents the value that
 * you want to check if it falls within the range defined by `min` and `max`.
 * @param {number} max - The `max` parameter represents the maximum value that the `number` parameter
 * can have in order for the `between` function to return `true`.
 * @returns The function `between` is returning a boolean value indicating whether the `number` is
 * within the range defined by `min` and `max`.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function between(min: number, number: number, max: number): boolean {
    return min <= number && number <= max;
}

/**
 * The `shake` function takes a number and adds a random amount within a specified range to it.
 * @param {number} number - The `number` parameter is the initial number that you want to shake or add
 * some randomness to.
 * @param {number} shakeAmount - The `shakeAmount` parameter represents the maximum amount by which the
 * `number` can be shaken. It determines the range within which the `number` can be adjusted randomly.
 * @returns The function `shake` is returning a number that is the result of adding the input `number`
 * to a random value within the range of `-shakeAmount` to `shakeAmount`.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function shake(number: number, shakeAmount: number): number {
    return number + randomSign() * randomUpTo(shakeAmount);
}

/**
 * The function calculates the absolute distance between two numbers.
 * @param {number} first - The `first` parameter is a number that represents the first value for which
 * you want to calculate the distance.
 * @param {number} second - The `second` parameter in the `distance` function represents the second
 * number for which you want to calculate the absolute difference from the first number.
 * @returns The function `distance` returns the absolute difference between the two input numbers
 * `first` and `second`.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function distance(first: number, second: number) {
    return Math.abs(first - second);
}

/**
 * Calculates the mean of two numbers.
 * @param first - The first number.
 * @param second - The second number.
 * @returns The mean of two numbers.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function mean(first: number, second: number): number {
    return (first + second) / 2;
}

/**
 * Translates a number from one number range to another.
 * 
 * E.g. `translate(0, 1, 40, 0, 100) === 0.4`, since 40 is located between 0 and 100 in the same
 * proportion as 0.4 between 0 and 1.
 * 
 * @param translationMin - The min of target range.
 * @param translationMax - The max of target range.
 * @param number - The number from original range.
 * @param originalMin - The min of original range.
 * @param originalMax - The max of original range.
 * @returns The number translated into a target range.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function translate(translationMin: number, translationMax: number, number: number, originalMin: number, originalMax: number) {
    const originalRange = distance(originalMin, originalMax);
    const translationRange = distance(translationMin, translationMax);
    const numberRate = number / originalRange;

    return numberRate * translationRange + translationMin;
}