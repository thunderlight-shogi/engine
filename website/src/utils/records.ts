export function convertToStringRecord<T extends Record<string, any>>(obj: T): Record<string, string> {
    const result: Record<string, string> = {};
    for (const key in obj) {
        if (Object.prototype.hasOwnProperty.call(obj, key)) {
            result[key] = obj[key].toString();
        }
    }
    return result;
}