const projects = import.meta.globEager('../../projects/*.md');
const projectList = Object.values(projects);
const projectMeta = projectList?.map((p) => {
	return p?.metadata;
});
// console.log('projectMeta', projectMeta);
//https://ar.al/2021/04/03/passing-data-from-layouts-to-pages-in-sveltekit/
export async function get() {
	return { body: projectMeta };
}
