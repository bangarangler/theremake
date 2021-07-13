<script>
	export let metaData;
	import Wave from '$components/Wave/Wave.svelte';
	import ProjectCardHeading from '$components/ProjectCard/ProjectCardHeading/ProjectCardHeading.svelte';
	/* const list = metaData && metaData?.map((md) => md?.techUsed); */
	/* console.log('list', list); */
</script>

{#each metaData as md}
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
					<p class="projectDescription__text">{techUsed}</p>
				{/each}
			</aside>
			<!--<li>-->
			<!--<a href={`/projects/${md.slug}`}>{md.projectTitle}</a>-->
			<!--</li>-->
			<div class="fancyBoxWrapper cardSection">
				<div class="fancyBox">
					<span />
					<h2>{md?.projectTitle}</h2>
				</div>
			</div>
		</div>
	</Wave>
{/each}

<style>
	.wrapper {
		/* border: 1px solid hotpink; */
		display: grid;
		grid-row-gap: 65px;
		grid-template-columns: minmax(0px, 1fr) minmax(0px, 1fr);
		grid-template-rows: minmax(0px, 1fr) minmax(0px, 1fr) minmax(0px, 1fr);
		grid-template-areas:
			'description description title'
			'techUsed card .'
			'techUsed card .';
	}
	.descriptionSection {
		grid-area: description;
		height: -moz-fit-content;
		height: fit-content;
	}
	.techUsedSection {
		grid-area: techUsed;
		height: -moz-fit-content;
		height: fit-content;
	}
	.titleSection {
		grid-area: title;
	}
	.cardSection {
		grid-area: card;
	}
	aside {
		/* border: 1px solid var(--lightGray); */
		width: -moz-fit-content;
		width: fit-content;
	}
	.description__text {
		font-family: var(--slantText);
		color: var(--descriptionColor);
	}
	/* .projectDescription__text { */
	/* font-size: var(--smallTextSize); */
	/* } */
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
		/* animation: rotateIt 10s linear infinite; */
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
		/* background: linear-gradient(315deg, var(--black), var(--cardBG), var(--black)); */
		background: var(--cardBGGradient);
		z-index: 2;
		transform: rotate(180deg);
	}

	.fancyBox h2 {
		position: relative;
		z-index: 3;
		text-align: center;
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
</style>
