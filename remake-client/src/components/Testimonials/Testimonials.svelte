<script lang="ts">
	import TestimonialCard from '$components/Testimonials/TestimonialCard/TestimonialCard.svelte';
	import { testimonials } from '$components/Testimonials/testimonials';
	import { paginate } from '$components/Testimonials/testimonialPagination';
	import IoIosArrowRoundForward from 'svelte-icons/io/IoIosArrowRoundForward.svelte';
	import IoIosArrowRoundBack from 'svelte-icons/io/IoIosArrowRoundBack.svelte';

	$: pageInfo = paginate(testimonials);
</script>

<section class="testimonialWrapper">
	<div class="arrowContainer">
		<p
			class:disabled={!pageInfo?.prePage}
			on:click={() => (pageInfo = paginate(testimonials, pageInfo.prePage))}
		>
			Prev
		</p>
		<span
			class:disabled={!pageInfo?.prePage}
			on:click={() => (pageInfo = paginate(testimonials, pageInfo.prePage))}
			><IoIosArrowRoundBack class="arrow" /></span
		>
		<span
			class:disabled={!pageInfo?.nextPage}
			on:click={() => (pageInfo = paginate(testimonials, pageInfo.nextPage))}
			><IoIosArrowRoundForward class="arrow" /></span
		>
		<p
			class:disabled={!pageInfo?.nextPage}
			on:click={() => (pageInfo = paginate(testimonials, pageInfo.nextPage))}
		>
			Next
		</p>
	</div>
	{#each pageInfo.data.slice(0, 3) as testimonial}
		<TestimonialCard {testimonial} />
	{/each}
	<div class="notFixed">
		<p
			class:disabled={!pageInfo?.prePage}
			on:click={() => (pageInfo = paginate(testimonials, pageInfo.prePage))}
		>
			Prev
		</p>
		<span
			class:disabled={!pageInfo?.prePage}
			on:click={() => (pageInfo = paginate(testimonials, pageInfo.prePage))}
			><IoIosArrowRoundBack class="arrow" /></span
		>
		<span
			class:disabled={!pageInfo?.nextPage}
			on:click={() => (pageInfo = paginate(testimonials, pageInfo.nextPage))}
			><IoIosArrowRoundForward class="arrow" /></span
		>
		<p
			class:disabled={!pageInfo?.nextPage}
			on:click={() => (pageInfo = paginate(testimonials, pageInfo.nextPage))}
		>
			Next
		</p>
	</div>
</section>

<style>
	.testimonialWrapper {
		display: grid;
		grid-template-rows: minmax(0, 1fr);
		grid-gap: 30px;
		place-items: center;
		padding: 2em;
		max-width: 90%;
		margin: 0 auto;
		position: relative;
	}
	@media (min-width: 500px) {
		.testimonialWrapper {
			max-width: 80%;
		}
	}
	@media (min-width: 910px) {
		.testimonialWrapper {
			grid-template-columns: repeat(3, minmax(0, 1fr));
			max-width: unset;
			margin: unset;
		}
	}
	@media (min-width: 1400px) {
		.testimonialWrapper {
			max-width: 1400px;
			margin: 0 auto;
		}
	}
	.testimonialWrapper .arrowContainer {
		position: absolute;
		top: 15px;
		right: 20px;
		text-align: center;
		display: flex;
		align-items: center;
		justify-content: space-between;
		/* border: 1px solid red; */
		width: max-content;
	}
	.notFixed {
		/* position: static; */
		text-align: center;
		display: flex;
		align-items: center;
		justify-content: space-between;
		/* border: 1px solid red; */
		width: max-content;
		place-self: end;
		grid-area: auto / -2;
	}
	@media (min-width: 1100px) {
		.notFixed {
			width: 250px;
		}
	}
	@media (min-width: 700px) {
		/* .testimonialWrapper #notFixed { */
		/* 	position: static; */
		/* 	grid-column: -4 / -1; */
		/* } */
	}
	@media (min-width: 1100px) {
		.testimonialWrapper .arrowContainer {
			top: 15px;
			right: 40px;
			width: 250px;
		}
	}
	.testimonialWrapper .arrowContainer span:hover {
		color: var(--darkAquaLightHotPink);
	}
	.testimonialWrapper .notFixed span:hover {
		color: var(--darkAquaLightHotPink);
	}
	.testimonialWrapper .arrowContainer:hover {
		cursor: pointer;
	}
	.testimonialWrapper .notFixed:hover {
		cursor: pointer;
	}
	.testimonialWrapper .arrowContainer span {
		max-height: 50px;
		max-width: 50px;
		display: block;
		animation: animate 15s linear infinite;
		height: 100%;
		width: 100%;
		object-fit: contain;
	}
	.testimonialWrapper .notFixed span {
		max-height: 50px;
		max-width: 50px;
		display: block;
		animation: animate 15s linear infinite;
		height: 100%;
		width: 100%;
		object-fit: contain;
	}
	.testimonialWrapper .arrowContainer p {
		margin-bottom: 0;
		font-size: var(--smallTextSize);
	}
	.testimonialWrapper .notFixed p {
		margin-bottom: 0;
		font-size: var(--smallTextSize);
	}
	.testimonialWrapper .arrowContainer p:hover {
		color: var(--darkAquaLightHotPink);
	}
	.testimonialWrapper .notFixed p:hover {
		color: var(--darkAquaLightHotPink);
	}
	.disabled {
		cursor: not-allowed;
		color: var(--projH6DarkLight);
	}
	@keyframes animate {
		0% {
			transform: rotate(0deg);
		}
		50% {
			transform: rotate(0deg);
		}
		55% {
			transform: rotate(360deg);
		}
		100% {
			transform: rotate(360deg);
		}
	}
</style>
