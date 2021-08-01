<script>
	import { fly, fade } from 'svelte/transition';
	import Portal from '$components/Portal/Portal.svelte';

	export let isModalOpen = false;

	function closeModal() {
		isModalOpen = false;
	}
</script>

{#if isModalOpen}
	<Portal>
		<div class="modalWrapper" transition:fly={{ opacity: 0, y: 100 }}>
			<button on:click={closeModal} aria-label="Close Modal!">Close</button>
			<slot />
		</div>
		<div class="modalBG" on:click={closeModal} transition:fade />
	</Portal>
{/if}

<style>
	.modalWrapper {
		position: fixed;
		inset: 25vh 0 0;
		/* min-width: 320px; */
		max-width: 530px;
		margin: 0 auto;
		width: 80%;
		z-index: 101;
		border-radius: 12px;
		max-height: 300px;
		padding: 1em;
		background: var(--cardBGGradient);
		display: grid;
		grid-template-rows: 20px auto;
		grid-row-gap: 10px;
	}
	@media (min-width: 620px) {
		.modalWrapper {
			width: 100%;
		}
	}
	.modalWrapper button {
		width: max-content;
		place-self: end;
	}
	.modalBG {
		background: var(--black);
		opacity: 0.8;
		cursor: pointer;
		position: fixed;
		inset: 0;
		z-index: 100;
	}
</style>
