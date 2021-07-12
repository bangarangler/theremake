<script context="module">
	export async function load({ page }) {
		try {
			const Project = await import(`../../projects/${page.params.slug}.md`);
			return {
				props: {
					Project: Project.default
				}
			};
		} catch (e) {
			console.log('err', e);
			return {
				status: 404,
				error: 'Project not found'
			};
		}
	}
</script>

<script>
	export let Project;
</script>

<Project />
