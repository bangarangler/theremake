import { writable } from 'svelte/store';

interface ProjectMetaDataType {
	projectTitle: string;
	slug: string;
	projectRole: string;
	projectDescription: string;
	techUsed: string[];
}

export const projectMetaData = writable<ProjectMetaDataType | undefined[]>([]);
