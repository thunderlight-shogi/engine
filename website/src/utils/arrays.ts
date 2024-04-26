export type Comparator<T> = (first: T, second: T) => number;

/**
 * The function `pick` randomly selects and returns an element from an array.
 * @param {T[]} array - The `pick` function takes an array of type `T` as a parameter.
 * @returns The function `pick` returns a randomly selected element from the input array `array`.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function pick<T>(array: T[]): T {
    if(array.length === 0) {
        throw new Error("Cannot pick an element from an empty array.");
    }

    const index: number = Math.floor(Math.random() * array.length);
    const element: T = array[index];

    return element;
}

/**
 * The function `pick` randomly selects and returns an element from an array.
 * @param {T[]} array - The `pick` function takes an array of type `T` as a parameter.
 * @returns The function `pick` returns a randomly selected element from the input array `array`.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function indexedPick<T>(array: T[]): {index: number, element: T} {
    if(array.length === 0) {
        throw new Error("Cannot pick an element from an empty array.");
    }

    const index: number = Math.floor(Math.random() * array.length);
    const element: T = array[index];

    return {index, element};
}

/**
 * The function `firstFrom` returns the first element of an array.
 * @param {T[]} array - The `array` parameter is an array of type `T`, which means it can hold elements
 * of any type. The function `firstFrom` takes an array as input and returns the first element of that
 * array. If the array is empty, it throws an error.
 * @returns The function `firstFrom` is returning the first element of the input array `array`.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
function firstFrom<T>(array: T[]): T {
    if (array.length === 0) {
        throw new Error("Cannot get the last element of an empty array.");
    }

    return array[0];
}

/**
 * The function `attachProperty` takes a value and a property extractor function, and returns an object
 * containing the original value and the extracted property.
 * @param {T} value - The `value` parameter represents the input value of type `T` that you want to
 * attach a property to.
 * @param propertyExtractor - The `propertyExtractor` parameter is a function that takes an element of
 * type `T` as its argument and returns a property of type `K` extracted from that element.
 * @returns The `attachProperty` function returns an object with two properties: `element` which is the
 * original value passed to the function, and `property` which is the result of applying the
 * `propertyExtractor` function on the original value.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
function attachProperty<T, K>(value: T, propertyExtractor: (element: T) => K): { element: T, property: K } {
    return {
        element: value,
        property: propertyExtractor(value)
    };
}

/**
 * The `minBy` function returns the minimum element from an iterable based on a specified property and
 * comparison function.
 * @param iterable - The `iterable` parameter is an iterable collection of elements from which you want
 * to find the minimum element based on a specific property. This can be an array, a Set, a Map, or any
 * other iterable data structure.
 * @param propertyExtractor - The `propertyExtractor` parameter is a function that takes an element of
 * type `T` and extracts a property of type `K` from it. This property will be used for comparison to
 * determine the minimum element in the iterable.
 * @param compare - The `compare` parameter in the `minBy` function is a function that compares two
 * values of type `K`. It is used to determine the minimum value based on the extracted property of
 * each element in the iterable. The `compare` function should follow the `Comparator` type, which
 * typically looks
 * @returns The `minBy` function returns the element from the input `iterable` that has the minimum
 * value when the `propertyExtractor` function is applied to it, based on the comparison defined by the
 * `compare` function.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function minBy<T, K>(iterable: Iterable<T>, propertyExtractor: (element: T) => K, compare: Comparator<K>): T {
    const array: T[] = [...iterable];

    if (array.length === 0) {
        throw new Error("Cannot get a maximum element of an empty array.");
    }
    
    const propertyMap = array.map(element => attachProperty(element, propertyExtractor));
    const sortedMap = propertyMap.sort((a, b) => compare(a.property, b.property));
    
    return firstFrom(sortedMap).element;
}

/**
 * The function `create2DArray` generates a 2D array with a specified width and height.
 * @param {number} width - The `width` parameter in the `create2DArray` function represents the number
 * of columns in the 2D array that will be created. It determines the horizontal size of the array.
 * @param {number} height - The `height` parameter in the `create2DArray` function represents the
 * number of rows in the 2D array that will be created.
 * @returns The `create2DArray` function is returning a 2D array of type `(E | undefined)[][]`, where
 * `E` is a generic type. Each element in the 2D array can hold a value of type `E` or `undefined`.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function create2DArray<E>(width: number, height: number): (E | undefined)[][] {
    const array = new Array<(E | undefined)[]>(height);

    for (let x = 0; x < height; x += 1) {
        array.push(new Array<E | undefined>(width));
    }

    return array;
}

/**
 * The `flatten` function in TypeScript takes a two-dimensional array and returns a flattened
 * one-dimensional array.
 * @param {E[][]} array - The `array` parameter in the `flatten` function is a two-dimensional array of
 * elements of type `E`.
 * @returns The `flatten` function returns a single-dimensional array containing all the elements from
 * the input two-dimensional array after flattening it.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function flatten<E>(array: E[][]): E[] {
    let flattenArray: E[] = [];

    for(const row of array) {
        flattenArray = flattenArray.concat(row);
    }

    return flattenArray;
}