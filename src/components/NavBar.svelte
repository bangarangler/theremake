<script lang="ts">
	import { browser } from '$app/env';

	let page: string;

	$: {
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
			page = getPage(page);
		}
	}
</script>

<svelte:head>
	<title>{page} | jonathandain.dev</title>
</svelte:head>

<nav>
	<a href="/">Home</a>
	<a href="/blog">Blog</a>
	<a href="/uses">Uses</a>
	<a href="/contact">Contact</a>
</nav>
