import { writable } from 'svelte/store';

export interface ProjectMetaDataType {
	projectTitle: string;
	slug: string;
	projectRole: string;
	projectDescription: string;
	techUsed: string[];
}

export const projectMetaData = writable<ProjectMetaDataType[]>([]);

export const projectSlugs = writable<string[]>([]);
