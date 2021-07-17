import { writable } from 'svelte/store';

export interface ProjectMetaDataType {
	projectTitle: string;
	slug: string;
	projectRole: string;
	projectDescription: string;
	techUsed: string[];
}

export interface PreviousAndNextProject {
	previous?: string;
	next?: string;
}

export const projectMetaData = writable<ProjectMetaDataType[]>([]);

export const previousAndNextProject = writable<PreviousAndNextProject>({});
