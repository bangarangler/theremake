<script lang="ts">
	import { browser } from '$app/env';
	import { theme } from '$stores/themeStore';
	/* $: console.log('theme', $theme); */
	$: bool = $theme === 'DARK' ? true : false;
	/* $: console.log('$theme', $theme); */
	$: if ($theme) {
		if (browser) {
			const el = window.document.body;
			$theme === 'DARK' ? el.classList.remove('lightTheme') : el.classList.add('lightTheme');
		}
	}
</script>

<div class="container">
	<input
		type="checkbox"
		id="toggle"
		class="toggle--checkbox"
		bind:checked={bool}
		on:click={theme.toggle($theme)}
	/>
	<label for="toggle" class="toggle--label">
		<span class="toggle--label-background" />
	</label>
	<div class="background" />
</div>

<style>
	:root {
		/** sunny side **/
		--blue-background: #c2e9f6;
		/* --blue-border: #76cce3; */
		--blue-border: #33ffff;
		--blue-color: var(--aqua);
		--yellow-background: #fffaa8;
		--yellow-border: #f5eb71;
		/** dark side **/
		--indigo-background: #808fc7;
		--indigo-border: #5d6baa;
		--indigo-color: #6b7abb;
		--gray-border: #e8e8ea;
		--gray-dots: #e8e8ea;
		/** general **/
		--switchWhite: #fff;
	}
	.toggle--checkbox {
		display: none;
	}
	.container {
		display: grid;
		place-items: center;
		position: relative;
		margin-top: 0px;
		margin-bottom: 8px;
	}
	@media (min-width: 500px) {
		.container {
			margin-top: 8px;
			margin-bottom: unset;
		}
	}
	.toggle--label {
		width: 100px;
		height: 50px;
		background: var(--blue-color);
		border-radius: 50px;
		border: 2.5px solid var(--blue-border);
		display: flex;
		position: relative;
		transition: all 350ms ease-in;
	}
	.toggle--label:hover {
		cursor: pointer;
	}
	.toggle--label:before {
		animation-name: reverse;
		animation-duration: 350ms;
		animation-fill-mode: forwards;
		transition: all 350ms ease-in;
		content: '';
		width: 41px;
		height: 41px;
		border: 2.5px solid var(--yellow-border);
		top: 2px;
		left: 2px;
		position: absolute;
		border-radius: 41px;
		background: var(--yellow-background);
	}
	.background {
		position: absolute;
		left: 0;
		top: 0;
		background: var(--blue-background);
		border-radius: 50%;
		z-index: -1;
		width: 100%;
		height: 100%;
		transition: all 250ms ease-in;
	}
	.toggle--label-background {
		width: 5px;
		height: 2.5px;
		border-radius: 2.5px;
		position: relative;
		background: var(--white);
		left: 67.5px;
		top: 22.5px;
		transition: all 150ms ease-in;
	}
	.toggle--label-background:before {
		content: '';
		position: absolute;
		top: -2.5px;
		width: 20px;
		height: 2.5px;
		border-radius: 2.5px;
		background: var(--white);
		left: -10px;
		transition: all 150ms ease-in;
	}
	.toggle--label-background:after {
		content: '';
		position: absolute;
		top: 2.5px;
		width: 20px;
		height: 2.5px;
		border-radius: 2.5px;
		background: var(--white);
		left: -5px;
		transition: all 150ms ease-in;
	}
	.toggle--checkbox:checked ~ .background {
		background: var(--indigo-background);
	}
	.toggle--checkbox:checked + .toggle--label {
		background: var(--indigo-color);
		border-color: var(--indigo-border);
	}
	.toggle--checkbox:checked + .toggle--label:before {
		background: var(--white);
		border-color: var(--gray-border);
		animation-name: switch;
		animation-duration: 350ms;
		animation-fill-mode: forwards;
	}
	.toggle--label:after {
		transition-delay: 0ms;
		transition: all 250ms ease-in;
		position: absolute;
		content: '';
		box-shadow: var(--gray-dots) -7px 0 0 1px, var(--gray-dots) -12px 7px 0 -1px;
		left: 71.5px;
		top: 11.5px;
		width: 5px;
		height: 5px;
		background: transparent;
		border-radius: 25%;
		opacity: 0;
	}
	.toggle--checkbox:checked + .toggle--label:after {
		transition-delay: 350ms;
		opacity: 1;
	}
	.toggle--checkbox:checked + .toggle--label .toggle--label-background {
		left: 30px;
		width: 2.5px;
	}
	.toggle--checkbox:checked + .toggle--label .toggle--label-background:before {
		width: 2.5px;
		height: 2.5px;
		top: -12.5px;
	}
	.toggle--checkbox:checked + .toggle--label .toggle--label-background:after {
		width: 2.5px;
		height: 2.5px;
		left: -15px;
		top: 10px;
	}
	@keyframes reverse {
		0% {
			left: 52px;
			width: 41px;
		}
		60% {
			left: 36px;
			width: 56px;
		}
		100% {
			left: 2px;
		}
	}
	@keyframes switch {
		0% {
			left: 2px;
		}
		60% {
			left: 2px;
			width: 56px;
		}
		100% {
			left: 52px;
			width: 41px;
		}
	}
</style>
