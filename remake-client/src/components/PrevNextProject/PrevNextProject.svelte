<script>
	import { page } from '$app/stores';
	import { projectMetaData, previousAndNextProject } from '$stores/projectMetaData';
	import IoIosArrowRoundBack from 'svelte-icons/io/IoIosArrowRoundBack.svelte';
	import IoIosArrowRoundForward from 'svelte-icons/io/IoIosArrowRoundForward.svelte';
	import joyride from '$images/undraw-joyride.svg';
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
</script>

<div class="prevNextContainer">
	<a
		on:click={(e) => e.stopPropagation()}
		class="prev {!prevPage && 'disabled'}"
		disabled={!prevPage}
		href={`/projects/${prevPage}/`}
		><span class="arrowIcons"><IoIosArrowRoundBack /></span>
		<p>prev</p></a
	>
	<div class="svgContainer">
		<img class="svg" src={joyride} alt="arrow svg art from undraw" />
	</div>
	<a
		on:click={(e) => e.stopPropagation()}
		class="next {!nextPage && 'disabled'}"
		disabled={!nextPage}
		href={`/projects/${nextPage}/`}
		><span class="arrowIcons"><IoIosArrowRoundForward /></span>
		<p>next</p></a
	>
	<nav class="projectNav">
		{#each $projectMetaData as md}
			<a
				on:click={(e) => e.stopPropagation()}
				href={`/projects/${md.slug}/`}
				rel="noopener"
				class={$page.path.includes(md.slug)
					? 'activeProj'
					: md.slug === nextPage
					? 'nextPage'
					: md.slug === prevPage
					? 'prevPage'
					: ''}>{md.projectTitle}</a
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
		width: 200px;
	}
	.svg {
		height: 100%;
		width: 100%;
		object-fit: contain;
	}
	@media (min-width: 1000px) {
		.svg {
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
		cursor: not-allowed;
		color: var(--lightGray);
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
		width: 125px;
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
		text-align: center;
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
		font-size: var(--h5);
		letter-spacing: 0.2rem;
		color: var(--hotpink) !important;
	}
	.nextPage {
		grid-area: nextProj;
	}
	.prevPage {
		grid-area: prevProj;
	}
</style>
