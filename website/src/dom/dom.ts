import { minBy } from "../utils/arrays";
import { Location2D, byDistanceTo } from "../utils/geometry";

export function locate(element: HTMLElement): Location2D {
    const rectangle = element.getBoundingClientRect();
    const x = rectangle.top;
    const y = rectangle.left;

    return new Location2D(x, y);
}

export function widthOf(element: HTMLElement): number {
    const rectangle = element.getBoundingClientRect();
    const width = rectangle.left - rectangle.right;

    return width;
}

export function heightOf(element: HTMLElement): number {
    const rectangle = element.getBoundingClientRect();
    const height = rectangle.top - rectangle.bottom;

    return height;
}

export function closest(neighborhood: Iterable<HTMLElement>, element: HTMLElement) {
    const location = locate(element);
    return minBy(neighborhood, locate, byDistanceTo(location));
}

export function byClass(className: string): HTMLElement[] {
    return [...document.getElementsByClassName(className) as HTMLCollectionOf<HTMLElement>];
}