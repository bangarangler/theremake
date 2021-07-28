<script>
	import setUp from '$static/imgForUses.png?w=300;600;1600&format=png&srcset';
	import laptopImg from '$static/laptop-undraw.svg?w=200;400;800&format=svg&srcset';
	import codingImg from '$static/undraw-coding.svg?w=200;400;800&format=svg&srcset';
	import programmingImg from '$static/undraw-programming.svg?w=200;400;800&format=svg&srcset';
	import botImg from '$static/undraw-firmware.svg?w=200;400;800&format=svg&srcset';
	import audioImg from '$static/undraw-audio.svg?w=200;400;800&format=svg&srcset';
</script>

<section>
	<h1>Uses...</h1>
	<div class="card wrapper">
		<picture class="imgWrapper">
			<source media="(min-width:1000px)" srcset={setUp} />
			<img
				class="img1"
				srcset={setUp}
				type="image/png"
				alt="Epic illistration from samji_illustrator!  Please check them out!"
			/>
		</picture>
	</div>
	<div class="usesGrid">
		<div class="usesTitleWrapper">
			<h3>Computers</h3>
			<picture>
				<source media="(min-width: 600px)" srcset={laptopImg} />
				<img srcset={laptopImg} alt="Computers!" type="image/svg" />
			</picture>
		</div>
		<ul>
			<li>Lenovo ThinkPad P52</li>
			<li>Apple MacBook Pro (15-inch), 16GB Ram, Intel Core i7 Space Gray</li>
		</ul>
		<div class="usesTitleWrapper">
			<h3>Hardware</h3>
			<picture>
				<source media="(min-width: 600px)" srcset={codingImg} />
				<img srcset={codingImg} alt="Computers!" type="image/svg" />
			</picture>
		</div>
		<ul>
			<li>Trackball Mouse</li>
			<li>Apple Magic TrackPad (don't see much use anymore)</li>
			<li>ErgoDox EZ split keyboard</li>
			<li>Handful of USB sticks (Linux Installers, Backups, misc)</li>
			<li>External SSD (not everything can live on github ; ) )</li>
			<li>LG Ultrawide Monitor (34 inch)</li>
		</ul>
		<div class="usesTitleWrapper">
			<h3>Software</h3>
			<picture>
				<source media="(min-width: 600px)" srcset={botImg} />
				<img srcset={botImg} alt="Computers!" type="image/svg" />
			</picture>
		</div>
		<ul>
			<li>Notion</li>
			<li>ULauncher (Linux) Alfred (Mac)</li>
			<li>MailSpring (Email client)</li>
			<li>Firefox (Main Browser)</li>
			<li>Firefox | Chrome | Edge | Vivaldi | Safari (Dev Browsers)</li>
		</ul>
		<div class="usesTitleWrapper">
			<h3>Terminal</h3>
			<picture>
				<source media="(min-width: 600px)" srcset={programmingImg} />
				<img srcset={programmingImg} alt="Computers!" type="image/svg" />
			</picture>
		</div>
		<ul>
			<li>Alacritty</li>
			<li>Tmux</li>
			<li>Neovim</li>
		</ul>
		<div class="usesTitleWrapper">
			<h3>Audio</h3>
			<picture>
				<source media="(min-width: 600px)" srcset={audioImg} />
				<img srcset={audioImg} alt="Computers!" type="image/svg" />
			</picture>
		</div>
		<ul>
			<li>Sony 1000xm3's</li>
			<li>Spotify</li>
		</ul>
	</div>
</section>

<style>
	section {
		margin: 0 auto;
		width: 90%;
	}
	.wrapper {
		display: grid;
		grid-template-columns: minmax(0 1fr);
		margin: 25px auto;
		width: 90%;
		padding: 0.5em;
	}
	@media (min-width: 500px) {
		.wrapper {
			max-width: 500px;
		}
	}
	.imgWrapper img {
		width: 100%;
		height: 100%;
		object-fit: contain;
	}
	section h1 {
		font-size: var(--h2);
		color: var(--darkAquaLightHotPink);
		margin-top: 30px;
	}
	section h3 {
		font-size: var(--h4);
		color: var(--dracPurpDarkGray);
		margin-top: 0;
		margin-bottom: 0;
		margin-right: 12px;
	}
	.usesGrid {
		display: grid;
		grid-template-columns: minmax(0, 1fr);
		width: 90%;
		margin: 0 auto;
	}
	.usesTitleWrapper {
		display: grid;
		grid-template-columns: max-content 70px;
		justify-content: flex-center;
		align-items: center;
	}
	.usesTitleWrapper picture {
		max-width: 70px;
	}
	.usesTitleWrapper picture img {
		height: 100%;
		width: 100%;
		object-fit: contain;
	}
	ul {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		justify-content: center;
		flex-wrap: wrap;
		margin-bottom: 30px;
	}
	li {
		list-style-type: disc;
	}
	ul li::marker {
		color: var(--darkAquaLightHotPink);
		font-size: 25px;
	}
</style>
