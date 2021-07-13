<script lang="ts">
	export let projectRole: string;
	export let projectTitle: string;
	export let slug: string;
</script>

<div class="container">
	<div class="card headingWrapper">
		<a href={`/projects/${slug}`}
			><h4>{projectRole}</h4>
			<p>{projectTitle}</p>
		</a>
	</div>
</div>

<style>
	.container {
		position: relative;
	}
	.headingWrapper {
		/* border: 1px solid hotpink; */
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: flex-end;
		width: -moz-fit-content;
		width: fit-content;
		/* position: absolute; */
		/* top: 0; */
		/* right: 0; */
	}
	.headingWrapper a {
		text-decoration: none;
		color: inherit;
	}
	h4 {
		margin-top: 0;
	}
	p {
		text-align: right;
		margin-bottom: 0;
		font-size: var(--smallTextSize);
	}
</style>
