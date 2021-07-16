const projects = import.meta.globEager('../../projects/*.md');
const projectList = Object.values(projects);
const projectMeta = projectList?.map((p) => {
	return p?.metadata;
});
const slugs = projectMeta.map((s) => {
	return s?.slug;
});
// console.log('slugs', slugs);
// console.log('projectMeta', projectMeta);
export async function get() {
	return { body: slugs };
}
