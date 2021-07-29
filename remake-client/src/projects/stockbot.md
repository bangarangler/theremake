---
projectTitle: Stock Bot
slug: stockbot
projectRole: Full Stack
projectDescription: 'Stock Bot is an trading robot.  Designed with the sole
purpose of making us money while we sleep.  I need to learn this. | Note to
future self... jon go study this!'
techUsed: ["typescript", "golang", "aws", "docker", "kubernetes"]
organization: J&J Studios
year: 2018 - Always
---

<script>
  import {onMount} from 'svelte'
  import ProjectInfo from '$components/ProjectInfo/ProjectInfo.svelte'
  //import tempImg from '$static/coming-soon-pixabay.jpg?w=600;700;1600&format=jpg&srcset'
  //import tempImg from '$static/the-new-beginning-pixabay.jpg?w=200;400;700&format=jpg&srcset'
  import tempImg from '$static/pexels-working.jpg?w=600;800;1600&format=jpg&srcset'
  import tempImg2 from '$static/pexels-phone-art.jpg?w=400;600;800&format=jpg&srcset'

  let headingText = "Software Engineer • Ongoing work in progress ; )"
  let innerWidth;
	let animation = false;
  let animations = [
  {domElement: null, isVis: false},
  {domElement: null, isVis: false},
  {domElement: null, isVis: false}
  ]
    onMount(() => {
  animation = true
  })
  $: if (innerWidth < 600) headingText = "Software Engineer • Partner • We do it all!"
  $: if (innerWidth > 800) headingText = "Software Engineer • Partner • 2019 - Present"
  $: if (innerWidth > 1000) headingText = "Software Engineer • Ongoing work in progress : )"
  $: if (innerWidth < 800) headingText = "Software Engineer • Ongoing work in progress ; )"
  $: if (innerWidth < 600) headingText = "Software Engineer •  We do it all!";

function isInViewport(element ) {
if (element) {
const rect = element.getBoundingClientRect();
    return (
        rect.top >= 0 &&
        rect.left >= 0 &&
        rect.bottom <= (window.innerHeight || document.documentElement.clientHeight) &&
        rect.right <= (window.innerWidth || document.documentElement.clientWidth)
    );
}
}

function animationEvent() {
  animations = animations.map((a, idx) => {
  return {domElement:  document.querySelector(`.domEle${idx}`),
  isVis:  isInViewport(a.domElement)}
  })
}

onMount(() => {
document.addEventListener('scroll', animationEvent)
return () => document.removeEventListener('scroll', animationEvent)
})

$: isInViewport(animations[0].domElement)
$: isInViewport(animations[1].domElement)
$: isInViewport(animations[2].domElement)
</script>

<svelte:window bind:innerWidth />

<article>
<div class="container">
<h1 class:display={animation}>{projectTitle}</h1>

<p class="headingText">{headingText}</p>
</div>

  <div class="card imgContainer">
    <picture>
      <source media="(min-width:1000px)" srcset={tempImg}>
        <img class="img1" srcset={tempImg} type="image/jpg" alt="Coming Soon!" />
    </picture>
    <!--<picture>
      <source media="(min-width:1200px)" srcset={tempImg1}>
      <img class="img2" srcset={tempImg1} type="image/jpg" alt="new beginning
      and working like crazy">
    </picture>-->
</div>

<div class="explanationContainer">

<h2>Whats the holdup?</h2>

_WORK ~ LIFE_ is the current hang. I work like a **crazy** person. 13 - 16 hour
days are not uncommon and I've been like that since I can remember. _"Workaholic"_ would be a fair assesment. In the midst of very long workdays; Personal work gets slowed down.

Reachinng out to all my **developers...** you guys know how it goes ; ) In
fairness life at a start up is nuts. Lost count of the hats that we all wear <span class="emoji">🎩</span>.

<h3 class:slideInLeft={animations[0].isVis} class="domEle0">Backend</h3>

This is our personal project! We have, are and will continue to use the best technology
for the job. We like to keep up to date and utilize new tech. If it's
performant and can scale we'll give it a go and see if we like it!

<h3 class:slideInRight={animations[1].isVis} class="domEle1">Frontend</h3>

**Jon** put in content here!

_Viewer_ See note for **Jon** above ; )

<h3 class:slideInLeft={animations[2].isVis} class="domEle2">Dev Ops</h3>

_Note to self..._ This needs content as well!

</div>

<div class="card oneImage">
  <picture>
  <source media="(min-width:1200px)" srcset={tempImg2}>
    <img srcset={tempImg2} type="image/jpg" alt="phone... man time flys">
  </picture>
</div>
<div class="caption"><p>Time keeps flying!</p></div>

<ProjectInfo {projectRole} {techUsed} {organization} {year}/>

</article>

<style>
article {
  margin: 40px 25px 0;
}
@media (min-width: 460px) {
article {
  margin: 0 25px 0;
}
}
@media (min-width: 500px) {
article {
  margin: 0 45px;
}
}
.container {
margin-bottom: 45px;
}
@media (min-width: 500px) {
.container {
margin-bottom: 65px;
}
}
h1 {
background: linear-gradient(271deg,var(--hotpink) 30%, 50%,var(--aqua) 70%,#a162e8 94%);
background-clip: border-box;
-webkit-background-clip: text;
-webkit-text-fill-color: transparent;
opacity: 0;
font-size: var(--h2);
}
@media (min-width: 500px) {
h1 {
font-size: var(--h1);
}
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
grid-template-columns: minmax(0, 1fr);
place-items: center;
}
@media (min-width: 600px) {
.imgContainer {
  grid-template-columns: minmax(0, 1fr);
}
}
@media (min-width: 1000px) {
.imgContainer {
  grid-template-columns: minmax(0, 1fr);
  width: 100%;
  place-items: center;
  max-width: 1000px;
  margin: 0 auto 25px;
  }
}

.imgContainer picture {
  width: auto;
}
.imgContainer picture:nth-child(2) {
display: none;
}
@media (min-width: 600px) {
.imgContainer picture:nth-child(2) {
display: block;
}
}
@media (min-width: 800px) {
.imgContainer picture {
  width: auto;
  max-width: 800px;
}
}
@media(min-width: 1000px) {
.imgContainer picture {
  height: unset;
  width: auto;
  max-width: 1000px;
}
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
  width: 90%;
  margin: 0 auto;
}
@media (min-width: 500px) {
.explanationContainer {
  width: 75%;
}
}
@media (min-width: 1000px) {
.explanationContainer {
  width: 50%;
  max-width: 800px;
  display: flex;
  flex-direction: column;
  align-items: center;
}
}
.explanationContainer h3 {
opacity: 0;
width: max-content;
}

@media (min-width: 1000px) {
  .explanationContainer h3 {
  opacity: 0;
  width: max-content;
  align-self: flex-start;
  }
}

.explanationContainer h2 {
text-align: left;
font-size: var(--h3);
}
@media (min-width: 1000px) {
  .explanationContainer h2 {
  text-align: left;
  font-size: var(--h3);
  align-self: flex-start;
  }
}

@media(min-width: 1000px) {
  .explanationContainer h2 {
    font-size: var(--h2);
  }
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
color: var(--darkAquaLightHotPink);
}
.emoji {
    background: var(--lightGray);
    border-radius: 12px;
    padding: 2px;
}

.oneImage {
margin: 25px auto 0 auto;
}
@media (min-width: 1000px) {
.oneImage {
height: 600px;
width: 800px;
max-height: 500px;
max-width: 800px;
}
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
