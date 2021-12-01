export interface Testimonial {
	id: number;
	firstName: string;
	lastName: string;
	title: string;
	company: string;
	twitter?: string;
	facebook?: string;
	linkedin?: string;
	instagram?: string;
	website?: string;
	reference: string;
}

export const testimonials: Testimonial[] = [
	{
		id: 1,
		firstName: 'Carson',
		lastName: 'Tao',
		title: 'Lead AI Scientist',
		company: 'Sumitovant Biopharma, Inc.',
		reference:
			'It has been a great pleasure to work with Jon for the past 8 months. I was directly involved in the day-to-day interaction with him when I was the CTO of Nowigence. Jon is a diligent developer with excellent project management, communication, and technical skills. He is also attentive to all sorts of technical issues and challenges we encountered during product development and maintenance. I felt much more confident when Jon is working on my team! Keep up the great work!'
	},
	{
		id: 2,
		firstName: 'Ryan',
		lastName: 'McCarthy',
		title: 'Product Manager',
		company: 'Energage',
		reference:
			'Jon\'s leadership at Nowigence was unparalleled in my time there.  Jon exemplifies the mindset of "to lead by example".  It was a pleasure working so closely with him for a year and a half! He will outwork every person in the room period. Jon had various responsibilities throughout his time at Nowigence and handled all additions to his day to day operation with ease.  I genuinely hope I have the opportunity to work with him again in the future.'
	},
	{
		id: 3,
		firstName: 'Ilya',
		lastName: 'Mayzus',
		title: 'Senior Software Engineer',
		company: 'Lucidus Solutions, LLC',
		reference:
			'I enjoyed greatly working with Jon P for over a year. While we handled different stacks of the application (me on Python backend, him on Node/JS), we collaborated on design of new projects that required delving into API (such as Microsoft OneDrive integration or GoogleDrive integration). What struck me was his fearless and optimistic attitude to tackling new problems, which is absolutely necessary when working on new things, especially in a start-up. I particularly valued his support and leadership, whether in managing a team of outside developers or researching ways to improve security of the application, or exploring new features of the next release of MongoDB. Perhaps more than his technical skills, I appreciated his people skills. Jon would be a great asset to any team he would join in the future. '
	},
	{
		id: 4,
		firstName: 'Jonathan',
		lastName: 'Bernal',
		title: 'Software Engineer',
		company: 'Nowigence Inc.',
		reference:
			"It's hard to find the words for how proud I am to know and to have had the pleasure to work with Jon.  He continuously shows his professionalism and ability to get work done in a timely manner while exceeding the goals set before not just himself, but his team as well.  He is constantly jumping in and putting out fires and not just fixing the problem (which is obviously important) but slowing down to walk you through it and make sure that you understand, and that you have a good grasp of how and why we needed to do something. I would work with Jon on any project and I certainly wouldn't hesitate to hire him. There are not enough good candidates as it is. I have zero doubt that Jonathan will be at the top of any team and bringing the entire organization up to such a high level of excellence right there with him. The bar for fantastic people to work with is ridiculously high. Jon, in my opinion, is very easily above and beyond that bar and will be a wonderful asset to any company or project that he is hired for."
	},
	{
		id: 5,
		firstName: 'Shawn',
		lastName: 'Antonucci',
		title: 'Software Developer',
		company: 'Nexient',
		reference:
			"Jon worked directly under me and was a phenominal teammate. We spun up an entire react native project to accompany a web project during a hackathon and had a fantastic time developing it.  I can say that he learns quickly and has a fantastic attitude even while facing new or challenging problems.  He works like a mad man!  I've never seen such a strong work ethic along with fantastic communication skills.  I look forward to working with him again in the future and feel confident in saying he will be a tremendous asset to any team!"
	},
	{
		id: 6,
		firstName: 'Shridatta',
		lastName: 'Zanjare',
		title: 'Frontend Software Developer',
		company: 'Lowes',
		reference:
			'I had a great time working with Jon at Nowigence. Jon is a great leader with great problem solving skills. I love his work ethics and he is always willing help and solve complex problems with great enthusiasm. He is very passionate about coding and is a great team mate. Not just that but he is great at managing people and is definitely a jack of all trades. I would have Jon on my team any day and would love to work with him in future if it happens.'
	}
];
