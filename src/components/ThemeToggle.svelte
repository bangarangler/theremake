<script lang="ts">
	import { theme } from '$stores/themeStore';
	$: isServer = typeof window === undefined ? true : false;
	$: bool = $theme === 'DARK' ? true : false;
	$: console.log('bool', bool);
	$: {
		doThis();
	}
	function doThis() {
		console.log('running do this...');
		if (isServer) {
			if (bool) {
				/* let test = window.document.body; */
				let test = window.document.getElementsByTagName('HTML');
				console.log('running if...');
				console.log('test', test);
				window.document.body.classList.add('lightTheme');
				/* test.classList.add('lightTheme'); */
			} else {
				let test = window.document.body;
				console.log('running else...');
				console.log('test', test);
				test.classList.remove('lightTheme');
				/* window.document.body.classList.add('darkTheme'); */
			}
		}
	}
</script>

<div class="container">
	<input
		type="checkbox"
		id="toggle"
		class="toggle--checkbox"
		bind:checked={bool}
		on:click={!$theme ? theme.lightMode : theme.darkMode}
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
		margin-top: 8px;
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
