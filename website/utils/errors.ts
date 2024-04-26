/**
 * The function `throws` returns a function that throws an error with the specified error
 * message. Use this function as a callback higher order function, e.g. for actions where 
 * an error must always be thrown.
 * @param {string} errorMessage - The `errorMessage` parameter is a string that represents the error
 * message that will be thrown when the returned function is called.
 * @returns A function is being returned that, when called, will throw an error with the specified
 * error message.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function throws(errorMessage: string): () => never {
    return () => {
        throw new Error(errorMessage);
    }
}