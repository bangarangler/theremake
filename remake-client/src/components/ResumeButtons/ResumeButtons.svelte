<script>
	import { browser } from '$app/env';
</script>

{#if browser}
	<div class="resWrap">
		<div class="resumeDownload">
			<p class="prettyResText">Pretty Resume</p>
			<button class="glow-on-hover" type="button"
				><a href="/software-engineer-jon-palacio-pretty.pdf" rel="external" target="_blank"
					>Download Pretty</a
				></button
			>
		</div>
		<div class="resumeDownload">
			<p class="proResText">Pro Resume</p>
			<button class="glow-on-hover" type="button"
				><a href="/software-engineer-jon-palacio-pro.pdf" rel="external" target="_blank"
					>Download Pro!</a
				></button
			>
		</div>
	</div>
{/if}

<style>
	.resWrap {
		display: grid;
		grid-template-columns: minmax(1px, 1fr);
		width: 95%;
		margin: 0 auto 100px;
	}
	@media (min-width: 580px) {
		.resWrap {
			width: 80%;
			grid-template-columns: repeat(2, minmax(1px, 1fr));
		}
	}
	@media (min-width: 950px) {
		.resWrap {
			width: 50%;
		}
	}
	@media (min-width: 1200px) {
		.resWrap {
			max-width: 800px;
		}
	}
	.resumeDownload {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		margin-bottom: 35px;
	}
	@media (min-width: 580px) {
		.resumeDownload {
			margin-bottom: unset;
		}
	}
	.resumeDownload p {
		font-size: var(--h4);
	}
	.resumeDownload .prettyResText {
		color: var(--darkAquaLightHotPink);
		font-family: var(--slantText);
		letter-spacing: 2px;
	}
	.resumeDownload .proResText {
		color: var(--dracPurp);
		font-family: var(--slantText);
		letter-spacing: 2px;
	}
	.glow-on-hover a {
		text-decoration: none;
		color: inherit;
	}
	.glow-on-hover {
		width: 220px;
		height: 50px;
		border: none;
		outline: none;
		color: var(--white);
		background: var(--darkGray);
		cursor: pointer;
		position: relative;
		z-index: 0;
		border-radius: 10px;
	}

	.glow-on-hover:before {
		content: '';
		background: linear-gradient(45deg, #0fc, #bd93f9, #ff1a75, #00ffd5, #002bff);
		position: absolute;
		top: -2px;
		left: -2px;
		background-size: 400%;
		z-index: -1;
		filter: blur(5px);
		width: calc(100% + 4px);
		height: calc(100% + 4px);
		animation: glowing 20s linear infinite;
		opacity: 0;
		transition: opacity 0.3s ease-in-out;
		border-radius: 10px;
	}

	.glow-on-hover:active {
		color: #000;
	}

	.glow-on-hover:active:after {
		background: transparent;
	}

	.glow-on-hover:hover:before {
		opacity: 1;
	}

	.glow-on-hover:after {
		z-index: -1;
		content: '';
		position: absolute;
		width: 100%;
		height: 100%;
		background: var(--darkGray);
		left: 0;
		top: 0;
		border-radius: 10px;
	}

	@keyframes glowing {
		0% {
			background-position: 0 0;
		}
		50% {
			background-position: 400% 0;
		}
		100% {
			background-position: 0 0;
		}
	}
</style>
