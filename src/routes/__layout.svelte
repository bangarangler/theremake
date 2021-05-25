<script lang="ts">
	import NavBar from '$components/NavBar.svelte';
	import { page } from '$app/stores';
	$: path = $page.path.replace('/', '');
	$: title = path !== '' ? path : 'home';
</script>

<svelte:head>
	<title>{title} | jonathandain.dev</title>
</svelte:head>

<main>
	<NavBar />
	<slot />
</main>
