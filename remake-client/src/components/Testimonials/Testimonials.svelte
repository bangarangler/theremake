<script lang="ts">
	import TestimonialCard from '$components/Testimonials/TestimonialCard/TestimonialCard.svelte';
	import { testimonials } from '$components/Testimonials/testimonials';
	import IoIosArrowRoundForward from 'svelte-icons/io/IoIosArrowRoundForward.svelte';
</script>

<section class="testimonialWrapper">
	<div class="arrowContainer">
		<span><IoIosArrowRoundForward class="arrow" /></span>
	</div>
	{#each testimonials as testimonial}
		<TestimonialCard {testimonial} />
	{/each}
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
		right: 40px;
	}
	.testimonialWrapper .arrowContainer span {
		max-height: 50px;
		max-width: 50px;
		display: block;
	}
	.testimonialWrapper .arrowContainer span {
		height: 100%;
		width: 100%;
		object-fit: contain;
	}
</style>
