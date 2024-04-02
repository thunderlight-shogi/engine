export function throws(errorMessage: string): () => never {
    return () => {
        throw new Error(errorMessage);
    }
}