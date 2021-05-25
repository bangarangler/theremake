<script lang="ts">
	import NavBar from '$components/NavBar.svelte';
	import { page } from '$app/stores';
	$: path = $page.path.replace('/', '');
	$: title = path !== '' ? path : 'home';
</script>

<svelte:head>
	<title>{title} | jonathandain.dev</title>
	<meta
		name="description"
		content="
Hi I'm Jon! I'm a web developer with a extreme interest in marketing and business! I've owned a few business's in my time and I still enjoy the thrill of planning and organizing a business. Regardless if it's day one or you have been in business for generations; I'm confident that I can make it even better!
							"
	/>
	<meta
		name="keywords"
		content="coding, developer, react, software,
		 portfolio, software developer, web developer, fullstack developer"
	/>
	<meta name="author" content="Jon Palacio" />
	<meta property="og:title" content="jonathandain.dev" />
	<meta
		property="og:description"
		content="
Hi I'm Jon! I'm a web developer with a extreme interest in marketing and
business! I've owned a few business's in my time and I still enjoy the thrill of
planning and organizing a business. Regardless if it's day one or you have been
in business for generations; I'm confident that I can make it even better!"
	/>
	<meta property="og:type" content="website" />
	<meta property="og:url" content="https://jonathandain.dev" />
	<!--<meta property="og:image" content="" />-->
	<meta name="twitter:card" content="summary" />
	<meta name="twitter:creator" content="Jon Palacio" />
	<meta name="twitter:title" content={title} />
	<meta
		name="twitter:description"
		content="Hi I'm Jon! I'm a web developer with a extreme interest in marketing and business! I've owned a few business's in my time and I still enjoy the thrill of planning and organizing a business. Regardless if it's day one or you have been in business for generations; I'm confident that I can make it even better!"
	/>
	<meta property="og:site_name" content="jonathandain.dev" />
	<meta
		name="twitter:image:alt"
		content="displaying website image for
jonathandain.dev"
	/>
</svelte:head>

<main>
	<NavBar />
	<slot />
</main>
