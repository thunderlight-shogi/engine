/**
 * The flipCoin function returns true with a specified probability.
 * @param {number} [chanceOfTrue=0.5] - The `chanceOfTrue` parameter in the `flipCoin` function
 * represents the probability of the function returning `true`. By default, if no value is provided for
 * `chanceOfTrue`, it is set to 0.5, meaning there is a 50% chance of the function to return true.
 * @returns A boolean value is being returned, which represents the result of flipping a coin based on
 * the given chance of it being true.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function flipCoin(chanceOfTrue: number = 0.5): boolean {
    return Math.random() < chanceOfTrue;
}