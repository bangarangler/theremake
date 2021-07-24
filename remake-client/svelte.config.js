import { mdsvex } from 'mdsvex';
import { mdsvexConfig } from './mdsvex.config.js';
import preprocess from 'svelte-preprocess';
import { imagetools } from 'vite-imagetools';
// import node from '@sveltejs/adapter-node';
// import pkg from "./package.json";
import path from 'path';
import adapter from '@sveltejs/adapter-static';
// import adapter from '@sveltejs/adapter-node';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	extensions: ['.svelte', ...mdsvexConfig.extensions],
	preprocess: [mdsvex(mdsvexConfig), preprocess()],

	kit: {
		// hydrate the <div id="svelte"> element in src/app.html
		// adapter: { adapter: node() },
		// adapter: node(),
		target: '#svelte',
		adapter: adapter(),
		vite: {
			resolve: {
				alias: {
					$components: path.resolve('./src/components'),
					$images: path.resolve('./src/images'),
					$stores: path.resolve('./src/stores'),
					$projects: path.resolve('./src/projects'),
					$blogposts: path.resolve('./src/blogposts'),
					$static: path.resolve('./static')
				}
			},
			plugins: [imagetools({ force: true })]
		}
	}
};

export default config;
