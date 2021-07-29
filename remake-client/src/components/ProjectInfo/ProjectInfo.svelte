<script>
	export let projectRole;
	export let techUsed;
	export let organization;
	export let year;
</script>

<div class="card projectInfoContaner">
	<div class="projSection">
		<h6>Position</h6>
		<p>{projectRole}</p>
		<h6>Organization</h6>
		<p>{organization}</p>
		<h6>Year</h6>
		<p>{year}</p>
	</div>
	<div class="projSection">
		<h6>Work</h6>
		<ul>
			{#each techUsed as tu}
				<li>{tu}</li>
			{/each}
		</ul>
	</div>
</div>

<style>
	.projectInfoContaner {
		display: grid;
		grid-template-columns: minmax(0, 1fr);
		width: 95%;
		margin: 65px auto 0 auto;
		place-items: center;
	}
	.projectInfoContaner .projSection {
		display: flex;
		flex-direction: column;
		width: 75%;
		align-items: flex-start;
		justify-content: center;
	}
	@media (min-width: 400px) {
		.projectInfoContaner {
			width: 80%;
		}
	}
	@media (min-width: 450px) {
		.projectInfoContaner {
			width: 80%;
		}
	}
	@media (min-width: 650px) {
		.projectInfoContaner {
			width: 80%;
			max-width: 600px;
			grid-template-columns: minmax(0, 1fr) minmax(0, 1fr);
			grid-column-gap: 0px;
		}
		.projectInfoContaner .projSection {
			width: unset;
		}
	}
	@media (min-width: 1000px) {
		.projectInfoContaner {
			grid-template-columns: minmax(0, 1fr) minmax(0, 1fr);
			width: 60%;
			grid-column-gap: 0px;
		}
	}
	@media (min-width: 1200px) {
		.projectInfoContaner {
			grid-template-columns: minmax(0, 1fr) minmax(0, 1fr);
			width: 50%;
			grid-column-gap: 30px;
		}
	}
	.projectInfoContaner H6 {
		color: var(--projH6DarkLight);
		border-bottom: 1px solid var(--hotpink);
		width: -moz-fit-content;
		width: fit-content;
	}
</style>
