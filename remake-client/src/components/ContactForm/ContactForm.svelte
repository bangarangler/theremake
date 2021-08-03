<script lang="ts">
	import contactImg from '$static/contact-image.svg?w=600;700;1000&format=svg&srcset';
	/* import Portal from '$components/Portal/Portal.svelte'; */
	import Loader from '$components/Loader/Loader.svelte';
	import ImageLoader from '$images/ImageLoader.svelte';
	import Modal from '$components/Modal/Modal.svelte';
	import Toast from '$components/Toast/Toast.svelte';
	import { isValid } from '$stores/toastStore';
	import { validate, ValidOptions } from './validate';
	import { BASE_URL } from '$lib/Env';
	export let isModalOpen = false;
	let name: string = '';
	let email: string = '';
	let message: string = '';
	let subject: string = '';
	let loading = false;
	let displayMsg = '';
	let displayEnding = '';
	async function submit() {
		displayMsg = '';
		/* console.log({ name, email, subject, message }); */
		$isValid = true;
		const ok = validate({ name, email, subject, message });
		if (ok === ValidOptions.ALL_GOOD) {
			loading = true;
			try {
				const res = await fetch(`http://${BASE_URL}/api/contact`, {
					/* const res = await fetch(`https://${BASE_URL}/api/contact`, { */
					method: 'post',
					headers: {
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({ name, email, subject, message })
				});
				if (res) {
					/* console.log('res', await res.json()); */
					displayMsg = `Hey ${name}; 🐚\n I got your message!`;
					displayEnding = '🎈 - Mischief Managed';
					resetInputs();
					isModalOpen = true;
				}
				$isValid = false;
				loading = false;
			} catch (err) {
				console.log('err fetching', err);
				displayMsg =
					'⛔ Oh my... Something\n went wrong! Please try\n again or contact me on slack!';
				displayEnding = '🎈 - Mischief NOT Managed';
				isModalOpen = true;
				$isValid = false;
				loading = false;
			}
		}
	}
	function resetInputs() {
		email = '';
		message = '';
		subject = '';
		name = '';
	}
</script>

<div class="wrapper">
	{#if !loading}
		<form on:submit|preventDefault={() => submit()}>
			<Toast duration={3000} />
			<Modal bind:isModalOpen>
				<div class="displayMsgWrapper">
					<pre>{displayMsg}</pre>
					<p class="mischief">{displayEnding}</p>
				</div>
			</Modal>
			<label for="name">Name</label>
			<input
				id="name"
				type="text"
				placeholder="Jane Doe"
				bind:value={name}
				on:blur={() => ($isValid = false)}
			/>
			<label for="email">Email</label>
			<input
				id="email"
				required
				type="email"
				placeholder="jane.doe@abracadabra.com"
				bind:value={email}
			/>
			<label for="subject">Subject</label>
			<input id="subject" type="text" placeholder="Mercury retrograde..." bind:value={subject} />
			<label for="message">Message</label>
			<textarea
				name="message"
				id="message"
				cols="30"
				rows="50"
				placeholder="My deep dark secret is..."
				bind:value={message}
			/>
			<div class="btnWrapper">
				<button class="glow-on-hover" type="submit" disabled={$isValid}>Send to Jon</button>
				<button class="glow-on-hover cancel" type="button" on:click={resetInputs}>Cancel</button>
			</div>
			<div>
				<h3>Let's Build Something!</h3>
			</div>
		</form>
	{/if}
	{#if loading}
		<form>
			<Loader />
			<div>
				<h3>Let's Build Something!</h3>
			</div>
			<div class="btnWrapper">
				<button class="glow-on-hover" type="submit" disabled={$isValid}
					>{loading ? 'Sending...' : 'Send to Jon'}</button
				>
				<button class="glow-on-hover cancel" type="button">Cancel</button>
			</div>
		</form>
	{/if}
	<div class="stuff">
		<div class="imgWrapper">
			<h2>Excited to chat!</h2>
			<picture>
				<source media="(min-width:1000px)" srcset={contactImg} />
				<ImageLoader
					ty="image/svg"
					srcset={contactImg}
					alt="Epic illistration from samji_illustrator!  Please check them out!"
				/>
			</picture>
			<p>reach out anytime!</p>
		</div>
	</div>
</div>

<style>
	.wrapper {
		display: grid;
		grid-template-columns: minmax(0, 1fr);
		width: 80%;
		height: 100%;
		margin: 0 auto;
		grid-row-gap: 50px;
		grid-template-areas:
			'stuff'
			'form';
	}
	@media (min-width: 1100px) {
		.wrapper {
			grid-template-columns: repeat(2, minmax(0, 1fr));
			margin: unset;
			width: 100%;
			grid-gap: unset;
			grid-template-areas: 'stuff form';
		}
	}
	.wrapper form {
		grid-area: form;
		display: flex;
		flex-direction: column;
		align-items: baseline;
		width: 90%;
		margin: 0 auto;
	}
	@media (min-width: 500px) {
		.wrapper form {
			width: 75%;
		}
	}
	@media (min-width: 1100px) {
		.wrapper form {
			align-items: unset;
		}
	}
	.wrapper form input::placeholder {
		color: var(--black);
	}
	.wrapper form textarea::placeholder {
		color: var(--black);
	}
	.wrapper form button {
		width: max-content;
		margin-bottom: 15px;
	}
	@media (min-width: 1100px) {
		.wrapper form button {
			margin-bottom: 10px;
		}
	}
	.wrapper form div h3 {
		font-family: var(--slantText);
		transform: skewY(-15deg);
		border-bottom: 4px solid var(--darkAquaLightHotPink);
		margin: 100px auto 120px;
		font-size: var(--h5);
	}
	@media (min-width: 400px) {
		.wrapper form div {
			font-size: var(--h4);
		}
	}
	@media (min-width: 500px) {
		.wrapper form div {
			font-size: var(--h4);
		}
	}
	@media (min-width: 700px) {
		.wrapper form div h3 {
			font-size: var(--h4);
			margin: 85px auto 120px;
		}
	}
	@media (min-width: 1200px) {
		.wrapper form div h3 {
			font-size: var(--h4);
			margin: 100px auto 120px;
		}
	}
	/*TODO: check on large screen*/
	@media (min-width: 1600px) {
		.wrapper form div h3 {
			font-size: var(--h4);
		}
	}
	.wrapper form div {
		align-self: center;
	}
	@media (min-width: 1100px) {
		.wrapper form div {
			align-self: unset;
		}
	}
	.wrapper .stuff {
		grid-area: stuff;
	}
	.imgWrapper {
		/* height: 100%; */
		/* width: 100%; */
		border-radius: 12px;
		background: var(--cardBG);
		min-height: auto;
		padding: 2em;
		display: grid;
		grid-template-rows: minmax(0, auto) minmax(0, 1fr) minmax(0, auto);
		grid-template-areas:
			'top'
			'middle'
			'bottom';
	}
	/* @media (min-width: 1100px) { */
	/* 	.imgWrapper { */
	/* 		border-radius: unset; */
	/* 	} */
	/* } */
	@media (min-width: 500px) {
		.imgWrapper {
			min-height: 60vh;
		}
	}
	@media (min-width: 1200px) {
		.imgWrapper {
			min-height: 80vh;
		}
	}
	.imgWrapper h2 {
		grid-area: top;
		text-align: center;
		margin: 0;
		color: var(--darkAquaLightHotPink);
		font-size: var(--h4);
	}
	@media (min-width: 500px) {
		.imgWrapper h2 {
			font-size: var(--h3);
		}
	}
	@media (min-width: 1200px) {
		.imgWrapper h2 {
			font-size: var(--h2);
		}
	}
	.imgWrapper picture {
		grid-area: middle;
		max-width: 400px;
		margin: 0 auto;
	}
	@media (min-width: 1200px) {
		.imgWrapper picture {
			max-width: 600px;
		}
	}
	.imgWrapper p {
		grid-area: bottom;
		text-align: center;
		margin: 0 auto;
		font-family: var(--slantText);
		letter-spacing: 1.5px;
		color: var(--darkAquaLightHotPink);
	}
	.btnWrapper {
		display: flex;
		flex-direction: column;
		width: 80%;
		align-items: center;
		justify-content: space-evenly;
		margin-top: 30px;
	}
	@media (min-width: 1100px) {
		.btnWrapper {
			flex-direction: row;
			width: 100%;
			margin-top: 50px;
		}
	}
	@media (min-width: 1360px) {
		.btnWrapper {
			margin-top: 100px;
		}
	}
	.cancel {
		height: 30px;
	}
	.glow-on-hover {
		width: 220px;
		height: 50px;
		border: none;
		outline: none;
		color: var(--white);
		background: var(--darkGray);
		cursor: pointer;
		position: relative;
		z-index: 0;
		border-radius: 10px;
	}

	.glow-on-hover:before {
		content: '';
		background: linear-gradient(45deg, #0fc, #bd93f9, #ff1a75, #00ffd5, #002bff);
		position: absolute;
		top: -2px;
		left: -2px;
		background-size: 400%;
		z-index: -1;
		filter: blur(5px);
		width: calc(100% + 4px);
		height: calc(100% + 4px);
		animation: glowing 20s linear infinite;
		opacity: 0;
		transition: opacity 0.3s ease-in-out;
		border-radius: 10px;
	}

	.glow-on-hover:active {
		color: #000;
	}

	.glow-on-hover:active:after {
		background: transparent;
	}

	.glow-on-hover:hover:before {
		opacity: 1;
	}

	.glow-on-hover:after {
		z-index: -1;
		content: '';
		position: absolute;
		width: 100%;
		height: 100%;
		background: var(--darkGray);
		left: 0;
		top: 0;
		border-radius: 10px;
	}

	@keyframes glowing {
		0% {
			background-position: 0 0;
		}
		50% {
			background-position: 400% 0;
		}
		100% {
			background-position: 0 0;
		}
	}
	.cancel {
		height: 35px;
	}
	.cancel:hover:before {
		background: linear-gradient(45deg, var(--red), var(--hotpink));
	}
	.displayMsgWrapper {
		/* border: 1px solid red; */
		display: grid;
		place-items: center;
	}
	pre {
		color: var(--darkAquaLightHotPink);
		font-weight: bold;
		border-bottom: 2px solid var(--dracPurp);
	}
	.mischief {
		font-family: var(--slantText);
		letter-spacing: 1.5px;
		font-size: var(--h5);
	}
</style>
