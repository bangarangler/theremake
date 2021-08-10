<script>
	console.log('hello');
	/* import { onMount } from 'svelte'; */
	import GiFastArrow from 'svelte-icons/gi/GiFastArrow.svelte';
	import plu from '$images/pluaris-dashboard.jpg';
	export let projInfo;
	$: console.log('projectImg', projInfo.projectTitle);
	/* onMount(() => { */
	/* }); */
</script>

<div class="flipCard">
	<div class="side front">
		<div class="img" id="dynamicBG">&nbsp;</div>
		<h4 class="heading"><span class="headingSpan">{projInfo.projectTitle}</span></h4>
		<h4 class="heading"><span class="headingSpan">{projInfo.projectTitle}</span></h4>
		<div class="details">
			<ul>
				{#each projInfo.techUsed as tU}
					<li>{tU}</li>
				{/each}
			</ul>
		</div>
	</div>
	<div class="side back">
		<p>{projInfo.projectTitle}</p>
		<div class="icon"><GiFastArrow /></div>
		<a href={`/projects/${projInfo.slug}`} class="btn">Go to Project</a>
	</div>
</div>

<style>
	.flipCard {
		--height: 15rem;
		perspective: 150rem;
		-moz-perspective: 150rem;
		position: relative;
		height: var(--height);
		margin-left: -90px;
	}
	@media (min-width: 630px) {
		.flipCard {
			margin-left: -85px;
		}
	}
	@media (min-width: 730px) {
		.flipCard {
			margin-left: 25px;
		}
	}
	@media (min-width: 820px) {
		.flipCard {
			margin-left: 55px;
		}
	}
	@media (min-width: 900px) {
		.flipCard {
			margin-left: 55px;
		}
	}
	@media (min-width: 985px) {
		.flipCard {
			margin-left: 45px;
		}
	}
	@media (min-width: 1395px) {
		.flipCard {
			margin-left: 15px;
		}
	}

	.flipCard .side {
		height: var(--height);
		width: 200px;
		transition: all 0.8s ease;
		position: absolute;
		top: 0;
		left: 0;
		backface-visibility: hidden;
		border-radius: 12px;
		overflow: hidden;
		box-shadow: var(--level2);
	}

	.flipCard .front {
		background: var(--cardBG);
	}

	.flipCard .back {
		background: linear-gradient(to right bottom, var(--hotpink), var(--dracPurp));
		transform: rotateY(180deg);
	}

	.flipCard:hover .front {
		transform: rotateY(-180deg);
	}

	.flipCard:hover .back {
		transform: rotateY(0);
	}
	/* @media (hover: hover) and (pointer: fine) { */
	/* 	.flipCard:hover .front { */
	/* 		transform: rotateY(-180deg); */
	/* 	} */
	/* } */
	/* @media (hover: hover) and (pointer: fine) { */
	/* 	.flipCard:hover .back { */
	/* 		transform: rotateY(0); */
	/* 	} */
	/* } */

	.front .img {
		background-image: linear-gradient(to right bottom, var(--hotpink), var(--dracPurp)),
			url('$static/pexels-working.jpg');
		background-size: cover;
		/* background-blend-mode: screen; */
		background-blend-mode: color;
		height: 6rem;
		-webkit-clip-path: polygon(0 0, 100% 0, 100% 85%, 0 100%);
		clip-path: polygon(0 0, 100% 0, 100% 85%, 0 100%);
	}
	.front .heading {
		font-family: var(--slantText);
		font-size: var(--h6);
		font-weight: 300;
		text-transform: uppercase;
		position: absolute;
		top: 1rem;
		right: 0.5rem;
		text-align: right;
		width: 60%;
	}
	.front .headingSpan {
		padding: 0 6px;
		-webkit-box-decoration-break: clone;
		box-decoration-break: clone;
		background-image: linear-gradient(
			to right bottom,
			rgba(255, 26, 117, 0.85),
			rgba(189, 147, 249, 0.85)
		);
	}
	.front .details ul {
		width: 80%;
		margin: 0 auto;
		display: grid;
		grid-template-columns: minmax(1px, 1fr) minmax(1px, 1fr);
		grid-column-gap: 15px;
		margin-top: 15px;
	}
	.front .details li {
		text-align: center;
		font-size: var(--smallTextSize);
		padding: 3px;
		border-bottom: 1px solid var(--dracPurp);
	}
	.back {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: space-between;
		padding: 1rem 1rem 2.5rem;
	}
	.back .icon {
		height: 50px;
		width: 50px;
	}
	.back a {
		text-decoration: none;
		color: var(--p);
		background: var(--headerBG);
		padding: 5px 8px;
		border-radius: 12px;
		transition: all 0.5s ease;
	}
	.back a:hover {
		color: var(--headerBG);
		background: var(--p);
		transform: skewY(5deg);
		border-bottom: 2px dotted var(--darkAquaLightHotPink);
	}
</style>
