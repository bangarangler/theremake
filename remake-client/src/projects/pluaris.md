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
  import pluarisDesktop from '$images/pluaris-dashboard.jpg?w=600;700;800&format=jpg&srcset'
  import myMemoryMobile from '$static/mymemory-pluaris-mobile.jpg?w=500;600;700&format=jpg&srcset'
  import myMemoryInfoMobile from '$images/mymemory-info-mobile.jpg?w=500;600;700&format=jpg&srcset'
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
    <picture>
      <img class="img1" srcset={myMemoryMobile} type="image/jpg" alt="Mobile view for Pluaris My Memory page" />
    </picture>
    <picture>
      <img class="img2" srcset={myMemoryInfoMobile} type="image/jpg" alt="Mobile view for Pluaris My Memory page information related to search.">
    </picture>
</div>

<div class="explanationContainer">

<h2>What we do?</h2>

_Pluaris_ specializes in reading and comprehending data, analyzing cause and effect, identifying benchmarks and measuring performance against them, tracing and linking intelligence by topics, extracting critical intelligence, alerting, answering questions on-the-fly, and synthesizing outputs. This saves each employee an average of 2 hours per day. It accelerates the pace of business resulting in revenue growth and increased profitability.

<h3 class:slideInLeft={exampleAnimate}>Backend</h3>

_Pluaris_ uses a few different backends to accomplish it's goals and deliver
for our customers in the **B2B** (Business to Business) and **B2C** (Business to Consumer) markets.

My primary responsibility is delivering **real time** information for our users
along with a beautful intuitive **UX** (user experience)!

What technology do we need to pull off such a task? Enter <span class="nodejs">NodeJS</span> and friends!

We started off by revamping the **Express** framework and giving it some amazing
superpowers by connecting it with **Apollo Server** to develop and deliver a <span class="graphql">GraphQL</span> API along with our REST API's. Which allowed us to address a fairly serious over fetching issue.

We have a ton of meta data that the machine needs inorder to tweak models and do all sorts of fancy **NLP** (Natural Language Processing) &amp; **Machine Learning**; However our users never needed or cared about a "shocking" amount of it.

Thats not to say it wasn't important data. Actually quite the opposite! However
there is zero reason to have giant objects or massive amounts of meta data being sent over the net (possibly ending up on a mobile client).

We created our <span class="graphql">GraphQL</span> API with two things in mind.
Performance and obliterating our over fetching issue.

For the DataScience side of the coin we use <span class="python">Python</span> along with a _Flask_ API. I won't waste time name droping all of the DataScience tools and libraries we utilize as this is not my area of experteise. However; I have pitched in on the _Flask_ API and helped on and off with things such as (architecture, preformance, deployment, debugging, scripting).

<h3 class:slideInRight={deploymentAnimate}>Deployment</h3>
<p>Aenean et laoreet nisi. Phasellus sagittis mauris at volutpat aliquam. Aenean vitae arcu quis est lobortis rutrum vitae sit amet ligula. Interdum et malesuada fames ac ante ipsum primis in faucibus. Nam posuere, libero sed fringilla eleifend, massa nulla porta ligula, a porta mi ligula et turpis. Sed scelerisque urna vel massa finibus euismod. In non mauris eros. In dapibus nunc ac quam auctor, sit amet fermentum diam vulputate. Praesent eu porttitor tortor. Nam sed nibh quam. Quisque molestie venenatis neque vel placerat. Donec nec euismod justo. Praesent a pharetra metus. Fusce eget tellus urna. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae</p>

<h3 class:slideInLeft={responsibilityAnimate}>Responsibility</h3>
<p>Nam mi est, dapibus ut aliquam ut, consequat vel nulla. Duis vitae pellentesque neque, ut tristique est. Morbi ut ligula fermentum, venenatis odio vel, laoreet mauris. Integer fermentum libero tortor, dictum semper mi tempus quis. Aliquam in lacinia leo. Morbi nec lobortis augue. Praesent sem nulla, accumsan id est at, facilisis molestie dolor. Aliquam semper sed felis a mollis. Vestibulum odio velit, lacinia quis felis vitae, convallis dictum felis. Maecenas non auctor justo.</p>
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
em {
  font-family: var(--slantText);
  letter-spacing: .2em;
  font-size: var(--h6);
  background: linear-gradient(271deg, #a162e8 30%, 50%,var(--aqua)
  70%,var(--hotpink) 94%);
  background-clip: border-box;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}
strong {
font-weight: bold;
color: var(--aqua);
}
.python {
 color: #4B8BBE;
}
.nodejs {
color: #68a063;
}
.graphql {
color: #e535ab;
}
</style>
