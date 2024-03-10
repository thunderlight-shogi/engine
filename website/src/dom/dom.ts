import { minBy } from "../utils/arrays";
import { Location2D, byDistanceTo } from "../utils/geometry";
import { mean } from "../utils/numbers";

function locate(element: HTMLElement): Location2D {
    const rectangle = element.getBoundingClientRect();
    const x = mean(rectangle.top, rectangle.bottom);
    const y = mean(rectangle.left, rectangle.right);

    return new Location2D(x, y);
}

export function closest(neighborhood: Iterable<HTMLElement>, element: HTMLElement) {
    const location = locate(element);
    return minBy(neighborhood, locate, byDistanceTo(location));
}

export function byClass(className: string): HTMLElement[] {
    return [...document.getElementsByClassName(className) as HTMLCollectionOf<HTMLElement>];
}