<script lang="ts">
	let name: string;
	function changeName(s: string): string {
		return (name = s.toUpperCase());
	}
</script>

<h1>Blog...</h1>
<p>{changeName('jon p')}</p>
