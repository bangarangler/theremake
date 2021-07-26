import { browser } from '$app/env';
import { writable } from 'svelte/store';

interface CreateTheme {
	subscribe: any;
	lightMode: () => void;
	darkMode: () => void;
	toggle: (whatTheme: string) => any;
	reset: () => void;
}

function updateLocalStorage(selection: string) {
	if (browser) {
		localStorage.setItem('userTheme', selection);
	}
}

function createTheme(): CreateTheme {
	let stored: string;
	if (browser) {
		stored = localStorage?.userTheme;
	}
	const { subscribe, set /*update*/ } = writable(stored || 'DARK');
	// const { subscribe, set /*update*/ } = writable('DARK');
	return {
		subscribe,
		lightMode: () => {
			updateLocalStorage('LIGHT');
			set('LIGHT');
		},
		darkMode: () => {
			updateLocalStorage('DARK');
			set('DARK');
		},
		toggle: (whatTheme) => {
			if (whatTheme === 'DARK') {
				updateLocalStorage('LIGHT');
				set('LIGHT');
			} else {
				updateLocalStorage('DARK');
				set('DARK');
			}
		},
		reset: () => {
			updateLocalStorage('DARK');
			set('DARK');
		}
	};
}

export const theme = createTheme();
