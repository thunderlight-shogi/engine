/**
 * Generates a version 4 UUID. 
 * Learn more: https://wikipedia.org/wiki/UUID/
 */
export function generateUUIDv4(): string {
    return "10000000-1000-4000-8000-100000000000".replace(/[018]/g, (entry: string) => {
        const int = parseInt(entry);
        return (int ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> int / 4).toString(16);
    });
}