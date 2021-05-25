<script lang="ts">
	import NavBar from '$components/NavBar.svelte';
	import { browser } from '$app/env';

	let page: string;

	$: {
		console.log('running this code');
		const getPage = (page: string): string => {
			page = window?.location?.pathname;
			console.log('page', page);
			switch (page) {
				case '/':
					return (page = 'home');
				case '/blog':
					return (page = 'blog');
				case '/uses':
					return (page = 'uses');
				case '/contact':
					return (page = 'contact');
				default:
					return (page = 'Loading...');
			}
		};
		if (browser) {
			console.log('running browser');
			page = getPage(page);
		}
	}
</script>

<svelte:head>
	<title>{page} | jonathandain.dev</title>
</svelte:head>

<main>
	<NavBar />
	<slot />
</main>
