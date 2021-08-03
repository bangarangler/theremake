<script lang="ts">
	import TestimonialCard from '$components/Testimonials/TestimonialCard/TestimonialCard.svelte';
	import { testimonials } from '$components/Testimonials/testimonials';
</script>

<section>
	{#each testimonials as testimonial}
		<TestimonialCard {testimonial} />
	{/each}
</section>

<style>
	section {
		display: grid;
		grid-template-rows: minmax(0, 1fr);
		grid-gap: 30px;
		place-items: center;
		padding: 2em;
		max-width: 90%;
		margin: 0 auto;
	}
	@media (min-width: 500px) {
		section {
			max-width: 80%;
		}
	}
	@media (min-width: 910px) {
		section {
			grid-template-columns: repeat(3, minmax(0, 1fr));
			max-width: unset;
			margin: unset;
		}
	}
	@media (min-width: 1400px) {
		section {
			max-width: 1400px;
			margin: 0 auto;
		}
	}
</style>
