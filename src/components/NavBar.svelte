<script lang="ts">
	import logo from '$images/beforeDawnTempLogo.svg';
	import ToggleSwitch from '$components/ThemeToggle.svelte';
</script>

<header>
	<nav>
		<a href="/">
			<img src={logo} alt="jonathandain.dev logo" />
		</a>
		<a href="/">Home</a>
		<a href="/blog">Blog</a>
		<a href="/uses">Uses</a>
		<a href="/contact">Contact</a>
	</nav>
	<ToggleSwitch />
</header>

<style>
	nav {
		display: flex;
		align-items: center;
		margin-top: 12px;
	}
	img {
		max-height: 80px;
		max-width: 80px;
	}
</style>
