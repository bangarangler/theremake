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
}

export const testimonials: Testimonial[] = [
	{
		id: 1,
		firstName: 'Carson',
		lastName: 'Tao',
		title: 'CEO',
		company: 'Nowigence Inc'
	},
	{
		id: 2,
		firstName: 'Ryan',
		lastName: 'Mccarthy',
		title: 'Product Manager',
		company: 'Nowigence Inc'
	},
	{
		id: 3,
		firstName: 'Jonathan',
		lastName: 'Bernal',
		title: 'Software Engineer',
		company: 'Nowigence Inc'
	}
];
