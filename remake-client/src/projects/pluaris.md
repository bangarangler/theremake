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
  import pluarisDesktop from '$images/pluaris-dashboard.jpg?w=600;700;1600&format=jpg&srcset'
  import myMemoryMobile from '$static/mymemory-pluaris-mobile.jpg?w=200;400;700&format=jpg&srcset'
  import myMemoryInfoMobile from '$images/mymemory-info-mobile.jpg?w=200;400;700&format=jpg&srcset'
  let scrollY;
	let animation = false;
  let backendAnimate = false;
  let frontendAnimate = false;
  let devOpsAnimate = false;
	$: if (scrollY > 350) {
  console.log("running...")
		backendAnimate = true
    }
    $: if (scrollY > 620) {
    frontendAnimate = true
    }
    $: if (scrollY > 1100) {
    devOpsAnimate = true
    }
    $: if(scrollY === 0) {
    backendAnimate = false
    frontendAnimate = false
    devOpsAnimate = false
    }
    onMount(() => {
  animation = true
  })
</script>

<svelte:window bind:scrollY />

<article>
<div class="container">
<h1 class:display={animation}>{projectTitle}</h1>

<p class="headingText">Software Engineer • Engineering Manager • Architecture / UI / UX / Development • 2019 - Present</p>

</div>
  <div class="card imgContainer">
    <picture>
      <source media="(min-width:1200px)" srcset={myMemoryMobile}>
        <img class="img1" srcset={myMemoryMobile} type="image/jpg" alt="Mobile view for Pluaris My Memory page" />
    </picture>
    <picture>
      <source media="(min-width:1200px)" srcset={myMemoryInfoMobile}>
      <img class="img2" srcset={myMemoryInfoMobile} type="image/jpg" alt="Mobile view for Pluaris My Memory page information related to search.">
    </picture>
</div>

<div class="explanationContainer">

<h2>What we do?</h2>

_Pluaris_ specializes in reading and comprehending data, analyzing cause and effect, identifying benchmarks and measuring performance against them, tracing and linking intelligence by topics, extracting critical intelligence, alerting, answering questions on-the-fly, and synthesizing outputs. This saves each employee an average of 2 hours per day. It accelerates the pace of business resulting in revenue growth and increased profitability.

<h3 class:slideInLeft={backendAnimate}>Backend</h3>

We are transitioning from a **Monolith** to a more **Microservice** approach. Some
services are now ran with <span class="docker">Docker</span>. With a goal of
having our **Microservices** containerized and running in <span
class="k8s">Kubernetes</span>.

We use <span class="nodejs">NodeJS</span>, ~~ExpressJS~~ Fastify, <span class="typescript">TypeScript</span>, <span class="graphql">GraphQL</span>, <span class="mongo">MongoDB</span>, <span class="redis">Redis</span>.

I'm a big fan of automation and have tried to streamline lots of different tasks
and workflows. ~~GoLang~~ <span class="go">Go</span> has been a fantastic
addition and a joy to work with! Not that <span class="nodejs">NodeJS</span> is
a slouch in the speed department but I have a feeling more **microservices**
will be written in <span class="go">Go</span> in the future.

For the DataScience side of the coin we use <span class="python">Python</span> along with a _Flask_ API. I won't waste time name droping all of the DataScience tools and libraries we utilize as this is not my area of experteise. However; I have pitched in on the _Flask_ API and helped on and off with things such as (architecture, preformance, deployment, debugging, scripting).

<!-- _Pluaris_ uses a few different backends to accomplish it's goals and deliver -->
<!-- for our customers in the **B2B** (Business to Business) and **B2C** (Business to Consumer) markets. -->
<!--  -->
<!-- My primary responsibility is delivering **real time** information for our users -->
<!-- along with a beautful intuitive **UX** (user experience)! -->

<!-- What technology do we need to pull off such a task? Enter <span class="nodejs">NodeJS</span> and friends! -->
<!--  -->
<!-- We started off by revamping the **Express** framework and giving it some amazing -->
<!-- superpowers by connecting it with **Apollo Server** to develop and deliver a <span class="graphql">GraphQL</span> API along with our REST API's. Which allowed us to address a fairly serious over fetching issue. -->
<!--  -->
<!-- We have a ton of meta data that the machine needs inorder to tweak models and do all sorts of fancy **NLP** (Natural Language Processing) &amp; **Machine Learning**; However our users never needed or cared about a "shocking" amount of it. -->
<!--  -->
<!-- Thats not to say it wasn't important data. Actually quite the opposite! However -->
<!-- there is zero reason to have giant objects or massive amounts of meta data being sent over the net (possibly ending up on a mobile client). -->
<!--  -->
<!-- We created our <span class="graphql">GraphQL</span> API with two things in mind. -->
<!-- Performance and obliterating our over fetching issue. As a kicker we also wanted -->
<!-- to utilize subscriptions via *web-sockets* and get some static typing going on. -->
<!--  -->

<h3 class:slideInRight={frontendAnimate}>Frontend</h3>

On the frontend we are using <span class="react">React</span>, along with <span
class="typescript">TypeScript</span> to handle the data and deliver rich _user_
tailored expierences for our clients.

Your ~~data~~ **Internal Memory**,
available to recall and curate a more personalized view into the inner workings
of either your business or your brain in **real-time**.

<h3 class:slideInLeft={devOpsAnimate}>Dev Ops</h3>

Our current **cloud provider** is <span class="digitalocean">Digital Ocean</span>. We have some _automation_ scripts in place to pull in the new code, run some **tests**, re-build, and **deploy**.

Our <span class="nodejs">NodeJS</span> servers are spread across multiple <span
class="digitalocean">droplets</span> and we are **load-balancing** and using **reverse proxies**.

A key piece of kit in this set up has been _Caddy_ and man is it awesome! Much
cleaner and easier to work with than an <span class="nginx">Nginx</span> config.
It handles auto **SSL** renewal with minimal fuss. And automating different
servers to run different _Caddyfile(s)_ has been a fantastic experience.

As we contine moving everything for **production** into <span
class="docker">Docker</span> and finishing our move to <span
class="k8s">K8s</span> we will have the freedom to take our platform and deploy
it on <span class="aws">AWS</span>, <span class="googlecloud">Google Cloud</span>, <span class="azure">Azure</span>, or anywhere we need to go!

</div>

<div class="card oneImage">
  <picture>
  <source media="(min-width:1200px)" srcset={pluarisDesktop}>
    <img srcset={pluarisDesktop} type="image/jpg" alt="Pluaris application view of the desktop">
  </picture>
</div>
<div class="caption"><p>Dashboard view of Pluaris!</p></div>

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
.imgContainer picture {
  max-width: 300px;
}
.img1 {
  height: 100%;
  width: 100%;
  object-fit: contain;
  border-radius: 12px;
}
.img2 {
  height: 100%;
  width: 100%;
  object-fit: contain;
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
.typescript {
color: #007acc;
}
.mongo {
color: #4DB33D;
}
.redis {
color: #D82C20;
}
.docker {
color: #0db7ed;
}
.k8s {
color: #3970e4;
}
.go {
color: #007d9c;
}
.react {
color: #61dbfb;
}
.digitalocean {
color: #008bcf;
}
.nginx {
color: #099639;
}
.aws {
color: #FF9900;
}
.azure {
color: #0080FF;
}
.googlecloud {
color: #F4B400;
}


.oneImage {
height: 600px;
width: 800px;
max-height: 500px;
max-width: 800px;
margin: 25px auto 0 auto;
}
.oneImage picture {
max-width: 800px;
}
.oneImage picture img {
width: 100%;
height: 100%;
object-fit: contain;
}
.caption p {
margin: 15px auto 0 auto;
text-align: center;
font-family: var(--slantText);
letter-spacing: .04em;
}
</style>
