import { minBy } from "../utils/arrays";
import { Location2D, Vector2D, byDistanceTo } from "../utils/geometry";
import { mean } from "../utils/numbers";
import { px } from "./css";

export function locate(element: HTMLElement): Location2D {
    const rectangle = element.getBoundingClientRect();
    const x = mean(rectangle.left, rectangle.right);
    const y = mean(rectangle.top, rectangle.bottom);

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

export function move(element: HTMLElement, absoluteLocation: Location2D) {
    element.style.left = px(absoluteLocation.x);
    element.style.top = px(absoluteLocation.y);
}

export function measure(element: HTMLElement) {
    return new Vector2D(widthOf(element), heightOf(element));
}

export function resize(element: HTMLElement, size: number) {
    element.style.width = px(size);
    element.style.height = px(size);
}

export function addClass(element: HTMLElement, className: string) {
    element.classList.add(className);
}

export function removeClass(element: HTMLElement, className: string) {
    element.classList.remove(className);
}

export function locateMouse(event: MouseEvent): Location2D {
    const x = event.pageX;
    const y = event.pageY;

    return new Location2D(x, y);
}