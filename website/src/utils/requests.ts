import { convertToStringRecord } from "./records";

export type HTTPMethod = "get" | "post"
export type HTTPParams = Record<string, unknown>;

/**
 * This class is designed to provide an easy access to APIs over HTTP protocol.
 * 
 * @author Anatoly Frolov <contact@anafro.ru>
 */
export class RestAPI {
    /**
     * Creates a new RestAPI client.
     * @param domain - A domain for the API. 
     */
    constructor(
        private readonly domain: string
    ) {}

    /**
     * Performs a GET request.
     * 
     * @param path - The path to GET.
     * @param parameters - The GET parameters.
     * @returns The response in an object.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    public async get(path: string, parameters: HTTPParams = {}): Promise<any> {
        return await this.fetch(path, "get", parameters);
    }

    /**
     * Performs a POST request.
     * 
     * @param path - The path to POST.
     * @param parameters - The POST parameters.
     * @returns The response in an object.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    public async post(path: string, parameters: HTTPParams = {}): Promise<any> {
        return await this.fetch(path, "post", parameters);
    }

    /**
     * Performs an HTTP request with provided parameters.
     * 
     * @param path - The path to a route.
     * @param parameters - The request parameters.
     * @returns The response in an object.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    private async fetch(path: string, method: HTTPMethod, parameters: Record<string, any> = {}): Promise<any> {
        const url = new URL(this.domain);
        url.pathname = path;
        
        if (method === "get") {
            url.search = new URLSearchParams(convertToStringRecord(parameters)).toString();
        }

        const response = await fetch(url, {
            ...(method === "post" && {body: JSON.stringify(parameters)}),
            method
        });

        if (!response.ok) {
            throw new Error(`
                ${this.convertRequestToString(path, method, parameters)} failed.
                The HTTP code was ${response.status}: ${response.statusText}
            `);
        }

        const text = await response.text();

        if (text.length === 0) {
            return {};
        }

        try {
            return JSON.parse(text);
        } catch(error: unknown) {
            throw new Error(`
                ${this.convertRequestToString(path, method, parameters)} did not respond with a valid JSON string.
                Instead, the response was:
                "${text}"
            `)
        }
    }

    /**
     * Converts HTTP request to a human-friendly string.
     * 
     * @param path - The path to a route.
     * @param method - The HTTP method.
     * @param parameters - The parameters to a route.
     * @returns A request as a string.
     * 
     * @author Anatoly Frolov <contact@anafro.ru>
     */
    private convertRequestToString(path: string, method: HTTPMethod, parameters: HTTPParams = {}): string {
        return `@${this.domain} <--${method}-- ${path}(${JSON.stringify(parameters)})`;
    }
}

