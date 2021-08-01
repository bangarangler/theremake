<script>
	import { toast, isValid } from '$stores/toastStore';
	import Portal from '$components/Portal/Portal.svelte';
	import { fly, fade } from 'svelte/transition';
	import { flip } from 'svelte/animate';
	import ToastMessage from '$components/Toast/ToastMessage.svelte';
	export let duration = 1000;
</script>

<Portal>
	<div class="toastWrapper">
		{#each $toast as message (message)}
			<div
				on:click={() => {
					toast.remove();
					$isValid = false;
				}}
				in:fly={{ opacity: 0, x: 100 }}
				out:fade
				animate:flip
				class="toast"
			>
				<ToastMessage {message} {duration} />
			</div>
		{/each}
	</div>
</Portal>

<style>
	.toastWrapper {
		position: fixed;
		top: 100px;
		right: 20px;
	}
	.toast {
		margin-bottom: 1rem;
		padding: 20px;
		border-radius: 12px;
		background: var(--cardBGGradient);
		box-shadow: var(--level2);
	}
</style>
