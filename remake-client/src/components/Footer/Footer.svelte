<script lang="ts">
	import { theme } from '$stores/themeStore';
	import FaInstagram from 'svelte-icons/fa/FaInstagram.svelte';
	import FaTwitter from 'svelte-icons/fa/FaTwitter.svelte';
	import FaFacebook from 'svelte-icons/fa/FaFacebook.svelte';
	import FaLinkedinIn from 'svelte-icons/fa/FaLinkedinIn.svelte';
	import FaSlack from 'svelte-icons/fa/FaSlack.svelte';
	import FaGithub from 'svelte-icons/fa/FaGithub.svelte';

	interface FooterSocialType {
		id: number;
		icon: any;
		path: string;
		text: string;
	}
	const footerSocial: FooterSocialType[] = [
		{
			id: 0,
			icon: FaInstagram,
			path: 'https://www.instagram.com/bangarangler/?hl=en',
			text: 'instagram'
		},
		{
			id: 1,
			icon: FaTwitter,
			path: 'https://twitter.com/bangarangler',
			text: 'twitter'
		},
		{
			id: 2,
			icon: FaFacebook,
			path: 'https://www.facebook.com/jonathan.palacio.9461?ref=bookmarks',
			text: 'facebook'
		},
		{
			id: 3,
			icon: FaLinkedinIn,
			/* icon: <FaLinkedinIn className={styles.linkedin} />, */
			path: 'https://www.linkedin.com/in/jonathan-palacio-401945a3/',
			text: 'linkedin'
		},
		{
			id: 4,
			icon: FaGithub,
			path: 'https://github.com/bangarangler',
			text: 'github'
		},
		{
			id: 5,
			icon: FaSlack,
			path:
				'https://join.slack.com/t/bangarangler/shared_invite/enQtNjA4OTc5NTQwMTk1LTNjOTRlYjIwOTkyYzBhYTNmYjBjOWY5MTFjNGRkNmUyZWQ3MThiMDI1ZWQ3NjM2Mzg2NDNmZjllMGI3ZmE5NmE',
			text: 'slack'
		}
	];
</script>

<footer>
	{#if $theme === 'DARK'}
		<h3>Jonathan Dain Palacio</h3>
	{:else}
		<h3 class="light">Jonathan Dain Palacio</h3>
	{/if}
	<div class="footerSocialWrapper">
		{#each footerSocial as fs}
			<a href={fs.path} class={fs.text} target="_blank"><svelte:component this={fs.icon} /></a>
		{/each}
	</div>
	<div class="copy">
		<p>&copy; 2017 - 2021 Jonathan Dain Palacio</p>
	</div>
</footer>

<style>
	footer {
		display: grid;
		grid-template-rows: minmax(0, 1fr) minmax(0, 1fr) minmax(0, 0.5fr);
		grid-template-areas:
			'name'
			'social'
			'copy';
		place-items: center;
	}
	footer h3 {
		grid-area: 'name';
		padding: 1.2rem;
		font-family: var(--slantText);
		border: 6px solid;
		background: -webkit-gradient(linear, right top, left bottom, from(#7a00cc), to(#ff1a75));
		background: linear-gradient(to bottom left, #7a00cc, #ff1a75);
		background-clip: border-box;
		-webkit-background-clip: text;
		-webkit-text-fill-color: transparent;
		-webkit-border-image: -webkit-gradient(
				linear,
				right top,
				left bottom,
				from(#7a00cc),
				to(#ff1a75)
			)
			1;
		-webkit-border-image: linear-gradient(to bottom left, #7a00cc, #ff1a75) 1;
		border-image: -webkit-gradient(linear, right top, left bottom, from(#7a00cc), to(#ff1a75)) 1;
		border-image: linear-gradient(to bottom left, #7a00cc, #ff1a75) 1;
		text-align: center;
		width: max-content;
		margin: 20px auto;
		font-size: var(--h4);
	}
	@media (min-width: 800px) {
		footer h3 {
			padding: 1.5rem;
			font-size: var(--h3);
		}
	}
	@media (min-width: 1200px) {
		footer h3 {
			padding: 1.5rem;
			font-size: var(--h3);
		}
	}
	footer .light {
		padding: 1.5rem;
		font-family: var(--slantText);
		border: 6px solid;
		background: -webkit-gradient(linear, right top, left bottom, from(var(--aqua)), to(#7a00cc));
		background: linear-gradient(to bottom left, var(--aqua), #7a00cc);
		background-clip: border-box;
		-webkit-background-clip: text;
		-webkit-text-fill-color: transparent;
		-webkit-border-image: -webkit-gradient(
				linear,
				right top,
				left bottom,
				from(#7a00cc),
				to(var(--aqua))
			)
			1;
		-webkit-border-image: linear-gradient(to bottom left, #7a00cc, var(--aqua)) 1;
		border-image: -webkit-gradient(linear, right top, left bottom, from(#7a00cc), to(var(-aqua))) 1;
		border-image: linear-gradient(to bottom left, #7a00cc, var(--aqua)) 1;
		text-align: center;
		width: max-content;
		margin: 20px auto;
	}
	footer .footerSocialWrapper {
		--iconSize: 40px;
		grid-area: 'social';
		width: 90%;
		display: flex;
		align-items: center;
		justify-content: space-between;
	}
	@media (min-width: 800px) {
		footer .footerSocialWrapper {
			--iconSize: 50px;
			width: 95%;
		}
	}
	@media (min-width: 1200px) {
		footer .footerSocialWrapper {
			--iconSize: 60px;
			width: 100%;
		}
	}
	.instagram {
		color: #5851db;
		width: var(--iconSize);
	}
	footer a:hover {
		color: var(--hotpink);
		transition: all 0.5s linear;
		transform: skew(-5deg);
	}
	.facebook {
		color: #3b5998;
		width: var(--iconSize);
	}
	.twitter {
		color: #1da1f2;
		width: var(--iconSize);
	}
	.linkedin {
		color: #0077b5;
		width: var(--iconSize);
	}
	.slack {
		color: #cf0e5b;
		width: var(--iconSize);
	}
	.github {
		color: #6cc644;
		width: var(--iconSize);
	}
	.copy {
		grid-area: 'copy';
	}
	.copy p {
		color: var(--hotpink);
		font-size: var(--h6);
	}
	@media (min-width: 800px) {
		.copy p {
			font-size: var(--h5);
		}
	}
	@media (min-width: 1200px) {
		.copy p {
			font-size: var(--h4);
		}
	}
</style>
