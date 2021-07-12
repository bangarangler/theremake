import { mdsvex } from 'mdsvex';
import { mdsvexConfig } from './mdsvex.config.js';
import preprocess from 'svelte-preprocess';
// import node from "@sveltejs/adapter-node";
// import pkg from "./package.json";
import { resolve } from 'path';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	extensions: ['.svelte', ...mdsvexConfig.extensions],
	preprocess: [mdsvex(mdsvexConfig), preprocess()],

	kit: {
		// hydrate the <div id="svelte"> element in src/app.html
		target: '#svelte',
		vite: {
			resolve: {
				alias: {
					$components: resolve('./src/components'),
					$images: resolve('./src/images'),
					$stores: resolve('./src/stores'),
					$projects: resolve('./src/projects'),
					$blogposts: resolve('./src/blogposts')
				}
			}
		}
	}
};

export default config;
