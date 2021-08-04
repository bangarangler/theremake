<script>
	import ImageLoader from '$images/ImageLoader.svelte';
	import temp from '$static/pexels-2screen.jpg?w=400;800;1400&format=svg&srcset';
	import bNDad from '$static/bnDad.jpg?w=400;600;1000&format=jpg&srcset';
</script>

<section>
	<div class="custom-shape-divider-top">
		<svg
			data-name="Layer 1"
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 1200 120"
			preserveAspectRatio="none"
		>
			<path
				d="M0,0V46.29c47.79,22.2,103.59,32.17,158,28,70.36-5.37,136.33-33.31,206.8-37.5C438.64,32.43,512.34,53.67,583,72.05c69.27,18,138.3,24.88,209.4,13.08,36.15-6,69.85-17.84,104.45-29.34C989.49,25,1113-14.29,1200,52.47V0Z"
				opacity=".25"
				class="shape-fill"
			/>
			<path
				d="M0,0V15.81C13,36.92,27.64,56.86,47.69,72.05,99.41,111.27,165,111,224.58,91.58c31.15-10.15,60.09-26.07,89.67-39.8,40.92-19,84.73-46,130.83-49.67,36.26-2.85,70.9,9.42,98.6,31.56,31.77,25.39,62.32,62,103.63,73,40.44,10.79,81.35-6.69,119.13-24.28s75.16-39,116.92-43.05c59.73-5.85,113.28,22.88,168.9,38.84,30.2,8.66,59,6.17,87.09-7.5,22.43-10.89,48-26.93,60.65-49.24V0Z"
				opacity=".5"
				class="shape-fill"
			/>
			<path
				d="M0,0V5.63C149.93,59,314.09,71.32,475.83,42.57c43-7.64,84.23-20.12,127.61-26.46,59-8.63,112.48,12.24,165.56,35.4C827.93,77.22,886,95.24,951.2,90c86.53-7,172.46-45.71,248.8-84.81V0Z"
				class="shape-fill"
			/>
		</svg>
		<div class="bio">
			<div class="imgWrap">
				<picture>
					<source media="(min-width:1000px)" srcset={temp} />
					<ImageLoader
						ty="image/jpg"
						srcset={temp}
						alt="developers desk with
two screens and nice background."
					/>
				</picture>
			</div>
			<div class="bioTitle">
				<h2>Jonathan Palacio</h2>
				<p>Software Engineer ~ Terminal Enthusiast</p>
				<div class="bnDadWrap">
					<picture>
						<source media="(min-width:1000px)" srcset={bNDad} />
						<ImageLoader
							ty="image/jpg"
							srcset={bNDad}
							alt="Jon and his son!  All the way to the moon!"
						/>
					</picture>
				</div>
			</div>
		</div>
	</div>
</section>

<div class="bioText">
	<h2>What does Jon do?</h2>
	<p>
		Jon is from North Carolina where he works as a Software Engineer at Nowigence Inc. He has
		significant experience with Web technologies, building and securing API's and developing
		performant and scalable backends, putting together frontends to give users a phenominal user
		experience and has recently fell down the Gopher (GoLang) rabbit hole! In his spare time, Jon
		enjoys Netflix, reading, and oddly enough... more coding!
	</p>
	<p>
		Jon's a proud father, husband, and Linux nerd! When not hard at work you will find him spending
		time with his family; Quite possibly multi-tasking and reading blog posts on anything and
		everything tech related.
	</p>
</div>

<style>
	section {
		position: relative;
		min-height: 100vh;
		background: linear-gradient(to bottom, var(--hotpink), transparent);
		max-width: 100%;
		margin: 0 auto;
	}
	@media (min-width: 500px) {
		section {
			max-width: 90%;
		}
	}
	@media (min-width: 600px) {
		section {
			max-width: 80%;
			margin: 0 auto;
		}
	}
	@media (min-width: 1060px) {
		section {
			max-width: unset;
			margin: unset;
		}
	}
	@media (min-width: 1600px) {
		section {
			max-width: 1600px;
			margin: 0 auto;
		}
	}
	.custom-shape-divider-top {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		overflow: hidden;
		line-height: 0;
	}

	.custom-shape-divider-top svg {
		position: relative;
		display: block;
		width: calc(100% + 1.3px);
		height: 185px;
		transform: rotateY(180deg);
	}
	.custom-shape-divider-top .shape-fill {
		fill: var(--headerBG);
	}
	.bio {
		display: grid;
		grid-template-columns: minmax(0, 1fr);
		place-items: center;
	}
	@media (min-width: 1060px) {
		.bio {
			grid-template-columns: repeat(2, minmax(0, 1fr));
			place-items: center;
			margin-top: 20px;
		}
	}
	.bio .imgWrap {
		background: var(--sweetImgWrapper);
		/* opacity: 0.7; */
		border-radius: 12px;
		padding: 0.5em;
		max-width: 350px;
		margin-top: 50px;
		grid-row: 2;
	}
	@media (min-width: 1060px) {
		.bio .imgWrap {
			max-width: 450px;
			grid-row: unset;
		}
	}
	@media (min-width: 1300px) {
		.bio .imgWrap {
			max-width: 600px;
			transform: rotate(-12deg);
		}
	}

	.bioTitle {
		border-radius: 12px;
		padding: 1em;
		display: flex;
		align-items: center;
		flex-direction: column;
		justify-content: center;
		margin-top: 0;
		padding-top: 0;
		box-shadow: var(--level4);
		background: var(--sweetImgWrapper);
		font-family: var(--slantText);
		text-align: center;
		/* max-width: 450px; */
		flex-wrap: wrap;
		margin-top: 50px;
	}
	@media (min-width: 460px) {
		.bioTitle {
			letter-spacing: 2px;
		}
	}
	@media (min-width: 1060px) {
		.bioTitle {
			max-width: unset;
			transform: rotate(-12deg);
			margin-bottom: 100px;
			margin-top: unset;
		}
	}
	@media (min-width: 1300px) {
		.bioTitle {
			transform: rotate(12deg);
			margin-bottom: 100px;
		}
	}
	.bioTitle h2 {
		margin-top: 15px;
		color: var(--sweetImgWrapText);
		font-size: var(--h4);
		text-align: center;
	}
	@media (min-width: 1060px) {
		.bioTitle h2 {
			font-size: var(--h3);
			max-width: 300px;
		}
	}
	@media (min-width: 1300px) {
		.bioTitle h2 {
			font-size: var(--h2);
		}
	}
	.bioText {
		width: 80%;
		margin: 50px auto 100px;
	}
	@media (min-width: 500px) {
		.bioText {
			margin: 130px auto 100px;
		}
	}
	@media (min-width: 1060px) {
		.bioText {
			width: 50%;
		}
	}
	.bioText h2 {
		margin-bottom: 40px;
		color: var(--dracPurpDarkGray);
		font-size: var(--h3);
	}
	@media (min-width: 980px) {
		.bioText h2 {
			font-size: var(--h2);
		}
	}
	.bioTitle .bnDadWrap {
		max-width: 300px;
		/* background: var(--sweetImgWrapper); */
	}
	@media (min-width: 1060px) {
		.bioTitle .bnDadWrap {
			max-width: 300px;
		}
	}
	@media (min-width: 1300px) {
		.bioTitle .bnDadWrap {
			max-width: 380px;
		}
	}
</style>
