<script>
	import { onMount } from 'svelte';
	import { isValid, toast } from '$stores/toastStore';
	export let message = '';
	export let duration = 1000;
	import { tweened } from 'svelte/motion';
	let progress = tweened(100, { duration });
	onMount(async () => {
		await progress.set(0);
		$isValid = false;
		toast.remove();
	});
</script>

<div style={`width: ${$progress}%`} />
<p>{message}</p>

<style>
	div {
		height: 15px;
		background: linear-gradient(to bottom left, #7a00cc, #ff1a75);
	}
	@media (min-width: 500px) {
		div {
			height: 10px;
			background: linear-gradient(to bottom left, #7a00cc, #ff1a75);
		}
	}
	p {
		margin: 0;
	}
</style>
