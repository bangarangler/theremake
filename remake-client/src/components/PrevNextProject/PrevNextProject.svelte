<script>
	import IoIosArrowRoundBack from 'svelte-icons/io/IoIosArrowRoundBack.svelte';
	import IoIosArrowRoundForward from 'svelte-icons/io/IoIosArrowRoundForward.svelte';
</script>

<div class="prevNextContainer">
	<div class="arrowContainer">
		<span class="arrowIcons"><IoIosArrowRoundBack /></span>
		<span class="arrowIcons"><IoIosArrowRoundForward /></span>
	</div>
</div>

<style>
	.prevNextContainer {
		width: 800px;
		height: 400px;
		min-width: 800px;
		min-height: 400px;
		border: 3px solid var(--darkGray);
		border-radius: 12px;
		margin: 100px auto 0 auto;
	}
	.arrowContainer {
		display: flex;
		justify-content: space-between;
		align-items: center;
		width: 100%;
		height: 100%;
	}
	.arrowIcons {
		display: block;
		width: 100px;
		height: 100px;
	}
</style>
