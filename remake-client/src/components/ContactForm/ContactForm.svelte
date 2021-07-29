<script>
	import contactImg from '$static/contact-image.svg?w=600;700;1000&format=svg&srcset';
</script>

<div class="wrapper">
	<form>
		<labe>Name</labe>
		<input type="text" placeholder="Jane Doe" />
		<labe>Email</labe>
		<input type="email" placeholder="jane.doe@abracadabra.com" />
		<labe>Message</labe>
		<textarea
			name="message"
			id="message"
			cols="30"
			rows="50"
			placeholder="My deep dark secret is..."
		/>
		<div>
			<h3>Let's Build Something!</h3>
		</div>
		<div class="btnWrapper">
			<button class="glow-on-hover" type="button">Send to Jon</button>
			<button class="cancel">Cancel</button>
		</div>
	</form>
	<div class="stuff">
		<div class="imgWrapper">
			<h2>Excited to chat!</h2>
			<picture>
				<source media="(min-width:1000px)" srcset={contactImg} />
				<img
					class="img1"
					srcset={contactImg}
					type="image/svg"
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
		grid-template-columns: repeat(2, minmax(0, 1fr));
		width: 100%;
		height: 100%;
		grid-gap: 15px;
		grid-template-areas: 'stuff form';
	}
	.wrapper form {
		grid-area: form;
		display: flex;
		flex-direction: column;
		width: 75%;
		margin: 0 auto;
	}
	.wrapper form input::placeholder {
		color: var(--reverseTextColor);
	}
	.wrapper form textarea::placeholder {
		color: var(--reverseTextColor);
	}
	.wrapper form button {
		width: max-content;
		margin-bottom: 10px;
		/* background: var(--hotpink); */
	}
	.wrapper form div h3 {
		font-family: var(--slantText);
		transform: skewY(-15deg);
		border-bottom: 1px solid var(--darkAquaLightHotPink);
		margin: 120px auto 200px;
	}
	.wrapper .stuff {
		grid-area: stuff;
	}
	.imgWrapper {
		/* height: 100%; */
		/* width: 100%; */
		background: var(--cardBG);
		min-height: 80vh;
		padding: 2em;
		display: grid;
		grid-template-rows: minmax(0, auto) minmax(0, 1fr) minmax(0, auto);
		grid-template-areas:
			'top'
			'middle'
			'bottom';
		/* display: flex; */
		/* flex-direction: column; */
		/* align-items: center; */
		/* justify-content: center; */
	}
	.imgWrapper h2 {
		grid-area: top;
		text-align: center;
		margin: 0;
		color: var(--darkAquaLightHotPink);
	}
	.imgWrapper picture {
		grid-area: middle;
		/* max-width: 400px; */
	}
	.imgWrapper picture img {
		height: 100%;
		width: 100%;
		object-fit: contain;
	}
	.imgWrapper p {
		grid-area: bottom;
		text-align: center;
		margin: 0;
		font-family: var(--slantText);
		letter-spacing: 1.5px;
		color: var(--darkAquaLightHotPink);
	}
	.btnWrapper {
		display: flex;
		width: 80%;
		align-items: center;
		justify-content: space-evenly;
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
</style>
