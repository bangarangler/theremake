<script context="module">
	export async function load({ context }) {
		const url = '/projects/projectlist';
		const res = await fetch(url);

		if (res.ok) {
			const urlData = await res.json();
			return { context: { urlData } };
		}

		return {
			status: res.status,
			error: new Error(`Could not load ${url}`)
		};
	}
</script>

<!-- https://ar.al/2021/04/03/passing-data-from-layouts-to-pages-in-sveltekit/-->

<slot />
