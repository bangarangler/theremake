<script context="module">
	export async function load({ page }) {
		try {
			/* const Project = await import(`../../projects/${page.params.slug}.md`); */
			const Blog = await import(`../../blogposts/${page.params.slug}.md`);
			return {
				props: {
					Blog: Blog.default
				}
			};
		} catch (e) {
			console.log('err', e);
			return {
				status: 404,
				error: 'Blog not found'
			};
		}
	}
</script>

<script>
	export let Blog;
</script>

<Blog />
