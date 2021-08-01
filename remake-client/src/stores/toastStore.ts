import { writable } from 'svelte/store';

const newToast = () => {
	const { subscribe, update } = writable([]);

	function send(msg: string) {
		update((state) => {
			return [...state, msg];
		});
	}

	function remove() {
		update((state) => {
			let [_, ...rest] = state;
			return [...rest];
		});
	}

	return { subscribe, send, remove };
};

export const toast = newToast();

export const isValid = writable(false);
