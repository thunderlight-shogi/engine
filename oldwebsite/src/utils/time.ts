/**
 * Returns the current time.
 * @returns - The current timestamp in milliseconds.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function currentTime(): number {
    return Date.now();
}

/**
 * Check whether the provided timestamp is in the past or not.
 * @param time - The checking timestamp in milliseconds.
 * @returns Is the provided time in the past?
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function isPast(time: number): boolean {
    return currentTime() > time;
}