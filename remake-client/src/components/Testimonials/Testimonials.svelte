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
		grid-template-columns: repeat(3, minmax(0, 1fr));
		grid-gap: 30px;
		place-items: center;
		padding: 2em;
	}
</style>
