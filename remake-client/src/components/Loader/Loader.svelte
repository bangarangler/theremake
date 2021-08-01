<script>
	console.log('loading');
</script>

<div class="loaderWrapper">
	<div class="loader" />
</div>

<style>
	.loaderWrapper {
		max-height: 200px;
		/* max-width: 300px; */
		width: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
		background: transparent;
	}
	.loader {
		position: relative;
		width: 150px;
		height: 150px;
		border-radius: 50%;
		background: linear-gradient(45deg, transparent, transparent, 40%, #e5f403);
		animation: animate 2s linear infinite;
	}
	@keyframes animate {
		0% {
			transform: rotate(0deg);
			filter: hue-rotate(0deg);
		}
		100% {
			transform: rotate(360deg);
			filter: hue-rotate(360deg);
		}
	}
	.loader:before {
		content: '';
		position: absolute;
		top: 6px;
		bottom: 6px;
		right: 6px;
		left: 6px;
		background: var(--black);
		border-radius: 50%;
		z-index: 1000;
	}
	.loader:after {
		content: '';
		position: absolute;
		top: 0px;
		bottom: 0px;
		right: 0px;
		left: 0px;
		background: linear-gradient(45deg, transparent, transparent, 40%, #e5f403);
		border-radius: 50%;
		z-index: 1000;
		z-index: 1;
		filter: blur(30px);
	}
</style>
