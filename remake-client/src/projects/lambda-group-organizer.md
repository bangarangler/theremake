---
projectTitle: Group Organizer
slug: lambda-group-organizer
projectRole: Full Stack
projectDescription: 'Winning hackathon project!  Lambda had some challenges with
students signing up for projects.  Lambda group organizer was our solution for
the problem and It won!'
techUsed: ["javascript", "react", "firebase", "netlify", "css", "html5"]
organization: Lambda
year: 2018
---

<script>
  import {onMount} from 'svelte'
  import ProjectInfo from '$components/ProjectInfo/ProjectInfo.svelte'
  import ImageLoader from '$images/ImageLoader.svelte'
  import tempImg from '$static/lambda-group-organizer.png?w=600;800;1600&format=png&srcset'
  import tempImg2 from '$static/lgroupOrgModal.png?w=400;600;800&format=png&srcset'

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
  $: if (innerWidth < 600) headingText = "Software Engineer • Lambda • Hackathon Winner!"
  $: if (innerWidth > 800) headingText = "Software Engineer • Lambda • 2018"
  $: if (innerWidth > 1000) headingText = "Software Engineer • Lambda • Winning Hackathon Project : )"
  $: if (innerWidth < 800) headingText = "Software Engineer • Lambda • Winning Hackathon Project ; )"
  $: if (innerWidth < 600) headingText = "Software Engineer • Lambda";

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
			<ImageLoader
				ty="image/png"
				srcset={tempImg}
				alt="Lambda Group Organizer sign in page"
			/>
    </picture>
</div>

<div class="explanationContainer">

<h2>Whats the problem?</h2>

_Lambda School_ put out a 2.5 day **hackathon challenge** for teams of up to 5 people. I worked on this with two other engineers, <a href="https://jon-bernal.netlify.app" rel="external" target="_blank">Jonathan Bernal</a> and <a href="https://www.linkedin.com/in/shawn-antonucci/">Shawn Antonucci</a>. The problem we decided to solve was the difficulty around signing up for a group project or hackathon. Previously there were issues with students accidentally erasing someone else's name from a group project since everything was done in a google sheet document.

We created _Lambda group organizer_ so a spreadsheet of projects **(csv)** could be loaded into the system each month for the group signups and the system would take care of pulling out group names, descriptions, project proposals, number of students per role, etc. The system would create a new named project session for students to sign up for by project role. Students would have a specific period in which they could sign up and then it would close. The administrator could then download the updated csv of all the data with student's names and roles associated with each project.

The system itself has three roles that are allowed, students, teachers assistants and admin. The students could only sign up for a project. The teacher's assistants could help with moving students from one group to another or adjusting group requirements as needed. The admin could of course do anything.

After we finished the hackathon we were asked to actually help them put the **project into action** for the school because they intended to use it for student group signups! Everyone from the group got hired to companies while working on getting it production ready, but we trained our replacements before we moved on. I hope _Lambda_ students are getting to enjoy working on it, it was a blast to make!

_Side note_: <a href="https://jon-bernal.netlify.app" rel="external" target="_blank">Jonathan Bernal</a> and <a href="https://linkedin.com/in/shawn-antonucci/">Shawn Antonucci</a> were amazing to work with. I couldn't recommend them as team members more strongly. In fact **Jon** and I work together currently at <a href="https://nowigence.com" target="_blank" rel="external">Nowigence Inc!</a>

<h3 class:slideInLeft={animations[0].isVis} class="domEle0">Backend</h3>

This one is super simple. We put it on _netlify_ and setup a simple **CI/CD** with github main branch. The database was using _Google Firebase_ for speed and ease of setup when it came to the hackathon, however it has since been ported to <span class="mongo">MongoDB</span> last we where involved while optimizing it for production use.

<h3 class:slideInRight={animations[1].isVis} class="domEle1">Frontend</h3>

We used <span class="react">React</span>, <span class="javascript">Javascript</span>, <span class="css">CSS</span> and Semantic UI. We needed the speed of a UI Library to quickly put something together that looked decent but was still somewhat modifiable. If I were going to work on this for years I would strip Semantic UI from the project to be honest. It can make for some messy css when you have to overide styles in a lot of places. However for a hackathon we didn't want to waste to much time on styling.

<h3 class:slideInLeft={animations[2].isVis} class="domEle2">Dev Ops</h3>

Not a whole lot to say here, pretty much everything was a managed service since we were going for speed. That wouldn't be the case if we were going to hang on to this for a long time for cost reasons.

</div>

<div class="card oneImage">
  <picture>
  <source media="(min-width:1200px)" srcset={tempImg2}>
    <ImageLoader srcset={tempImg2} ty="image/png" alt="Lambda group organizer
    Modal view of a project with someone nammed Jimmy already signed up." />
  </picture>
</div>
<div class="caption"><p>Quick detail view for a project!</p></div>

<div class="linkWrap"><a class="draw-outline draw-outline--tandem" href="https://lambda-group-organizer.firebaseapp.com/" rel="external" target="_blank">Lambda Group Organizer</a></div>

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
.caption p {
margin: 15px auto 0 auto;
text-align: center;
font-family: var(--slantText);
letter-spacing: .04em;
}
.linkWrap {
  width: 100%;
  margin: 50px auto;
display: flex;
justify-content: center;
}
.draw-outline {
  display: inline-block;
  padding: 16px 28px;
  border: 2px var(--dracPurp) solid;
  text-align: center;
  text-decoration: none;
  color: var(--dracPurp);
  position: relative;
  transition: border-color 0.35s ease-in-out;
  z-index: 1;
}
.draw-outline:before, .draw-outline:after {
  content: '';
  position: absolute;
  width: 0px;
  height: 0px;
  box-sizing: content-box;
  z-index: -1;
  transition: transform 0.25s ease-in-out;
  padding-left: 2px;
}
.draw-outline:before {
  top: -2px;
  left: -2px;
  border-top: 2px transparent solid;
  border-right: 2px transparent solid;
}
.draw-outline:after {
  bottom: -2px;
  right: -2px;
  border-bottom: 2px transparent solid;
  border-left: 2px transparent solid;
}
.draw-outline:hover {
  color: var(--hotpink);
  transition: color 0.35s ease-in-out, border-color 0.35s ease-in-out;
  border-color: var(--hotpink);
  animation: outline-reset 0.35s 1 forwards;
}
.draw-outline:hover:before {
  animation: top-right-border 0.75s 1 0.35s forwards;
}
.draw-outline:hover:after {
  animation: bottom-left-border 0.75s 1 1.1s forwards;
}
.draw-outline--tandem:hover:after {
  animation: bottom-left-border 0.75s 1 0.35s forwards;
}
.draw-outline:active:before, .draw-outline:active:after {
  transform: scale(1.05);
}
@keyframes outline-reset {
  0% {
    border-color: var(--hotpink);
  }
  100% {
    border-color: transparent;
  }
}
@keyframes top-right-border {
  0% {
    border-color: var(--hotpink);
    width: 0px;
    height: 0;
  }
  50% {
    width: 100%;
    height: 0;
  }
  100% {
    border-color: var(--hotpink);
    width: 100%;
    height: 100%;
  }
}
@keyframes bottom-left-border {
  0% {
    border-color: var(--hotpink);
    width: 0px;
    height: 0;
  }
  50% {
    width: 100%;
    height: 0;
  }
  100% {
    border-color: var(--hotpink);
    width: 100%;
    height: 100%;
  }
}
.react {
color: #61dbfb;
}
.mongo {
color: #4DB33D;
}
.javascript {
color: #fcdc00;
}
.css {
color: #1a0dab;
}
a:hover {
  color: var(--dracPurp);
  border-bottom: 2px solid var(--dracPurp);
}
</style>
