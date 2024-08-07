import { convertToStringRecord } from "./records";

export type HTTPMethod = "get" | "post"
export type HTTPParams = Record<string, unknown>;

export class RestAPI {
    constructor(
        private readonly domain: string
    ) {}

    public async get(path: string, parameters: HTTPParams = {}): Promise<any> {
        return await this.fetch(path, "get", parameters);
    }

    public async post(path: string, parameters: HTTPParams = {}): Promise<any> {
        return await this.fetch(path, "post", parameters);
    }

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

    private convertRequestToString(path: string, method: HTTPMethod, parameters: HTTPParams = {}): string {
        return `@${this.domain} <--${method}-- ${path}(${JSON.stringify(parameters)})`;
    }
}

