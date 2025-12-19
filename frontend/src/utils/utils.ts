import type { RouteLocationNormalizedLoaded } from 'vue-router'

export function importName(route: RouteLocationNormalizedLoaded): string {
    const list = route.name?.toString().split("-") || [];
    let lastWord = list[list.length - 1] || "";

    if (lastWord.endsWith("s")) {
        lastWord = lastWord.slice(0, -1);
    }

    return lastWord;
}