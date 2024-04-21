/**
 * The `px` function in TypeScript takes a number value representing pixels and returns it as a string
 * with "px" appended.
 * @param {number} pixels - The `pixels` parameter is a number representing the value in pixels that
 * you want to convert to a string with the "px" unit appended.
 * @returns The function `px` takes a number as input and returns a string with the number followed by
 * "px".
 */
export function px(pixels: number): string {
    return `${pixels}px`;
}