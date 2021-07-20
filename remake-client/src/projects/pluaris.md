---
projectTitle: Pluaris
slug: pluaris
projectRole: Full Stack
projectDescription: 'Pluaris enables its users to quickly analyze boatloads of data, whether structured or unstructured, extract actionable insights from all the critical touch points happening within their ecosystem. '
techUsed: ['python', 'node', 'typescript', 'graphql', 'typescript', 'redis', 'mongo']
---

<script>
	import {fly, fade} from 'svelte/transition'
  import {onMount} from 'svelte'
  import pluarisDesktop from '$images/pluaris-dashboard.jpg'
  import myMemoryMobile from '$images/mymemory-pluaris-mobile.jpg'
  import myMemoryInfoMobile from '$images/mymemory-info-mobile.jpg'
  let scrollY;
	let ani = false;
  let exampleAnimate = false;
  let deploymentAnimate = false;
  let responsibilityAnimate = false;
	$: if (scrollY > 350) {
  console.log("running...")
		exampleAnimate = true
    }
    $: if (scrollY > 620) {
    deploymentAnimate = true
    }
    $: if (scrollY > 1100) {
    responsibilityAnimate = true
    }
    $: if(scrollY === 0) {
    exampleAnimate = false
    deploymentAnimate = false
    responsibilityAnimate = false
    }
    onMount(() => {
  ani = true
  })
</script>

<svelte:window bind:scrollY />

<article>
<div class="container">
<h1 class:display={ani}>{projectTitle}</h1>

<p class="headingText">Software Engineer • Engineering Manager • Architecture / UI / UX / Development • 2019 - Present</p>

</div>
<div class="card imgContainer">
<img class="img1" src={myMemoryMobile} alt="Mobile view for Pluaris My Memory
page" />
<img class="img2" src={myMemoryInfoMobile} alt="Mobile view for Pluaris My Memory page
information related to search.">
</div>
</article>

<style>
article {
  margin: 0 45px;
}
.container {
margin-bottom: 65px;
}
h1 {
background: linear-gradient(271deg,var(--hotpink) 30%, 50%,var(--aqua) 70%,#a162e8 94%);
background-clip: border-box;
-webkit-background-clip: text;
-webkit-text-fill-color: transparent;
opacity: 0;
}

.display {
animation: 1.2s ease dispalyAnimation;
opacity: 1;
}

@keyframes dispalyAnimation {
0% {
 opacity: 0;
 margin-left: 200px;
 transform: skewX(35deg);
}
100% {
 opacity: 1;
 margin-left: 0px;
 transform: skewX(0deg);
}
}

.headingText {
  max-width: -moz-fit-content;
  max-width: fit-content;
  padding-right: 5px;
  overflow: hidden;
  border-right: .15em solid var(--aqua);
  white-space: nowrap;
  animation:
    typing 6.5s steps(80, end),
    blink-caret .75s step-end infinite;
}
@keyframes typing {
  from { width: 0 }
  to { width: 100% }
}
@keyframes blink-caret {
  from, to { border-color: transparent }
  50% { border-color: var(--aqua); }
}

.imgContainer {
display: grid;
grid-template-columns: repeat(2, minmax(0, 1fr));
width: 100%;
place-items: center;
margin-bottom: 25px;
}
.img1 {
  max-height: 700px;
  border-radius: 12px;
}
.img2 {
  max-height: 700px;
  border-radius: 12px;
}
</style>
