import { GetAllTasks } from '$lib/wailsjs/go/main/App';

export const load = async () => {
	const tasks = await GetAllTasks();

	const dumped = tasks.filter((task) => task.important === null && task.urgent === null);
	const eliminated = tasks.filter((task) => task.important === false && task.urgent === false);
	const delegated = tasks.filter((task) => task.important === true && task.urgent === false);
	const scheduled = tasks.filter((task) => task.important === false && task.urgent === true);
	const toBeDone = tasks.filter((task) => task.important === true && task.urgent === true);

	return { tasks, dumped, eliminated, delegated, scheduled, toBeDone };
};
