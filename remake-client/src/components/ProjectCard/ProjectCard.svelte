<script>
	export let metaData;
	import Wave from '$components/Wave/Wave.svelte';
</script>

{#each metaData as md}
	<Wave>
		<div class="card">
			<h2>{md?.projectTitle}</h2>
			<li>
				<a href={md.slug}>{md.projectTitle}</a>
			</li>
		</div>
	</Wave>
{/each}
