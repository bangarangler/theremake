import { writable } from 'svelte/store';

interface CreateTheme {
	subscribe: any;
	lightMode: () => void;
	darkMode: () => void;
	reset: () => void;
}

function createTheme(): CreateTheme {
	const { subscribe, set, update } = writable('DARK');
	return {
		subscribe,
		lightMode: () => update(() => 'LIGHT'),
		darkMode: () => update(() => 'DARK'),
		reset: () => set('DARK')
	};
}

export const theme = createTheme();
