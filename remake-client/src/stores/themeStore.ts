import { writable } from 'svelte/store';

interface CreateTheme {
	subscribe: any;
	lightMode: () => void;
	darkMode: () => void;
	toggle: (whatTheme: string) => any;
	reset: () => void;
}

function createTheme(): CreateTheme {
	const { subscribe, set /*update*/ } = writable('DARK');
	return {
		subscribe,
		lightMode: () => set('LIGHT'),
		darkMode: () => set('DARK'),
		toggle: (whatTheme) => {
			if (whatTheme === 'DARK') {
				set('LIGHT');
			} else {
				set('DARK');
			}
		},
		reset: () => set('DARK')
	};
}

export const theme = createTheme();

/* on:click={!$theme ? theme.lightMode : theme.darkMode} */
// bind:checked={$theme}
