const projects = import.meta.globEager('../../projects/*.md');
const projectList = Object.values(projects);
const projectMeta = projectList?.map((p) => {
	return p?.metadata;
});
// console.log('projectMeta', projectMeta);
export async function get() {
	return { body: projectMeta };
}
