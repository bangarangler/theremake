<script>
	import { projectMetaData } from '$stores/projectMetaData';
	import IoIosArrowRoundBack from 'svelte-icons/io/IoIosArrowRoundBack.svelte';
	import IoIosArrowRoundForward from 'svelte-icons/io/IoIosArrowRoundForward.svelte';
	import joyride from '$images/undraw-joyride.svg';
</script>

<div class="prevNextContainer">
	<div class="prev">
		<span class="arrowIcons"><IoIosArrowRoundBack /></span>
		<p>prev</p>
	</div>
	<div class="svgContainer">
		<img class="svg" src={joyride} alt="arrow svg art from undraw" />
	</div>
	<div class="next">
		<span class="arrowIcons"><IoIosArrowRoundForward /></span>
		<p>next</p>
	</div>
	<div class="projectNav">
		{#each $projectMetaData as md}
			<a href={`/projects/${md.slug}`}>{md.projectTitle}</a>
		{/each}
	</div>
</div>

<style>
	.prevNextContainer {
		width: 800px;
		height: 400px;
		min-width: 800px;
		min-height: 400px;
		border: 3px solid var(--darkGray);
		border-radius: 12px;
		margin: 100px auto 0 auto;
		display: grid;
		grid-template-columns: repeat(4, minmax(0, 1fr));
		place-items: center;
		grid-template-areas:
			'prev svg svg next'
			'prev svg svg next';
	}
	.prev {
		grid-area: prev;
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		cursor: pointer;
	}
	.svgContainer {
		grid-area: svg;
	}
	.svg {
		max-height: 200px;
		max-width: 300px;
	}
	.next {
		grid-area: next;
		display: flex;
		flex-direction: column;
		align-items: flex-end;
		cursor: pointer;
	}
	.arrowIcons {
		display: block;
		width: 100px;
		height: 100px;
	}
	.projectNav a {
		text-decoration: none;
	}
</style>
