<script lang="ts">
	/* import { TEST } from '$lib/Env'; */
	/* <p>Testing... : {TEST}</p> */
	import { fly, fade } from 'svelte/transition';
	/* import { flip } from 'svelte/animate'; */
	import { theme } from '$stores/themeStore';
	import ProjectCard from '$components/ProjectCard/ProjectCard.svelte';
	import { onMount } from 'svelte';
	let visible = false;
	onMount(() => {
		visible = true;
	});
</script>

<section>
	<div class="titleWrapper">
		<h1>Hi. I'm Jon. Fullstack Developer...</h1>
		{#if visible}
			<blockquote in:fade={{ delay: 100, duration: 1200 }} out:fade={{ duration: 100 }}>
				<p
					class="quote"
					in:fly={{ delay: 100, x: -100, opacity: 0, duration: 1200 }}
					out:fade={{ duration: 100 }}
				>
					A Foolish Consistency Is the Hobgoblin of Little Minds
				</p>
				{#if $theme === 'DARK'}
					<div
						class="quoteFooter"
						in:fly={{ delay: 100, x: 100, opacity: 0, duration: 1000 }}
						out:fade={{ duration: 100 }}
					>
						- Ralph Waldo Emerson
					</div>
				{:else}
					<div
						class="quoteFooter light"
						in:fly={{ delay: 100, x: 100, opacity: 0, duration: 1000 }}
						out:fade={{ duration: 100 }}
					>
						- Ralph Walso Emerson
					</div>
				{/if}
			</blockquote>
		{/if}
	</div>
	<div class="line" />
	<div class="blobWrapper">
		<h2>Ponderings and expertise; With Love!</h2>
		<div class="blob" />
	</div>
</section>

<ProjectCard />

<style>
	/* Small devices (landscape phones 576px and up) */
	/* @media (min-width: 576px) { */
	/* } */
	section {
		padding: var(--containerPadding);
		display: grid;
		grid-template-columns: minmax(0, 1fr);
		grid-template-areas:
			'title'
			'line'
			'blob';
		margin-bottom: 125px;
	}
	@media (min-width: 900px) {
		/* section { */
		/* 	grid-template-columns: minmax(0, 1fr) 3px minmax(50vmin, 1fr); */
		/* 	grid-template-areas: */
		/* 		'title title title' */
		/* 		'line line line' */
		/* 		'blob blob blob'; */
		/* } */
	}
	@media (min-width: 1200px) {
		section {
			padding: var(--containerPadding);
			display: grid;
			grid-template-columns: minmax(0, 1fr) 3px minmax(50vmin, 1fr);
			grid-template-areas: 'title line blob';
			margin-bottom: 200px;
		}
	}
	.titleWrapper {
		width: 90%;
		margin: 0 auto;
	}
	@media (min-width: 900px) {
		.titleWrapper {
			width: 70%;
			display: flex;
			align-items: center;
			justify-content: space-evenly;
		}
	}
	@media (min-width: 1200px) {
		.titleWrapper {
			display: grid;
			grid-area: 'title';
			place-self: center;
			justify-self: center;
			width: 60%;
		}
	}

	h1 {
		font-size: var(--h3);
		background: linear-gradient(271deg, var(--hotpink) 50%, var(--aqua) 70%, #a162e8 94%);
		background-clip: border-box;
		-webkit-background-clip: text;
		-webkit-text-fill-color: transparent;
		text-align: center;
		margin-bottom: 35px;
	}
	@media (min-width: 900px) {
		h1 {
			font-size: var(--h2);
		}
	}
	@media (min-width: 1200px) {
		h1 {
			font-size: var(--h1);
			background: linear-gradient(271deg, var(--hotpink) 50%, var(--aqua) 70%, #a162e8 94%);
			background-clip: border-box;
			-webkit-background-clip: text;
			-webkit-text-fill-color: transparent;
			text-align: unset;
		}
	}

	blockquote {
		text-align: center;
		margin-bottom: 35px;
	}

	.quote {
		quotes: '\201C''\201D''\2018''\2019';
	}

	.quote:before {
		content: open-quote;
		display: inline;
		height: 0;
		line-height: 0;
		left: -10px;
		position: relative;
		top: 20px;
		color: var(--hotpink);
		font-size: 3em;
	}

	.quote::after {
		content: close-quote;
		display: inline;
		height: 0;
		line-height: 0;
		left: 10px;
		position: relative;
		top: 25px;
		color: var(--aqua);
		font-size: 3em;
	}

	.quoteFooter {
		margin: 0;
		text-align: right;
		font-style: italic;
		font-family: var(--slantText);
		padding: 0;
		text-align: center;
		background: linear-gradient(to bottom left, #7a00cc, #ff1a75);
		background-clip: border-box;
		-webkit-background-clip: text;
		-webkit-text-fill-color: transparent;
		font-size: 1.3rem;
	}
	@media (min-width: 500px) {
	}
	@media (min-width: 1200px) {
		.quote {
			quotes: '\201C''\201D''\2018''\2019';
		}
	}

	.light {
		-webkit-text-fill-color: unset;
	}

	.line {
		width: 90%;
		min-width: 90%;
		height: 3px;
		min-height: 3px;
		background: var(--hotpink);
		margin: 25px auto 50px auto;
	}

	@media (min-width: 500px) {
	}
	@media (min-width: 1200px) {
		.line {
			width: 3px;
			height: 100%;
			background: var(--hotpink);
		}
	}

	h2 {
		position: absolute;
		top: 21px;
		left: 36px;
		z-index: 1;
		font-size: var(--h3);
		color: var(--textColor);
		width: min-content;
		background: var(--waveBG);
		border-radius: 12px;
		opacity: 0.8;
		padding: 8px;
	}
	@media (min-width: 600px) {
		h2 {
			top: 21px;
			left: 35px;
			width: min-content;
		}
	}
	@media (min-width: 900px) {
		h2 {
			top: 100px;
			left: 140px;
			font-size: var(--h1);
			width: -moz-fit-content;
			width: fit-content;
			background: var(--waveBG);
			border-radius: 12px;
			opacity: 0.8;
			padding: 8px;
		}
	}
	@media (min-width: 1200px) {
		h2 {
			position: absolute;
			top: 24px;
			left: 46px;
			z-index: 1;
			font-size: var(--h2);
			width: min-content;
			background: var(--waveBG);
			border-radius: 12px;
			opacity: 0.8;
			padding: 12px;
		}
	}

	.blobWrapper {
		width: 45vmin;
		height: 45vmin;
		grid-area: 'blob';
		position: relative;
		place-self: center;
	}
	@media (min-width: 900px) {
		.blobWrapper {
			width: 55vmin;
			height: 55vmin;
		}
	}
	@media (min-width: 1200px) {
		.blobWrapper {
			width: 50vmin;
			height: 50vmin;
			grid-area: 'blob';
			position: relative;
			place-self: end;
		}
	}

	.blob {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background: var(--aqua);
		background-image: none;
		background-image: linear-gradient(45deg, #a162e8 15%, var(--aqua) 80%);
		border-radius: 30% 70% 70% 30% / 30% 30% 70% 70%;
		box-shadow: -10vmin 10vmin 0 rgba(255, 255, 255, 0.07);
	}
	@media (min-width: 900px) {
		/* .blob { */
		/* 	box-shadow: -8vmin 8vmin 0 rgba(255, 255, 255, 0.07); */
		/* } */
	}
	@media (min-width: 1200px) {
		.blob {
			position: absolute;
			top: 0;
			left: 0;
			width: 100%;
			height: 100%;
			background: var(--aqua);
			background-image: none;
			background-image: linear-gradient(45deg, #a162e8 15%, var(--aqua) 80%);
			border-radius: 30% 70% 70% 30% / 30% 30% 70% 70%;
			box-shadow: -8vmin 8vmin 0 rgba(255, 255, 255, 0.07);
		}
	}
</style>
