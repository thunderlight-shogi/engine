import { minBy } from "../utils/arrays";
import { Location2D, Vector2D, byDistanceTo } from "../utils/geometry";
import { mean } from "../utils/numbers";
import { px } from "./css";

/**
 * The function `locate` calculates the center coordinates of a given HTML element.
 * @param {HTMLElement} element - An HTMLElement representing an element in the DOM (Document Object
 * Model) of a web page.
 * @returns The function `locate` returns a `Location2D` object with the x and y coordinates of the
 * center of the bounding rectangle of the given `HTMLElement` element.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function locate(element: HTMLElement): Location2D {
    const rectangle = element.getBoundingClientRect();
    const x = mean(rectangle.left, rectangle.right);
    const y = mean(rectangle.top, rectangle.bottom);

    return new Location2D(x, y);
}

/**
 * The function `widthOf` calculates the width of an HTML element using its bounding rectangle.
 * @param {HTMLElement} element - The `element` parameter in the `widthOf` function is expected to be
 * an HTMLElement object representing an element in the HTML document. This function calculates and
 * returns the width of the element by getting its bounding rectangle and calculating the difference
 * between its left and right coordinates.
 * @returns the width of the element by calculating the difference between the left and right
 * boundaries of the element's bounding rectangle.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function widthOf(element: HTMLElement): number {
    const rectangle = element.getBoundingClientRect();
    const width = rectangle.left - rectangle.right;

    return width;
}

/**
 * The function `heightOf` calculates the height of an HTML element using its bounding rectangle.
 * @param {HTMLElement} element - The `element` parameter in the `heightOf` function is expected to be
 * an HTMLElement object representing an element in the HTML document. This function calculates and
 * returns the width of the element by getting its bounding rectangle and calculating the difference
 * between its top and bottom coordinates.
 * @returns the height of the element by calculating the difference between the top and bottom
 * boundaries of the element's bounding rectangle.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function heightOf(element: HTMLElement): number {
    const rectangle = element.getBoundingClientRect();
    const height = rectangle.top - rectangle.bottom;

    return height;
}

/**
 * The function `closest` finds the closest element in a neighborhood to a given element based on their
 * locations.
 * @param neighborhood - The `neighborhood` parameter is an iterable collection of HTMLElement objects
 * representing elements in a specific neighborhood or group.
 * @param {HTMLElement} element - The `element` parameter represents the HTML element for which you
 * want to find the closest neighbor within a given neighborhood of elements.
 * @returns The `closest` function is returning the element in the `neighborhood` iterable that is
 * closest to the `element` based on their locations.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function closest(neighborhood: Iterable<HTMLElement>, element: HTMLElement) {
    const location = locate(element);
    return minBy(neighborhood, locate, byDistanceTo(location));
}

/**
 * The function `closestTo` returns the element in a neighborhood that is closest to a given location
 * in a 2D space.
 * @param neighborhood - The `neighborhood` parameter is an iterable collection of HTMLElement objects
 * representing elements in a neighborhood or area.
 * @param {Location2D} location - The `location` parameter represents a 2D location in a coordinate
 * system. It could be an object with `x` and `y` properties representing the coordinates on a 2D
 * plane.
 * @returns The `closestTo` function is returning the element in the `neighborhood` that is closest to
 * the specified `location` based on the distance calculated using the `byDistanceTo` function.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function closestTo(neighborhood: Iterable<HTMLElement>, location: Location2D) {
    return minBy(neighborhood, locate, byDistanceTo(location));
}

/**
 * The function `byClass` returns an array of HTMLElement objects with the specified class name.
 * @param {string} className - The `className` parameter is a string that represents the class name of
 * the elements you want to select from the document.
 * @returns An array of HTMLElement objects with the specified class name is being returned.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function byClass(className: string): HTMLElement[] {
    return [...document.getElementsByClassName(className) as HTMLCollectionOf<HTMLElement>];
}

/**
 * The `move` function sets the left and top style properties of an HTML element to the x
 * and y coordinates of a given Location2D object, respectively.
 * @param {HTMLElement} element - The `element` parameter is an HTMLElement, which represents an HTML
 * element in the DOM (Document Object Model). It could be any HTML element like a div, span, button,
 * etc.
 * @param {Location2D} absoluteLocation - The `absoluteLocation` parameter represents the absolute
 * position of the element on the screen. It is defined by a 2D coordinate system, typically with an
 * x-coordinate (horizontal position) and a y-coordinate (vertical position).
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function move(element: HTMLElement, absoluteLocation: Location2D) {
    element.style.left = px(absoluteLocation.x);
    element.style.top = px(absoluteLocation.y);
}

/**
 * The `measure` function returns a `Vector2D` object representing the width and height
 * of an HTML element.
 * @param {HTMLElement} element - An HTMLElement representing an element in the HTML document.
 * @returns The `measure` function is returning a new `Vector2D` object that contains the width and
 * height of the provided `HTMLElement` element.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function measure(element: HTMLElement) {
    return new Vector2D(widthOf(element), heightOf(element));
}

/**
 * The `resize` function resizes an HTML element by setting its width and height to a
 * specified size in pixels.
 * @param {HTMLElement} element - The `element` parameter is an HTMLElement, which represents an HTML
 * element in the DOM (Document Object Model).
 * @param {number} size - The `size` parameter in the `resize` function is a number that represents the
 * desired width and height (in pixels) to which the `element` should be resized.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function resize(element: HTMLElement, size: number) {
    element.style.width = px(size);
    element.style.height = px(size);
}

/**
 * The function `addClass` adds a specified class to an HTML element.
 * @param {HTMLElement} element - The `element` parameter is an HTMLElement object, which represents an
 * HTML element in the DOM.
 * @param {string} className - The `className` parameter is a string that represents the class name you
 * want to add to the specified `HTMLElement` element.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function addClass(element: HTMLElement, className: string) {
    element.classList.add(className);
}

/**
 * The function `removeClass` removes a specified class from an HTML element.
 * @param {HTMLElement} element - The `element` parameter is an HTMLElement object, which represents an
 * HTML element in the DOM.
 * @param {string} className - The `className` parameter is a string that represents the class name you
 * want to remove from the specified `element`.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function removeClass(element: HTMLElement, className: string) {
    element.classList.remove(className);
}

/**
 * The function `triggerReflow` forces a reflow on a given HTML element by accessing its `offsetHeight`
 * property.
 * @param {HTMLElement} element - The `element` parameter is an HTMLElement object representing an
 * element in the HTML document.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function triggerReflow(element: HTMLElement) {
    element.offsetHeight;
}

/**
 * The function `triggerClass` removes a class from an HTML element, triggers a reflow, and then adds
 * the class back after a 3-second delay.
 * @param {HTMLElement} element - The `element` parameter is an HTMLElement object, which represents an
 * HTML element in the DOM (Document Object Model). It could be any HTML element like a div, span,
 * button, etc.
 * @param {string} className - The `className` parameter in the `triggerClass` function is a string
 * that represents the class name that you want to add to the specified `element`.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function triggerClass(element: HTMLElement, className: string) {
    removeClass(element, className);
    triggerReflow(element);
    setTimeout(() => addClass(element, className), 0);
}

/**
 * The function `locateMouse` takes a MouseEvent as input and returns a Location2D object with the x
 * and y coordinates of the mouse pointer.
 * @param {MouseEvent} event - A MouseEvent parameter that contains information about the mouse event
 * that occurred, such as the position of the mouse cursor on the page.
 * @returns An instance of the `Location2D` class with the x and y coordinates of the mouse pointer on
 * the page.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function locateMouse(event: MouseEvent): Location2D {
    const x = event.pageX;
    const y = event.pageY;

    return new Location2D(x, y);
}

/**
 * The function `indexOfChild` returns the index of a given HTML element among its parent's children.
 * @param {HTMLElement} child - The `child` parameter in the `indexOfChild` function is an
 * `HTMLElement` representing the child element for which you want to find the index within its
 * parent's list of children.
 * @returns The function `indexOfChild` returns the index of the given `child` element among its
 * siblings within the parent element.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export function indexOfChild(child: HTMLElement): number {
    const parent = child.parentNode;

    if (parent === null) {
        throw Error("The element has no parent, hence doesn't have child index.");
    }

    const siblings: HTMLElement[] = [...parent?.children] as HTMLElement[];
    return siblings.indexOf(child);
}