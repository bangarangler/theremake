<script>
	import { projectMetaData } from '$stores/projectMetaData';
	import Wave from '$components/Wave/Wave.svelte';
	import ProjectCardHeading from '$components/ProjectCard/ProjectCardHeading/ProjectCardHeading.svelte';
</script>

{#each $projectMetaData as md}
	<Wave>
		<div class="card wrapper">
			<div class="titleSection">
				<ProjectCardHeading
					projectTitle={md.projectTitle}
					projectRole={md.projectRole}
					slug={md.slug}
				/>
			</div>
			<aside class="card descriptionSection">
				<h5 class="description__text">Description:</h5>
				<p class="projectDescription__text">
					{md.projectDescription}
				</p>
			</aside>
			<aside class="card techUsedSection">
				<h5 class="description__text">Tech Used:</h5>
				{#each md.techUsed as techUsed}
					<p class="projectDescription__text">
						{techUsed}
					</p>
				{/each}
			</aside>
			<div class="fancyBoxWrapper cardSection">
				<div class="fancyBox">
					<span />
					<h2><a href={`/projects/${md?.slug}`}>{md?.projectTitle}</a></h2>
				</div>
			</div>
		</div>
	</Wave>
{/each}

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
		color: var(--descriptionColor);

		transition: 0.3s ease opacity;
	}
	@media (min-width: 1100px) {
		.description__text {
			font-family: var(--slantText);
			color: var(--descriptionColor);

			transition: 0.3s ease opacity;
		}
		.description__text {
			font-family: var(--slantText);
			color: var(--descriptionColor);

			transition: 0.3s ease opacity;
			margin: 0 0 1.05rem;
		}
		.projectDescription__text {
			margin-bottom: 1.15em;
		}
	}
	@media (min-width: 1200px) {
		.description__text {
			font-family: var(--slantText);
			color: var(--descriptionColor);

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
	/* .hidden { */
	/* 	opacity: 0; */
	/* 	transition: 0.3s ease opacity; */
	/* } */
</style>
