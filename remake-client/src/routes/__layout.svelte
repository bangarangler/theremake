<script context="module">
	export async function load() {
		const projects = import.meta.globEager('../projects/*.md');
		const projectList = Object.values(projects);
		//console.log('projectList', projectList);
		// get the metaData
		const projectMeta = projectList?.map((p) => {
			return p?.metadata;
		});
		//console.log('projectMeta', projectMeta);
		let prevNext = [];
		projectMeta.forEach((p, i) => {
			const current = projectMeta[i];
			// @ts-ignore
			const prev = projectMeta[i - 1];
			/* console.log('prev', prev); */
			// @ts-ignore
			const nxt = projectMeta[i + 1];
			/* console.log('next', nxt); */
			const nxtNPrev = {
				current: current ? current?.slug : '',
				previous: prev ? prev?.slug : '',
				next: nxt ? nxt?.slug : ''
			};
			/* console.log('nxtNPrev', nxtNPrev); */
			prevNext.push(nxtNPrev);
		});
		/* projectMeta.prevNextt = prevNext; */
		/* console.log('prevNext', prevNext); */
		return {
			props: {
				metaData: projectMeta,
				prevNext: prevNext
			}
		};
	}
</script>

<script lang="ts">
	import { page } from '$app/stores';
	import { projectMetaData, previousAndNextProject } from '$stores/projectMetaData';
	import type { ProjectMetaDataType } from '$stores/projectMetaData';
	import NavBar from '$components/NavBar/NavBar.svelte';
	import Footer from '$components/Footer/Footer.svelte';
	$: path = $page.path.replace('/', '');
	$: title = path !== '' ? path : 'home';

	export let metaData: ProjectMetaDataType[];
	export let prevNext: any;
	if (prevNext) {
		previousAndNextProject.update(() => prevNext);
	}
	if (metaData) {
		projectMetaData.update(() => metaData);
	}
	//$: console.log('projectMetaData from allergenguardian', $projectMetaData);
	//$: console.log('previousAndNextProject from allergenguardian', $previousAndNextProject);
</script>

<svelte:head>
	<title>{title.charAt(0).toUpperCase() + title.slice(1)} | Jonathan Dain Palacio</title>
	<meta
		name="description"
		content="Hi I'm Jon! I'm a web developer with a extreme interest in marketing and business! I've owned a few business's in my time and I still enjoy the thrill of planning and organizing a business. Regardless if it's day one or you have been in business for generations; I'm confident that I can make it even better!"
	/>
	<meta
		name="keywords"
		content="coding, developer, react, software,
		 portfolio, software developer, web developer, fullstack developer"
	/>
	<meta name="author" content="Jon Palacio" />
	<meta property="og:title" content="jonathandain.dev" />
	<meta
		property="og:description"
		content="
Hi I'm Jon! I'm a web developer with a extreme interest in marketing and
business! I've owned a few business's in my time and I still enjoy the thrill of
planning and organizing a business. Regardless if it's day one or you have been
in business for generations; I'm confident that I can make it even better!"
	/>
	<meta property="og:type" content="website" />
	<meta property="og:url" content="https://jonathandain.dev" />
	<!--<meta property="og:image" content="" />-->
	<meta name="twitter:card" content="summary" />
	<meta name="twitter:creator" content="Jon Palacio" />
	<meta name="twitter:title" content={title} />
	<meta
		name="twitter:description"
		content="Hi I'm Jon! I'm a web developer with a extreme interest in marketing and business! I've owned a few business's in my time and I still enjoy the thrill of planning and organizing a business. Regardless if it's day one or you have been in business for generations; I'm confident that I can make it even better!"
	/>
	<meta property="og:site_name" content="jonathandain.dev" />
	<meta
		name="twitter:image:alt"
		content="displaying website image for
jonathandain.dev"
	/>
</svelte:head>

<main class="layout">
	<NavBar />
	<slot />
</main>

<div class="fancyDivide" />

<Footer />

<style>
	.fancyDivide {
		width: 100%;
		height: 3px;
		background: linear-gradient(to left, #7a00cc, var(--aqua));
	}
</style>
