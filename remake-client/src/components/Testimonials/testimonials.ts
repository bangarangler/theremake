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
		firstName: 'Jonathan',
		lastName: 'Bernal',
		title: 'Software Engineer',
		company: 'Nowigence Inc.',
		reference:
			"It's hard to find the words for how proud I am to know and to have had the pleasure to work with Jonathan.  He continuously shows his professionalism and ability to get work done in a timely manner while exceeding the goals set for the team.  Placeholder. Placeholder. Placeholder. Placeholder Placeholder Placeholder Placeholder Placeholder Placeholder Placeholder Placeholder Placeholder Placeholder Placeholder Placeholder Placeholder Placeholder Placeholder Placeholder"
	}
];
