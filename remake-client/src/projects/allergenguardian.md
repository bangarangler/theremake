---
projectTitle: Allergen Guardian
slug: allergenguardian
projectRole: Full Stack
projectDescription: 'Shop | Scan | Safe'
techUsed: ['next.js', 'node', 'typescript', 'mongodb', 'react', 'graphql', 'redis', 'fastify']
---

<script>
	import {fly, fade} from 'svelte/transition'
  import {onMount} from 'svelte'
  import test from '$images/beforeDawnTempLogo.svg'
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
    onMount(() => {
  ani = true
  })
</script>

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
display: flex;
width: 100%;
align-items: center;
justify-content: center;
gap: 40px;
margin-bottom: 25px;
}
.img1 {
  max-height: 300px;
  max-width:600px;
}
.img2 {
  height: 300px;
  width: 600px;
  max-height: 300px;
  max-width:600px;
  background: hotpink;
  border-radius: 12px;
}
.explanationContainer {
display: flex;
flex-direction: column;
align-items: flex-start;
width: 50%;
margin: 0 auto;
}
.explanationContainer h3 {
opacity: 0;
}
.explanationContainer h2 {
text-align: left;
}
.slideInLeft {
  animation: 1.2s ease slideInLeft;
  opacity: 1 !important;
}

.slideInRight {
  animation: 1.2s ease slideInRight;
  opacity: 1 !important;
}

@keyframes slideInLeft {
  0% {
   opacity: 0;
   margin-left: 200px;
  }
  100% {
  opacity: 1;
  margin-left: 0px;
  }
}

@keyframes slideInRight {
  0% {
   opacity: 0;
   margin-left: -200px;
  }
  100% {
  opacity: 1;
  margin-left: 0px;
  }
}
.fakeVideo {
height: 600px;
width: 800px;
max-height: 500px;
max-width: 800px;
background: purple;
margin: 25px auto 0 auto;
}
.fakeCaption p {
margin: 15px auto 0 auto;
text-align: center;
font-family: var(--slantText);
}
.projectInfoContaner {
display: grid;
grid-template-columns: minmax(0, 1fr) minmax(0, 1fr);
width: 50%;
margin: 65px auto 0 auto;
place-items: center;
grid-column-gap: 30px;
}
.projectInfoContaner h6 {
color: var(--lightGray);
border-bottom: 1px solid var(--hotpink);
width: -moz-fit-content;
width: fit-content;
}
</style>

<svelte:window bind:scrollY />

<article>
<div class="container">
<h1 class:display={ani}>{projectTitle}</h1>

<p class="headingText">Engineer • Partner / UI / UX / Development • 2021</p>

</div>
<div class="card imgContainer">
<img class="img1" src={test} alt="testing" />
<div class="img2"/>
</div>

<div class="explanationContainer">
<h2>What's the problem?</h2>
<p>Aliquam molestie vestibulum elit in feugiat. Nam suscipit justo erat. Donec hendrerit volutpat dolor. In ullamcorper laoreet lacus nec mattis. In luctus sem id ornare scelerisque. Fusce in magna mattis, facilisis leo id, tempor leo. Fusce non lobortis dolor, maximus bibendum velit. Curabitur scelerisque urna a metus tempus euismod. In pulvinar libero risus, id volutpat magna sollicitudin vel. Mauris a magna metus. In sit amet lorem nec erat convallis fringilla et ut nisl. Vivamus vel consequat risus. Suspendisse lorem libero, laoreet quis fringilla sit amet, accumsan a augue.</p>

<h3 class:slideInLeft={exampleAnimate}>Example</h3>
<p>Donec leo lorem, tristique quis ipsum a, suscipit commodo diam. Vestibulum cursus odio augue. Nunc laoreet magna eu tempor malesuada. Donec blandit sapien et auctor suscipit. Sed ac congue eros. Sed in dictum orci. Pellentesque placerat, dolor in dapibus interdum, urna velit semper velit, vel tincidunt augue arcu at odio.</p>

<h3 class:slideInRight={deploymentAnimate}>Deployment</h3>
<p>Aenean et laoreet nisi. Phasellus sagittis mauris at volutpat aliquam. Aenean vitae arcu quis est lobortis rutrum vitae sit amet ligula. Interdum et malesuada fames ac ante ipsum primis in faucibus. Nam posuere, libero sed fringilla eleifend, massa nulla porta ligula, a porta mi ligula et turpis. Sed scelerisque urna vel massa finibus euismod. In non mauris eros. In dapibus nunc ac quam auctor, sit amet fermentum diam vulputate. Praesent eu porttitor tortor. Nam sed nibh quam. Quisque molestie venenatis neque vel placerat. Donec nec euismod justo. Praesent a pharetra metus. Fusce eget tellus urna. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae</p>

<h3 class:slideInLeft={responsibilityAnimate}>Responsibility</h3>
<p>Nam mi est, dapibus ut aliquam ut, consequat vel nulla. Duis vitae pellentesque neque, ut tristique est. Morbi ut ligula fermentum, venenatis odio vel, laoreet mauris. Integer fermentum libero tortor, dictum semper mi tempus quis. Aliquam in lacinia leo. Morbi nec lobortis augue. Praesent sem nulla, accumsan id est at, facilisis molestie dolor. Aliquam semper sed felis a mollis. Vestibulum odio velit, lacinia quis felis vitae, convallis dictum felis. Maecenas non auctor justo.</p>
</div>

<div class="fakeVideo"/>
<div class="fakeCaption"><p>Something about the video here</p></div>

<div class="explanationContainer">
<h2>Our Solution</h2>
<p>Lorem ipsum dolor sit amet. </p>
<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris leo nulla, ullamcorper sit amet luctus vitae, ornare vel sem. Nam dapibus aliquet dui et venenatis. Ut fringilla nisl eget diam lobortis, at iaculis sem aliquam. Nam fringilla tincidunt augue id rhoncus. Mauris accumsan sem nulla, non vulputate quam eleifend at. Curabitur consequat consectetur sem ut tincidunt. Vivamus hendrerit facilisis risus sit amet tincidunt. Donec et ante viverra, eleifend dui vitae, blandit risus. Nam ac porta lectus, sit amet vehicula elit. </p>
<p>Duis eget libero aliquet, consequat leo fermentum, pulvinar risus. Vivamus at enim fringilla est tristique lobortis at nec velit. Pellentesque mollis porttitor est vitae volutpat. Donec ipsum nibh, scelerisque at tellus vel, consequat interdum lectus. Duis blandit, urna iaculis elementum rhoncus, turpis nisi imperdiet neque, nec tempus ipsum ante sed erat. Praesent eu fermentum lectus, vitae rhoncus metus. Nunc a mauris a sem sollicitudin iaculis. Nullam cursus felis ipsum, sit amet consequat eros cursus sit amet. </p>
<p>Duis eget libero aliquet, consequat leo fermentum, pulvinar risus. Vivamus at enim fringilla est tristique lobortis at nec velit. Pellentesque mollis porttitor est vitae volutpat. Donec ipsum nibh, scelerisque at tellus vel, consequat interdum lectus. Duis blandit, urna iaculis elementum rhoncus, turpis nisi imperdiet neque, nec tempus ipsum ante sed erat. Praesent eu fermentum lectus, vitae rhoncus metus. Nunc a mauris a sem sollicitudin iaculis. Nullam cursus felis ipsum, sit amet consequat eros cursus sit amet. </p>
</div>

<div class="fakeVideo"/>
<div class="fakeCaption"><p>Something about the video here</p></div>

<div class="card projectInfoContaner">
<div>
<h6>Position</h6>
<p>{projectRole}</p>
<h6>Organization</h6>
<p>J&JStudio</p>
<h6>Year</h6>
<p>2020</p>
</div>
<div>
<h6>Work</h6>
<ul>
{#each techUsed as tu}
<li>{tu}</li>
{/each}
</ul>
</div>
</div>

</article>
