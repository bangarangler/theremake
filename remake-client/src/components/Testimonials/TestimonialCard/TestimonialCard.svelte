<script>
	import carsonImg from '$static/carson.svg';
	import jonBImg from '$static/jon-b-image.svg';
	import ryanImg from '$static/ryan-image.svg';
	import ilyaImg from '$static/undraw-profile.svg';
	import ImageLoader from '$images/ImageLoader.svelte';
	export let testimonial;
	$: ({ firstName, lastName, title, company, reference } = testimonial);
</script>

<div class="testimonialWrapper">
	<span class="quote" />
	<div class="refSection">
		<p>{reference}</p>
	</div>
	<p class="name">{firstName} {lastName}</p>
	<span>{title}: {company}</span>
	{#if firstName === 'Carson'}
		<picture>
			<source media="(min-width:1000px)" srcset={carsonImg} />
			<div>
				<ImageLoader srcset={carsonImg} ty="image/svg" alt="{title} {firstName}" />
			</div>
		</picture>
		<svg class="carson" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1440 320"
			><path
				fill="currentColor"
				fill-opacity="1"
				d="M0,192L34.3,202.7C68.6,213,137,235,206,240C274.3,245,343,235,411,202.7C480,171,549,117,617,96C685.7,75,754,85,823,128C891.4,171,960,245,1029,261.3C1097.1,277,1166,235,1234,197.3C1302.9,160,1371,128,1406,112L1440,96L1440,320L1405.7,320C1371.4,320,1303,320,1234,320C1165.7,320,1097,320,1029,320C960,320,891,320,823,320C754.3,320,686,320,617,320C548.6,320,480,320,411,320C342.9,320,274,320,206,320C137.1,320,69,320,34,320L0,320Z"
			/></svg
		>
	{/if}
	{#if firstName === 'Ryan'}
		<picture>
			<source media="(min-width:500px)" srcset={ryanImg} />
			<div>
				<ImageLoader srcset={ryanImg} ty="image/svg" alt="{title} {firstName}" />
			</div>
		</picture>
		<svg class="ryan" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1440 320"
			><path
				fill="currentColor"
				fill-opacity="1"
				d="M0,192L34.3,202.7C68.6,213,137,235,206,240C274.3,245,343,235,411,202.7C480,171,549,117,617,96C685.7,75,754,85,823,128C891.4,171,960,245,1029,261.3C1097.1,277,1166,235,1234,197.3C1302.9,160,1371,128,1406,112L1440,96L1440,320L1405.7,320C1371.4,320,1303,320,1234,320C1165.7,320,1097,320,1029,320C960,320,891,320,823,320C754.3,320,686,320,617,320C548.6,320,480,320,411,320C342.9,320,274,320,206,320C137.1,320,69,320,34,320L0,320Z"
			/></svg
		>
	{/if}
	{#if firstName === 'Ilya'}
		<picture>
			<source media="(min-width:500px)" srcset={jonBImg} />
			<div>
				<ImageLoader srcset={ilyaImg} ty="image/svg" alt="{title} {firstName}" />
			</div>
		</picture>
		<svg class="ilya" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1440 320"
			><path
				fill="currentColor"
				fill-opacity="1"
				d="M0,192L34.3,202.7C68.6,213,137,235,206,240C274.3,245,343,235,411,202.7C480,171,549,117,617,96C685.7,75,754,85,823,128C891.4,171,960,245,1029,261.3C1097.1,277,1166,235,1234,197.3C1302.9,160,1371,128,1406,112L1440,96L1440,320L1405.7,320C1371.4,320,1303,320,1234,320C1165.7,320,1097,320,1029,320C960,320,891,320,823,320C754.3,320,686,320,617,320C548.6,320,480,320,411,320C342.9,320,274,320,206,320C137.1,320,69,320,34,320L0,320Z"
			/></svg
		>
	{/if}
	{#if firstName === 'Jonathan'}
		<picture>
			<source media="(min-width:500px)" srcset={jonBImg} />
			<div>
				<ImageLoader srcset={jonBImg} ty="image/svg" alt="{title} {firstName}" />
			</div>
		</picture>
		<svg class="jonb" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1440 320"
			><path
				fill="currentColor"
				fill-opacity="1"
				d="M0,192L34.3,202.7C68.6,213,137,235,206,240C274.3,245,343,235,411,202.7C480,171,549,117,617,96C685.7,75,754,85,823,128C891.4,171,960,245,1029,261.3C1097.1,277,1166,235,1234,197.3C1302.9,160,1371,128,1406,112L1440,96L1440,320L1405.7,320C1371.4,320,1303,320,1234,320C1165.7,320,1097,320,1029,320C960,320,891,320,823,320C754.3,320,686,320,617,320C548.6,320,480,320,411,320C342.9,320,274,320,206,320C137.1,320,69,320,34,320L0,320Z"
			/></svg
		>
	{/if}
	{#if firstName === 'Shawn'}
		<picture>
			<source media="(min-width:500px)" srcset={ilyaImg} />
			<div>
				<ImageLoader srcset={ilyaImg} ty="image/svg" alt="{title} {firstName}" />
			</div>
		</picture>
		<svg class="shawn" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1440 320"
			><path
				fill="currentColor"
				fill-opacity="1"
				d="M0,192L34.3,202.7C68.6,213,137,235,206,240C274.3,245,343,235,411,202.7C480,171,549,117,617,96C685.7,75,754,85,823,128C891.4,171,960,245,1029,261.3C1097.1,277,1166,235,1234,197.3C1302.9,160,1371,128,1406,112L1440,96L1440,320L1405.7,320C1371.4,320,1303,320,1234,320C1165.7,320,1097,320,1029,320C960,320,891,320,823,320C754.3,320,686,320,617,320C548.6,320,480,320,411,320C342.9,320,274,320,206,320C137.1,320,69,320,34,320L0,320Z"
			/></svg
		>
	{/if}
	{#if firstName === 'Shridatta'}
		<picture>
			<source media="(min-width:500px)" srcset={ilyaImg} />
			<div>
				<ImageLoader srcset={ilyaImg} ty="image/svg" alt="{title} {firstName}" />
			</div>
		</picture>
		<svg class="shridatta" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1440 320"
			><path
				fill="currentColor"
				fill-opacity="1"
				d="M0,192L34.3,202.7C68.6,213,137,235,206,240C274.3,245,343,235,411,202.7C480,171,549,117,617,96C685.7,75,754,85,823,128C891.4,171,960,245,1029,261.3C1097.1,277,1166,235,1234,197.3C1302.9,160,1371,128,1406,112L1440,96L1440,320L1405.7,320C1371.4,320,1303,320,1234,320C1165.7,320,1097,320,1029,320C960,320,891,320,823,320C754.3,320,686,320,617,320C548.6,320,480,320,411,320C342.9,320,274,320,206,320C137.1,320,69,320,34,320L0,320Z"
			/></svg
		>
	{/if}
</div>

<style>
	.testimonialWrapper {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		margin-bottom: 15px;
	}
	@media (min-width: 910px) {
		.testimonialWrapper {
			margin-bottom: 0px;
		}
	}
	picture div {
		max-width: 100px;
		max-height: 100px;
		border-radius: 50%;
		overflow: hidden;
	}
	.ryan {
		color: var(--aqua);
	}
	.carson {
		color: var(--dracPurp);
	}
	.ilya {
		color: var(--hotpink);
	}
	.jonb {
		color: var(--dracPurp);
	}
	.shawn {
		color: var(--aqua);
	}
	.shridatta {
		color: var(--hotpink);
	}
	.name {
		font-size: var(--h4);
		font-family: var(--slantText);
		letter-spacing: 3px;
		background: linear-gradient(to bottom left, #7a00cc, #ff1a75);
		background-clip: border-box;
		-webkit-background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.quote {
		quotes: '\201C''\201D''\2018''\2019';
		margin-bottom: 12px;
		align-self: flex-start;
	}
	.refSection {
		margin-bottom: 30px;
		background: var(--cardBG);
		border-radius: 12px;
		padding: 1em;
		max-height: 400px;
		overflow: auto;
	}
	.refSection p {
		margin-bottom: 0;
	}
	.testimonialWrapper span {
		font-size: 0.9em;
		/* color: var(--darkAquaLightHotPink); */
		margin-bottom: 30px;
		text-align: center;
		font-weight: bold;
		font-style: italic;
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
		color: transparent;
		font-size: 3em;
	}
</style>
