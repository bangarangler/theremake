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
		<textarea name="message" id="message" cols="30" rows="10" />
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
			<p>Put something epic here Jon</p>
		</div>
		<!--<h2>Fill this out!</h2>-->
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
	}
	.wrapper .stuff {
		grid-area: stuff;
	}
	.imgWrapper {
		/* height: 100%; */
		/* width: 100%; */
		background: var(--darkGray);
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
	}
</style>
