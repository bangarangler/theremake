<script>
	import { page } from '$app/stores';
	import { projectMetaData, previousAndNextProject } from '$stores/projectMetaData';
	import ImageLoader from '$images/ImageLoader.svelte';
	import IoIosArrowRoundBack from 'svelte-icons/io/IoIosArrowRoundBack.svelte';
	import IoIosArrowRoundForward from 'svelte-icons/io/IoIosArrowRoundForward.svelte';
	/* import setUp from '$static/imgForUses.png?w=300;600;1000&format=png&srcset'; */
	import joyride from '$images/undraw-joyride.svg?w=300;600;800&format=svg&srcset';
	$: activePage = $projectMetaData.filter((p) => $page.path.includes(p.slug))[0];
	let nextPage;
	let prevPage;
	function getPrevNext(prevNext, slug) {
		for (let i = 0; i < prevNext.length; i++) {
			if (prevNext[i].current === slug) {
				nextPage = prevNext[i].next;
				prevPage = prevNext[i].previous;
				return;
			}
		}
	}
	$: getPrevNext($previousAndNextProject, activePage.slug);
	$: console.log('projectMetaData', $projectMetaData);
</script>

<div class="prevNextContainer">
	<a
		on:click={(e) => e.stopPropagation()}
		class="prev {!prevPage && 'disabled'}"
		disabled={!prevPage}
		sveltekit:prefetch
		href={`/projects/${prevPage}/`}
		><span class="arrowIcons"><IoIosArrowRoundBack /></span>
		<p>prev</p></a
	>
	<div class="svgContainer">
		<ImageLoader srcset={joyride} ty="image/svg" alt="arrow svg art from undraw" />
	</div>
	<a
		on:click={(e) => e.stopPropagation()}
		class="next {!nextPage && 'disabled'}"
		disabled={!nextPage}
		sveltekit:prefetch
		href={`/projects/${nextPage}/`}
		><span class="arrowIcons"><IoIosArrowRoundForward /></span>
		<p>next</p></a
	>
	<nav class="projectNav">
		{#each $projectMetaData as md}
			<a
				on:click={(e) => e.stopPropagation()}
				sveltekit:prefetch
				href={`/projects/${md.slug}/`}
				rel="noopener"
				class={$page.path.includes(md.slug)
					? 'activeProj'
					: md.slug === nextPage
					? 'nextPage'
					: md.slug === prevPage
					? 'prevPage'
					: 'nope'}>{md.projectTitle}</a
			>
		{/each}
	</nav>
</div>

<style>
	.prevNextContainer {
		width: 90%;
		/*height: 400px; */
		min-width: 90%;
		/*min-height: 400px;*/
		border: 3px solid var(--darkGray);
		border-radius: 12px;
		margin: 100px auto 0 auto;
		padding: 15px;
		display: grid;
		grid-template-columns: repeat(4, minmax(0, 1fr));
		place-items: center;
		grid-template-areas:
			'prev svg svg next'
			'prev svg svg next'
			'projectNav projectNav projectNav projectNav';
		margin-bottom: 100px;
	}

	@media (min-width: 800px) {
		.prevNextContainer {
			width: 80%;
			min-width: 80%;
		}
	}
	@media (min-width: 1000px) {
		.prevNextContainer {
			width: 800px;
			height: 400px;
			min-width: 800px;
			min-height: 400px;
			padding: unset;
			grid-template-areas:
				'prev svg svg next'
				'prev svg svg next'
				'projectNav projectNav projectNav projectNav';
		}
	}
	.prev {
		grid-area: prev;
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		cursor: pointer;
		color: inherit;
		text-decoration: none;
	}
	.svgContainer {
		grid-area: svg;
		max-width: 200px;
	}
	@media (min-width: 1000px) {
		.svgContainer {
			max-height: 200px;
			max-width: 300px;
		}
	}
	.next {
		grid-area: next;
		display: flex;
		flex-direction: column;
		align-items: flex-end;
		cursor: pointer;
		background: none;
		color: inherit;
		text-decoration: none;
	}
	.disabled {
		color: var(--projH6DarkLight);
		pointer-events: none;
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
		width: 100px;
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
		text-align: center;
		font-size: var(--smallTextSize);
	}
	@media (min-width: 500px) {
		.projectNav a {
			width: 125px;
			font-size: var(--bodyFont);
		}
	}
	@media (min-width: 600px) {
		.projectNav a {
			width: unset;
		}
	}
	.activeProj {
		grid-area: activeProj;
		font-family: var(--slantText);
		font-weight: bold;
		font-size: var(--h6) !important;
		letter-spacing: 0.2rem;
		color: var(--hotpink) !important;
	}
	@media (min-width: 500px) {
		.activeProj {
			font-size: var(--h5) !important;
		}
	}
	.nextPage {
		grid-area: nextProj;
	}
	.prevPage {
		grid-area: prevProj;
	}
	.nope {
		display: none;
	}
</style>
