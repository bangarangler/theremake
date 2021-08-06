<script lang="ts">
	import setUp from '$static/imgForUses.png?w=300;600;1000&format=png&srcset';
	import laptopImg from '$static/laptop-undraw.svg?w=200;400;800&format=svg&srcset';
	import codingImg from '$static/undraw-coding.svg?w=200;400;800&format=svg&srcset';
	import programmingImg from '$static/undraw-programming.svg?w=200;400;800&format=svg&srcset';
	import botImg from '$static/undraw-firmware.svg?w=200;400;800&format=svg&srcset';
	import audioImg from '$static/undraw-audio.svg?w=200;400;800&format=svg&srcset';
	import ImageLoader from '$images/ImageLoader.svelte';
</script>

<section>
	<h1>Uses...</h1>
	<div class="card wrapper">
		<picture class="imgWrapper">
			<source media="(min-width:1000px)" srcset={setUp} />
			<ImageLoader
				ty="image/png"
				srcset={setUp}
				alt="Epic illistration from samji_illustrator!  Please check them out!"
			/>
		</picture>
	</div>
	<div class="usesGrid">
		<div class="computerWrapper">
			<div class="usesTitleWrapper">
				<h3>Computers</h3>
				<picture>
					<source media="(min-width: 600px)" srcset={laptopImg} />
					<ImageLoader ty="image/svg" srcset={laptopImg} alt="Computers!" />
				</picture>
			</div>
			<ul>
				<li>Lenovo ThinkPad P52</li>
				<li>Apple MacBook Pro (15-inch), 16GB Ram, Intel Core i7 Space Gray</li>
			</ul>
		</div>
		<div class="hardwareWrapper">
			<div class="usesTitleWrapper">
				<h3>Hardware</h3>
				<picture>
					<source media="(min-width: 600px)" srcset={codingImg} />
					<ImageLoader ty="image/svg" srcset={codingImg} alt="Hardware!" />
				</picture>
			</div>
			<ul>
				<li>Kensington Orbit Trackball Mouse</li>
				<li>Apple Magic TrackPad 2 (don't see much use anymore)</li>
				<a href="https://ergodox-ez.com/" target="_blank" rel="external"
					><li>ErgoDox split keyboard</li></a
				>
				<li>Handful of USB sticks (Linux Installers, Backups, misc)</li>
				<li>External SSD (not everything can live on github ; ) )</li>
				<li>LG Ultrawide Monitor (34 inch)</li>
			</ul>
		</div>
		<div class="softwareWrapper">
			<div class="usesTitleWrapper">
				<h3>Software</h3>
				<picture>
					<source media="(min-width: 600px)" srcset={botImg} />
					<ImageLoader ty="image/svg" srcset={botImg} alt="Software!" />
				</picture>
			</div>
			<ul>
				<a href="https://www.notion.so" target="_blank" rel="external"><li>Notion</li></a>
				<a href="https://ulauncher.io/" target="_blank" rel="external"
					><li>ULauncher (Linux) ~ Alfred (Mac)</li></a
				>
				<a href="https://getmailspring.com/" target="_blank" rel="external"
					><li>MailSpring (Email client)</li></a
				>
				<a href="https://www.mozilla.org/en-US/firefox/developer/" target="_blank" rel="external"
					><li>Firefox Dev Edition (Main Browser)</li></a
				>
				<li>Firefox | Chrome | Edge | Vivaldi | Safari (Dev Browsers)</li>
			</ul>
		</div>
		<div class="terminalWrapper">
			<div class="usesTitleWrapper">
				<h3>Terminal</h3>
				<picture>
					<source media="(min-width: 600px)" srcset={programmingImg} />
					<ImageLoader ty="image/svg" srcset={programmingImg} alt="Terminal!" />
				</picture>
			</div>
			<ul>
				<a href="https://github.com/alacritty/alacritty" target="_blank" rel="external"
					><li>Alacritty</li></a
				>
				<a href="https://github.com/tmux/tmux/wiki" target="_blank" rel="external"><li>Tmux</li></a>
				<a href="https://neovim.io/" target="_blank" rel="external"><li>Neovim</li></a>
				<a href="https://curl.se/" target="_blank" rel="external"><li>curl</li></a>
				<a href="https://github.com/bangarangler/stow-dotfiles" target="_blank" rel="external"
					><li>Dotfiles</li></a
				>
			</ul>
		</div>
		<div class="audioWrapper">
			<div class="usesTitleWrapper">
				<h3>Audio</h3>
				<picture>
					<source media="(min-width: 600px)" srcset={audioImg} />
					<!--<img loading="lazy" srcset={audioImg} alt="Computers!"
type="image/svg" />-->
					<ImageLoader ty="image/svg" srcset={audioImg} alt="Audio!" />
				</picture>
			</div>
			<ul>
				<li>Sony 1000xm3's</li>
				<a href="https://www.spotify.com" target="_blank" rel="external"><li>Spotify</li></a>
			</ul>
		</div>
	</div>
</section>
<p class="notAffiliate">
	Not affiliate links! Just wanted to share incase your interested in improving your terminal
	experience ; )
</p>

<style>
	section {
		margin: 0 auto;
		width: 90%;
	}
	@media (min-width: 1600px) {
		section {
			max-width: 1600px;
		}
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
	@media (min-width: 820px) {
		.wrapper {
			margin-bottom: 50px;
		}
	}
	@media (min-width: 1000px) {
		.wrapper {
			margin-bottom: 100px;
		}
	}
	section h1 {
		font-size: var(--h2);
		color: var(--darkAquaLightHotPink);
		margin-top: 30px;
		text-align: center;
	}
	@media (min-width: 460px) {
		section h1 {
			margin-top: unset;
		}
	}
	@media (min-width: 750px) {
		section h1 {
			font-size: var(--h1);
		}
	}
	section h3 {
		font-size: var(--h4);
		color: var(--dracPurpDarkGray);
		margin-top: 0;
		margin-bottom: 0;
		margin-right: 12px;
	}
	@media (min-width: 750px) {
		section h3 {
			font-size: var(--h3);
		}
	}
	.usesGrid {
		display: grid;
		grid-template-columns: minmax(0, 1fr);
		width: 80%;
		margin: 0 auto 100px;
	}
	@media (min-width: 750px) {
		.usesGrid {
			grid-template-columns: repeat(2, minmax(0, 1fr));
			grid-template-areas:
				'comp hard'
				'audio hard'
				'soft term'
				'soft .';
			grid-column-gap: 30px;
			width: 90%;
		}
		.computerWrapper {
			grid-area: comp;
		}
		.hardwareWrapper {
			grid-area: hard;
		}
		.softwareWrapper {
			grid-area: soft;
		}
		.terminalWrapper {
			grid-area: term;
		}
		.audioWrapper {
			grid-area: audio;
		}
	}
	@media (min-width: 820px) {
		.usesGrid {
			grid-template-columns: minmax(0, 1fr);
			grid-template-areas:
				'comp'
				'hard'
				'soft'
				'term'
				'audio';
		}
		.computerWrapper {
			margin: 30px auto 30px;
			width: 80%;
		}
		.hardwareWrapper {
			margin: 30px auto 30px;
			width: 80%;
		}
		.softwareWrapper {
			margin: 30px auto 30px;
			width: 80%;
		}
		.terminalWrapper {
			margin: 30px auto 30px;
			width: 80%;
		}
		.audioWrapper {
			margin: 30px auto 30px;
			width: 80%;
		}
	}
	@media (min-width: 1600px) {
		.computerWrapper {
			width: 90%;
		}
		.hardwareWrapper {
			margin: 30px auto 30px;
			width: 90%;
		}
		.softwareWrapper {
			margin: 30px auto 30px;
			width: 90%;
		}
		.terminalWrapper {
			margin: 30px auto 30px;
			width: 90%;
		}
		.audioWrapper {
			margin: 30px auto 30px;
			width: 90%;
		}
	}
	.usesTitleWrapper {
		display: grid;
		grid-template-columns: max-content 50px;
		justify-content: flex-center;
		align-items: center;
	}
	.usesTitleWrapper picture {
		max-width: 50px;
		margin-top: 10px;
	}
	@media (min-width: 820px) {
		.usesTitleWrapper {
			grid-template-columns: max-content 1fr;
		}
		.usesTitleWrapper picture {
			max-width: 200px;
			min-width: 200px;
			margin-top: unset;
			place-self: flex-end;
		}
	}
	@media (min-width: 1000px) {
		.usesTitleWrapper picture {
			max-width: 400px;
			min-width: 400px;
		}
	}
	ul {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		justify-content: center;
		flex-wrap: wrap;
		margin-bottom: 30px;
	}
	ul a {
		text-decoration: none;
		color: inherit;
		border-bottom: 1px solid var(--dracPurpDarkGray);
	}
	ul a:hover {
		color: var(--darkAquaLightHotPink);
		border-bottom: 1px solid var(--darkAquaLightHotPink);
	}
	li {
		list-style-type: disc;
	}
	ul li::marker {
		color: var(--darkAquaLightHotPink);
		font-size: 25px;
	}
	.notAffiliate {
		margin: auto auto 100px;
		text-align: center;
		width: 50%;
	}
</style>
