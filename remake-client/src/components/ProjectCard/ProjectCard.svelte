<script lang="ts">
	import { projectMetaData } from '$stores/projectMetaData';
  //import type { ProjectMetaDataType } from '$stores/projectMetaData';
	import { paginate } from '$components/Testimonials/testimonialPagination';
	import Wave from '$components/Wave/Wave.svelte';
	import TinyFlipCard from '$components/ProjectCard/TinyFlipCard.svelte';
	//import Wave from '$lib/Wave/Wave.svelte';
	import ProjectCardHeading from '$components/ProjectCard/ProjectCardHeading/ProjectCardHeading.svelte';
	import eyeball from '$images/undraw-surveillance.svg?w=200;&format=svg&srcset';
	export let innerWidth = null;
	let toLargeText = 'Screen is getting kinda large there...';
	$: if (innerWidth >= 1800) toLargeText = 'No Seriously; How wide does this go?';
	$: if (innerWidth < 1700) toLargeText = 'Screen is getting kinda large there...';
	$: if (innerWidth >= 1900) toLargeText = "That's it. No more. good day ; )";
	export let stopRotation = false;
	$: if (innerWidth < 500) stopRotation = true;
	$: if (innerWidth > 500) stopRotation = false;
	$: projInfo = paginate($projectMetaData);
</script>

<svelte:window bind:innerWidth />

{#each projInfo.data.slice(0, 3) as md}
	<div class="maxCutOff">
		<Wave>
			<div class="card wrapper">
				<div class="titleSection">
					<ProjectCardHeading
						projectTitle={md?.projectTitle}
						projectRole={md?.projectRole}
						slug={md?.slug}
					/>
				</div>
				<aside class="card descriptionSection">
					<h5 class="description__text">Description:</h5>
					<p class="projectDescription__text">
						{md?.projectDescription}
					</p>
				</aside>
				<aside class="card techUsedSection">
					<h5 class="description__text">Tech Used:</h5>
					{#each md?.techUsed as techUsed}
						<p class="projectDescription__text">
							{techUsed}
						</p>
					{/each}
				</aside>
				<div class="fancyBoxWrapper cardSection">
					<div class="fancyBox" class:stopRotation>
						<span />
						<h2><a href={`/projects/${md?.slug}`}>{md?.projectTitle}</a></h2>
					</div>
				</div>
				<div class="toLarge">
					<picture
						><img
							srcset={eyeball}
							type="image/svg"
							alt="Eyeball from
				https://undraw.co"
						/></picture
					>
					<p>{toLargeText}</p>
				</div>
			</div>
		</Wave>
	</div>
{/each}

<div class="gridWrap">
	{#each projInfo.data.slice(3, projInfo.data.length) as md}
		<TinyFlipCard projInfo={md} />
	{/each}
</div>

<style>
	.wrapper {
		border-radius: 0; /* overwrite card border-radis for use with wave */
		display: grid;
		grid-row-gap: 0;
		grid-template-columns: minmax(0px, 1fr);
		grid-template-rows:
			minmax(0px, 200px) minmax(0px, auto) minmax(0px, 1fr)
			minmax(0, auto);
		grid-template-areas:
			'title'
			'description'
			'card'
			'techUsed';
	}
	@media (min-width: 1100px) {
		.wrapper {
			border-radius: 0; /* overwrite card border-radis for use with wave */
			display: grid;
			grid-row-gap: 35px;
			grid-template-columns: minmax(0px, 2fr) minmax(0px, 1fr);
			grid-template-rows: minmax(0px, 250px) minmax(0px, 1fr);
			grid-template-areas:
				'description title'
				'card techUsed';
		}
	}
	@media (min-width: 1200px) {
		.wrapper {
			border-radius: 0; /* overwrite card border-radis for use with wave */
			display: grid;
			/* grid-row-gap: 65px; */
			grid-template-columns: minmax(0px, 1fr) minmax(0px, 1fr);
			grid-template-rows: minmax(0px, 1fr) minmax(0px, 1fr) minmax(0px, 1fr);
			grid-template-areas:
				'description description title'
				'techUsed card .'
				'techUsed card .';
		}
	}
	.toLarge {
		grid-area: toLarge;
		display: none;
	}
	@media (min-width: 1700px) {
		.wrapper {
			grid-template-columns: repeat(3, minmax(0, 1fr));
			grid-template-rows: minmax(0px, 1fr) minmax(0px, 1fr) minmax(0px, 1fr);
			grid-template-areas:
				'description description title'
				'techUsed card toLarge'
				'techUsed card toLarge';
			border-radius: 12px;
			margin-bottom: 150px;
		}
		@media (min-width: 1900px) {
			.wrapper {
				border-radius: 12px;
				margin-bottom: 150px;
			}
		}
		.toLarge {
			grid-area: toLarge;
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			font-family: var(--slantText);
			position: relative;
		}
		.toLarge img {
			max-width: 200px;
			margin-bottom: 15px;
		}
		.toLarge p {
			background: linear-gradient(271deg, #a162e8 50%, var(--aqua) 70%, var(--hotpink) 94%);
			background-clip: border-box;
			-webkit-background-clip: text;
			-webkit-text-fill-color: transparent;
		}
	}
	@media (min-width: 1900px) {
		.maxCutOff {
			max-width: 1700px;
			margin: 0 auto;
		}
		.toLarge p {
			transform: rotate(90deg);
			background: unset;
			background-clip: unset;
			-webkit-background-clip: unset;
			-webkit-text-fill-color: unset;
			font-size: var(--h2);
			position: fixed;
			top: 470px;
			right: -270px;
		}
	}
	.titleSection {
		grid-area: title;
		align-self: center;
		justify-self: end;
	}
	.descriptionSection {
		grid-area: description;
		margin-bottom: 100px;
	}
	.techUsedSection {
		grid-area: techUsed;
		height: -moz-fit-content;
		height: fit-content;
		display: flex;
		align-items: center;
		justify-content: space-evenly;
		width: 100%;
		margin-top: 100px;
		flex-wrap: wrap;
		gap: 5px;
	}

	.techUsedSection .description__text {
		margin-bottom: 0px;
	}
	.techUsedSection .projectDescription__text {
		margin-bottom: 0px;
	}
	.cardSection {
		grid-area: card;
	}
	@media (min-width: 1100px) {
		.titleSection {
			align-self: unset;
		}
		.descriptionSection {
			align-self: center;
			justify-self: start;
		}
		.techUsedSection {
			justify-self: end;
			display: block;
			width: auto;
			margin-top: 0px;
		}
		.techUsedSection .description__text {
			margin: 0 0 1.05rem;
		}
		.techUsedSection .projectDescription__text {
			margin-bottom: 1.05rem;
		}
	}
	@media (min-width: 1200px) {
		.descriptionSection {
			/* grid-area: description; */
			height: -moz-fit-content;
			height: fit-content;
			margin-bottom: 0;
		}
		.techUsedSection {
			align-self: center;
			justify-self: start;
			height: -moz-fit-content;
			height: fit-content;
		}
	}
	aside {
		width: -moz-fit-content;
		width: fit-content;
	}
	.description__text {
		font-family: var(--slantText);
		color: var(--darkAquaLightHotPink);

		transition: 0.3s ease opacity;
	}
	@media (min-width: 1100px) {
		.description__text {
			font-family: var(--slantText);
			color: var(--darkAquaLightHotPink);

			transition: 0.3s ease opacity;
		}
		.description__text {
			font-family: var(--slantText);
			color: var(--darkAquaLightHotPink);

			transition: 0.3s ease opacity;
			margin: 0 0 1.05rem;
		}
		.projectDescription__text {
			margin-bottom: 1.15em;
		}
	}
	.fancyBoxWrapper {
		max-width: 250px;
		margin: 0 auto;
	}
	@media (min-width: 500px) {
		.fancyBoxWrapper {
			max-width: unset;
		}
	}
	.fancyBox {
		max-width: 250px;
	}
	.fancyBox h2 {
		font-size: var(--h3);
	}
	@media (min-width: 500px) {
		.fancyBox {
			position: relative;
			width: 280px;
			height: 380px;
			background: #fff;
			display: flex;
			justify-content: center;
			align-items: center;
			padding: 5px;
			max-width: unset;
		}
	}
	@media (min-width: 1200px) {
		.description__text {
			font-family: var(--slantText);
			color: var(--darkAquaLightHotPink);

			transition: 0.3s ease opacity;
			margin: 0 0 1.05rem;
		}
		.projectDescription__text {
			margin-bottom: 1.15rem;
		}
	}
	.fancyBoxWrapper {
		background: var(--cardBG);
		display: flex;
		justify-content: center;
		align-items: center;
		min-height: 400px;
	}
	.fancyBox {
		position: relative;
		width: 280px;
		height: 380px;
		background: #fff;
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 5px;
	}

	.fancyBox::before {
		content: '';
		position: absolute;
		top: -4px;
		left: -4px;
		right: -4px;
		bottom: -4px;
		background: linear-gradient(315deg, var(--aqua), var(--cardBG), var(--hotpink));
		transform: skewX(2deg) skewY(4deg);
	}

	.fancyBox::after {
		content: '';
		position: absolute;
		top: -4px;
		left: -4px;
		right: -4px;
		bottom: -4px;
		background: linear-gradient(315deg, var(--aqua), var(--cardBG), var(--hotpink));
		filter: blur(50px);
		animation: rotateIt 20s linear infinite;
		transform: skewX(2deg) skewY(4deg);
	}

	.fancyBox span {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background: var(--cardBGGradient);
		z-index: 2;
		transform: rotate(180deg);
	}
	.fancyBox a {
		text-decoration: none;
		color: inherit;
	}

	.fancyBox h2 {
		position: relative;
		z-index: 3;
		text-align: center;
		font-size: var(--h2);
	}
	@media (min-width: 700px) {
		.fancyBoxWrapper {
			min-height: 400px;
		}
		.fancyBox {
			width: 300px;
			height: 400px;
		}
	}
	@media (min-width: 1100px) {
		.fancyBoxWrapper {
			background: var(--cardBG);
			display: flex;
			justify-content: center;
			align-items: center;
			min-height: 300px;
		}
		.fancyBox {
			position: relative;
			width: 300px;
			height: 400px;
			background: #fff;
			display: flex;
			justify-content: center;
			align-items: center;
			padding: 5px;
		}
	}
	@media (min-width: 1200px) {
		.fancyBoxWrapper {
			background: var(--cardBG);
			display: flex;
			justify-content: center;
			align-items: center;
			min-height: 500px;
		}
		.fancyBox {
			position: relative;
			width: 300px;
			height: 400px;
			background: #fff;
			display: flex;
			justify-content: center;
			align-items: center;
			padding: 5px;
		}

		.fancyBox::before {
			content: '';
			position: absolute;
			top: -4px;
			left: -4px;
			right: -4px;
			bottom: -4px;
			background: linear-gradient(315deg, var(--aqua), var(--cardBG), var(--hotpink));
			transform: skewX(2deg) skewY(4deg);
		}

		.fancyBox::after {
			content: '';
			position: absolute;
			top: -4px;
			left: -4px;
			right: -4px;
			bottom: -4px;
			background: linear-gradient(315deg, var(--aqua), var(--cardBG), var(--hotpink));
			filter: blur(50px);
			animation: rotateIt 20s linear infinite;
			transform: skewX(2deg) skewY(4deg);
		}

		.fancyBox span {
			position: absolute;
			top: 0;
			left: 0;
			width: 100%;
			height: 100%;
			background: var(--cardBGGradient);
			z-index: 2;
			transform: rotate(180deg);
		}
		.fancyBox a {
			text-decoration: none;
			color: inherit;
		}

		.fancyBox h2 {
			position: relative;
			z-index: 3;
			text-align: center;
			font-size: var(--h3);
		}
	}

	@keyframes rotateIt {
		0% {
			transform: rotate(180deg);
		}
		20% {
			transform: rotate(0deg);
		}
		100% {
			transform: rotate(0deg);
		}
	}
	.stopRotation,
	.stopRotation::before,
	.stopRotation::after {
		animation: unset !important;
	}
	/* .hidden { */
	/* 	opacity: 0; */
	/* 	transition: 0.3s ease opacity; */
	/* } */
	.gridWrap {
		display: grid;
		grid-template-columns: minmax(1px, 1fr);
		grid-row-gap: 30px;
		place-items: center;
		margin: 0 auto 200px;
		width: min-content;
	}
	@media (min-width: 730px) {
		.gridWrap {
			grid-template-columns: repeat(3, minmax(1px, 1fr));
			width: 95%;
		}
	}
	@media (min-width: 985px) {
		.gridWrap {
			width: 85%;
		}
	}
	@media (min-width: 985px) {
		.gridWrap {
			width: 65%;
		}
	}
	@media (min-width: 1395px) {
		.gridWrap {
			width: 45%;
		}
	}
</style>
