/**
 * Sleeps for a provided amount of time in milliseconds.
 * @param milliseconds - The amount of milliseconds to sleep.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function sleep(milliseconds: number): Promise<void> {
    return new Promise(resolve => setTimeout(resolve, milliseconds));
}