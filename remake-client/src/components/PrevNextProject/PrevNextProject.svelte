<!-- https://ar.al/2021/04/03/passing-data-from-layouts-to-pages-in-sveltekit/-->
<script>
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { projectMetaData } from '$stores/projectMetaData';
	import IoIosArrowRoundBack from 'svelte-icons/io/IoIosArrowRoundBack.svelte';
	import IoIosArrowRoundForward from 'svelte-icons/io/IoIosArrowRoundForward.svelte';
	import joyride from '$images/undraw-joyride.svg';
	/*$: console.log('page.path', $page.path);*/
	/* const test = $page.path.indexOf(md?.slug) ? 'activeProj' : null */
	/* const test = $page.path.includes('test'); */
	/* console.log('test', test); */
	/* on:click={() => goto(`/projects/${md?.slug}`)} */
	/* small store issue, if they try to go straight to a project page or reload
	 * on one the store is empty because it's set in the indeex page... need to see
	 * if that data can be set server side as well so that if they go to a page that
	 * information is availabe */
</script>

<!--<script context="module">
	let urlData;
	export async function load({ context }) {
		urlData = await context.urlData;
		console.log('urlData here ; )', urlData);
		return true;
	}
</script>-->



<div class="prevNextContainer">
	<div class="prev">
		<span class="arrowIcons"><IoIosArrowRoundBack /></span>
		<p>prev</p>
	</div>
	<div class="svgContainer">
		<img class="svg" src={joyride} alt="arrow svg art from undraw" />
	</div>
	<div class="next" on:click={() => goto(`/projects/`)}>
		<span class="arrowIcons"><IoIosArrowRoundForward /></span>
		<p>next</p>
	</div>
	<div class="projectNav">
		{#each $projectMetaData as md}
			<a href={`/projects/${md?.slug}`} class={$page.path.includes(md?.slug) ? 'activeProj' : ''}
				>{md?.projectTitle}</a
			>
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
			'prev svg svg next'
			'projectNav projectNav projectNav projectNav';
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
	.projectNav {
		grid-area: projectNav;
		width: 100%;
		display: grid;
		grid-template-columns: repeat(3, minmax(0, 1fr));
		grid-template-areas: 'prevProj activeProj nextProj';
		place-items: center;
	}
	.projectNav a {
		color: var(--aqua);
		text-decoration: none;
	}
	.activeProj {
		grid-area: activeProj;
		font-family: var(--slantText);
		font-weight: bold;
		font-size: var(--h5);
		letter-spacing: 0.2rem;
		color: var(--hotpink) !important;
	}
</style>
