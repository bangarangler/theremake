<script>
	export let metaData;
	import Wave from '$components/Wave/Wave.svelte';
</script>

{#each metaData as md}
	<Wave>
		<div class="card">
			<li>
				<a href={`/projects/${md.slug}`}>{md.projectTitle}</a>
			</li>
			<div class="fancyBoxWrapper">
				<div class="fancyBox">
					<span />
					<h2>{md?.projectTitle}</h2>
				</div>
			</div>
		</div>
	</Wave>
{/each}

<style>
	.fancyBoxWrapper {
		background: var(--cardBG);
		display: flex;
		justify-content: center;
		align-items: center;
		min-height: 600px;
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
		background: var(--cardBG);
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
