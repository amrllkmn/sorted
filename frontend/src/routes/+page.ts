import { GetAllTasks } from '$lib/wailsjs/go/main/App';

export const load = async () => {
	const tasks = await GetAllTasks();

	console.log(tasks);

	return { tasks };
};
