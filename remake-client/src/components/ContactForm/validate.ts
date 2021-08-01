import { toast } from '$stores/toastStore';

export interface ContactFormInput {
	name: string;
	email: string;
	subject: string;
	message: string;
}

export enum ValidOptions {
	NAME = '(Name): I need your name! 📛',
	EMAIL = '(Email): Whats your email? ✉',
	SUBJECT = '(Subject): Whats our conversation about? 🌙',
	MESSAGE = '(Message): Need some details in your message pretty please. ⛵',
	ALL_GOOD = 'ok to submit ✅',
	BAD_NAME = '(Name): Name not long enough. 📛',
	// email handled by good old HTML and ofcourse on the backend
	BAD_SUBJECT = '(Subject): Subject not long enough. 🌙',
	BAD_MESSAGE = '(Message): Message not long enough. ⛵'
}

export interface ValidInfo {
	id: number;
	field: string;
	errType: ValidOptions;
}

export function validate(input: ContactFormInput) {
	const { name, email, subject, message } = input;
	// NAME
	if (!name) {
		return toast.send(ValidOptions.NAME);
	}
	if (name.length <= 2) return toast.send(ValidOptions.BAD_NAME);

	// EMAIL validation handled by html5 and backend
	if (!email) return toast.send(ValidOptions.EMAIL);

	// SUBJECT
	if (!subject) return toast.send(ValidOptions.SUBJECT);
	if (subject.length <= 5) return toast.send(ValidOptions.BAD_SUBJECT);

	// MESSAGE
	if (!message) return toast.send(ValidOptions.MESSAGE);
	if (message.length <= 10) return toast.send(ValidOptions.BAD_MESSAGE);
	return ValidOptions.ALL_GOOD;
}
