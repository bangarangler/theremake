<script>
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { projectMetaData, previousAndNextProject } from '$stores/projectMetaData';
	import IoIosArrowRoundBack from 'svelte-icons/io/IoIosArrowRoundBack.svelte';
	import IoIosArrowRoundForward from 'svelte-icons/io/IoIosArrowRoundForward.svelte';
	import joyride from '$images/undraw-joyride.svg';
	$: console.log('previousAndNextProject from PrevNextProject', $previousAndNextProject);
	$: activePage = $projectMetaData.filter((p) => $page.path.includes(p.slug))[0];
	console.log('activePage HERE', activePage);
	let nextPage;
	let prevPage;
	function getPrevNext(prevNext, slug) {
		/* console.log('prevNexxt', prevNext); */
		/* console.log('slug from func', slug); */
		for (let i = 0; i < prevNext.length; i++) {
			/* console.log('prevNext[i]', prevNext[i]); */
			/* console.log('slug in for loop', slug); */
			if (prevNext[i].current === slug) {
				nextPage = prevNext[i].next;
				prevPage = prevNext[i].previous;
				return;
			}
		}
	}
	$: getPrevNext($previousAndNextProject, activePage.slug);
	$: console.log('nextPage reactive', nextPage);
	$: console.log('prevPage reactive', prevPage);
</script>

<div class="prevNextContainer">
	<button
		class="prev {!prevPage && 'disabled'}"
		on:click={() => goto(`/projects/${prevPage}`)}
		disabled={!prevPage}
	>
		<span class="arrowIcons"><IoIosArrowRoundBack /></span>
		<p>prev</p>
	</button>
	<div class="svgContainer">
		<img class="svg" src={joyride} alt="arrow svg art from undraw" />
	</div>
	<button
		class="next {!nextPage && 'disabled'}"
		on:click={() => goto(`/projects/${nextPage}`)}
		disabled={!nextPage}
	>
		<span class="arrowIcons"><IoIosArrowRoundForward /></span>
		<p>next</p>
	</button>
	<nav class="projectNav">
		{#each $projectMetaData as md}
			<a
				on:click={(e) => e.stopPropagation()}
				href={`/projects/${md.slug}/`}
				rel="noopener"
				class={$page.path.includes(md.slug) ? 'activeProj' : ''}>{md.projectTitle}</a
			>
		{/each}
	</nav>
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
		background: none;
	}
	.disabled {
		cursor: not-allowed;
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
		background: none;
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
