<script>
	import { page } from '$app/stores';
	/* import { goto } from '$app/navigation'; */
	import { projectMetaData } from '$stores/projectMetaData';
	import IoIosArrowRoundBack from 'svelte-icons/io/IoIosArrowRoundBack.svelte';
	import IoIosArrowRoundForward from 'svelte-icons/io/IoIosArrowRoundForward.svelte';
	import joyride from '$images/undraw-joyride.svg';
	/* let pages = $projectMetaData.map((p) => { */
	/* 	return p.slug; */
	/* }); */
	/* let activePage = pages.filter((p) => $page.path.includes(p)); */
	/* console.log('pages', pages); */
	/* console.log('activePage', activePage); */
	/* function getNext(list, current) { */
	/* 	console.log('list', list); */
	/* 	console.log('current', current); */
	/* 	let currentIdx = list.indexOf(current); */
	/* 	console.log('currentIdx', currentIdx); */
	/* 	if (list) */
	/* } */
	/* getNext(pages, activePage[0]); */
	// have array
	// need lenght of array
	// click button and change
	/* [1, 2*, 3] */
	/* [1, 2, 3*] */
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
