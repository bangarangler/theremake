<script context="module">
	/* export const prerender = true; */
</script>

<section>
	<h1>Invisible Blog</h1>
	<span>Ok. not Invisible.</span>
	<span>Certainly improbable...</span>
	<div class="card wrapper">
		<picture class="imgWrapper">
			<source
				media="(min-width:1000px)"
				srcset={'https://media.giphy.com/media/l3E6F9xoxTabw3JYY/source.gif'}
			/>
			<img
				class="img1"
				srcset={'https://media.giphy.com/media/l3E6F9xoxTabw3JYY/source.gif'}
				type="image/gif"
				alt="Sillicon Valley... funny moment"
			/>
		</picture>
	</div>
</section>

<style>
	section {
		padding-bottom: 4em;
	}
	h1 {
		text-align: center;
		color: var(--darkAquaLightHotPink);
	}
	span {
		text-align: center;
		display: block;
	}
	.wrapper {
		display: grid;
		grid-template-columns: minmax(0, 1fr);
		margin: 25px auto;
		width: 90%;
		padding: 0.5em;
	}
	.imgWrapper img {
		width: 100%;
		height: 100%;
		object-fit: contain;
	}
	@media (min-width: 1200px) {
		.wrapper {
			max-width: 1000px;
			margin: 25px auto;
			padding: 2em;
		}
		.imgWrapper img {
			width: 100%;
			height: 100%;
			object-fit: contain;
		}
	}
</style>
