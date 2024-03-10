export type Comparator<T> = (first: T, second: T) => number;

export function pick<T>(array: T[]): T {
    if(array.length === 0) {
        throw new Error("Cannot pick an element from an empty array.");
    }

    const index: number = Math.floor(Math.random() * array.length);
    const element: T = array[index];

    return element;
}

function firstFrom<T>(array: T[]): T {
    if (array.length === 0) {
        throw new Error("Cannot get the last element of an empty array.");
    }

    return array[0];
}

function attachProperty<T, K>(value: T, propertyExtractor: (element: T) => K): { element: T, property: K } {
    return {
        element: value,
        property: propertyExtractor(value)
    };
}

export function minBy<T, K>(iterable: Iterable<T>, propertyExtractor: (element: T) => K, compare: Comparator<K>): T {
    const array: T[] = [...iterable];

    if (array.length === 0) {
        throw new Error("Cannot get a maximum element of an empty array.");
    }
    
    const propertyMap = array.map(element => attachProperty(element, propertyExtractor));
    const sortedMap = propertyMap.sort((a, b) => compare(a.property, b.property));
    
    return firstFrom(sortedMap).element;
}