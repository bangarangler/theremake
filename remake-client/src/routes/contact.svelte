<script>
	import Moon from '$components/Moon/Moon.svelte';
	import ContactForm from '$components/ContactForm/ContactForm.svelte';
	import ResumeButtons from '$components/ResumeButtons/ResumeButtons.svelte';
	import { onMount } from 'svelte';
	import { fly, fade } from 'svelte/transition';
	import { theme } from '$stores/themeStore';
	let visible = false;
	onMount(() => {
		visible = true;
	});
</script>

<section>
	<div class="wrapper">
		<h1 class="heading">Looking for Jon?</h1>
		{#if visible}
			<blockquote in:fade={{ delay: 100, duration: 1200 }} out:fade={{ duration: 100 }}>
				<p
					class="quote"
					in:fly={{ delay: 100, x: -100, opacity: 0, duration: 1200 }}
					out:fade={{ duration: 100 }}
				>
					This is the way... <br /> ... to contact me!
				</p>
				{#if $theme === 'DARK'}
					<div
						class="quoteFooter"
						in:fly={{ delay: 100, x: 100, opacity: 0, duration: 1000 }}
						out:fade={{ duration: 100 }}
					>
						<p>- I have spoken ; )</p>
					</div>
				{:else}
					<div
						class="quoteFooter light"
						in:fly={{ delay: 100, x: 100, opacity: 0, duration: 1000 }}
						out:fade={{ duration: 100 }}
					>
						<p>- I have spoken ; )</p>
					</div>
				{/if}
			</blockquote>
		{/if}
		<div class="moon">
			<Moon />
		</div>
	</div>
	<div class="formWrapper">
		<ContactForm />
	</div>
	<ResumeButtons />
</section>

<style>
	.wrapper {
		display: grid;
		/* margin-bottom: 50px; */
		/*grid-template-columns: minmax(0, auto);*/
		grid-template-rows: minmax(0, auto) 1fr;
		place-items: center;
		grid-template-areas:
			'head'
			'moon'
			'wrap';
	}
	/* @media (min-width: 500px) { */
	/* 	.wrapper { */
	/* 		margin-bottom: 100px; */
	/* 	} */
	/* } */
	@media (min-width: 1000px) {
		.wrapper {
			width: 90%;
			margin: 0 auto;
			grid-column-gap: 30px;
			/* grid-template-columns: repeat(2, minmax(0, 1fr)); */
			grid-template-columns: repeat(2, minmax(0, auto));
			grid-template-rows: minmax(0, auto) 1fr;
			grid-template-areas:
				'head head'
				'wrap moon';
		}
	}
	@media (min-width: 1400px) {
		.wrapper {
			max-width: 1400px;
		}
	}
	blockquote {
		grid-area: wrap;
		text-align: center;
		border-radius: 12px;
		padding-bottom: 12px;
		margin-bottom: 35px;
	}
	@media (min-width: 500px) {
		blockquote p {
			font-size: 1.5rem;
			padding-bottom: 0.2em;
		}
	}

	.quote {
		quotes: '\201C''\201D''\2018''\2019';
	}

	.quote:before {
		content: open-quote;
		display: inline;
		height: 0;
		line-height: 0;
		left: -10px;
		position: relative;
		top: 20px;
		color: var(--hotpink);
		font-size: 3em;
	}

	.quote::after {
		content: close-quote;
		display: inline;
		height: 0;
		line-height: 0;
		left: 10px;
		position: relative;
		top: 25px;
		color: var(--aqua);
		font-size: 3em;
	}

	.quoteFooter {
		margin: 0;
		text-align: right;
		font-style: italic;
		font-family: var(--slantText);
		padding: 0;
		text-align: center;
		background: linear-gradient(to bottom left, #7a00cc, #ff1a75);
		background-clip: border-box;
		-webkit-background-clip: text;
		-webkit-text-fill-color: transparent;
		font-size: 1.3rem;
		letter-spacing: 1.5;
	}
	@media (min-width: 1200px) {
		.quote {
			quotes: '\201C''\201D''\2018''\2019';
		}
	}
	.quoteFooter p {
		letter-spacing: 1.5px;
	}

	.light {
		-webkit-text-fill-color: unset;
	}
	.heading {
		grid-area: head;
		font-size: var(--h2);
		margin-top: 35px;
		background: linear-gradient(271deg, var(--hotpink) 30%, var(--aqua) 60%, #a162e8 74%);
		background-clip: border-box;
		-webkit-background-clip: text;
		-webkit-text-fill-color: transparent;
		text-align: center;
	}
	@media (min-width: 500px) {
		.heading {
			font-size: var(--h1);
			margin-top: unset;
		}
	}
	.moon {
		grid-area: moon;
		max-width: 300px;
	}
	@media (min-width: 500px) {
		.moon {
			max-width: 400px;
		}
	}
	@media (min-width: 1000px) {
		.moon {
			max-width: 500px;
		}
	}
	@media (min-width: 1200px) {
		.moon {
			max-width: 600px;
		}
	}
	.formWrapper {
		width: 100%;
		height: 100%;
		margin: 0px auto;
	}
	@media (min-width: 910px) {
		.formWrapper {
			margin: 100px auto 100px;
		}
	}
	@media (min-width: 1200px) {
		.formWrapper {
			/* height: 80vh; */
			height: unset;
		}
	}
	@media (min-width: 1400px) {
		.formWrapper {
			max-width: 1400px;
		}
	}
</style>
