import { writable } from "svelte/store";
import type { Route } from "$types/route";

export const routes = writable<Route[]>([]);

export const getRoutes = async (url: string): Promise<Route[] | undefined> => {
	try {
		const res = await fetch(url);
		if (res.ok) {
			const data: Route[] = await res.json();
			return data;
		} else {
			throw new Error(res.statusText);
		}
	} catch (err) {
		console.log(err);
	}
};
