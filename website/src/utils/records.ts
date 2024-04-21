/**
 * Converts a record of `any` to a record of strings via `toString()`.
 * @param obj - An object to convert to a record of strings.
 * @returns A record of strings.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function convertToStringRecord<T extends Record<string, any>>(obj: T): Record<string, string> {
    const result: Record<string, string> = {};
    for (const key in obj) {
        if (Object.prototype.hasOwnProperty.call(obj, key)) {
            result[key] = obj[key].toString();
        }
    }
    return result;
}