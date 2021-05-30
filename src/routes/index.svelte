<script lang="ts">
	import { TEST } from '$lib/Env';
	/* <p>Testing... : {TEST}</p> */
	import Wave from '$components/Wave/Wave.svelte';
</script>

<section>
	<div class="titleWrapper">
		<h1>Hi. I'm Jon. Fullstack Developer...</h1>
	</div>
	<div class="line" />
	<div class="blobWrapper">
		<h2>Ponderings and expertise; With Love!</h2>
		<div class="blob" />
	</div>
</section>

<Wave>
	<div class="testing">
		<p class="remove-margin">show me this thing</p>
	</div>
</Wave>

<style>
	.testing {
		background: var(--waveBG);
		min-width: 100%;
		min-height: 100%;
	}
	.remove-margin {
		margin-bottom: 0;
	}
	section {
		padding: var(--containerPadding);
		display: grid;
		grid-template-columns: minmax(0, 1fr) 3px minmax(50vmin, 1fr);
		grid-template-areas: 'title line blob';
		margin-bottom: 200px;
	}
	.titleWrapper {
		grid-area: 'title';
		place-self: center;
		justify-self: center;
		width: 60%;
	}
	h1 {
		background: linear-gradient(271deg, var(--hotpink) 50%, var(--aqua) 70%, #a162e8 94%);
		background-clip: border-box;
		-webkit-background-clip: text;
		-webkit-text-fill-color: transparent;
		/* width: 90%; */
		/* border: 1px solid red; */
		/* grid-area: 'title'; */
	}
	.line {
		width: 3px;
		height: 100%;
		background: var(--hotpink);
	}
	h2 {
		position: absolute;
		top: 115px;
		left: -56px;
		z-index: 1;
	}
	.blobWrapper {
		width: 50vmin;
		height: 50vmin;
		grid-area: 'blob';
		position: relative;
		place-self: end;
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
</style>
