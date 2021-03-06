import { writable } from 'svelte/store';

export interface ProjectMetaDataType {
	projectTitle: string;
	slug: string;
	projectRole: string;
	projectDescription: string;
	techUsed: string[];
}

export interface PreviousAndNextProjectType {
	previous?: string;
	next?: string;
}

// export enum ToolChoices {
// 	'JAVASCRIPT' = 'javascript',
// 	'TYPESCRIPT' = 'typescript'
// }

// const tools: ToolChoices[] = [ToolChoices.JAVASCRIPT, ToolChoices.TYPESCRIPT];

export const projectMetaData = writable<ProjectMetaDataType[]>([]);

export const previousAndNextProject = writable<PreviousAndNextProjectType>({});

// export const languageIcons = writable<ToolChoices[]>(tools);
