export function pick<T>(array: T[]): T {
    if(array.length === 0) {
        throw new Error("Cannot pick an element from an empty array.");
    }

    const index: number = Math.floor(Math.random() * array.length);
    const element: T = array[index];

    return element;
}