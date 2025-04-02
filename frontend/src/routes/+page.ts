import { tasks } from "../data"

export const load = () => {
  const dumped = tasks.filter((task) => task.important === null && task.urgent === null)
	const scheduled = tasks.filter((task) => task.important === false && task.urgent === true)

	const delegated = tasks.filter((task) => task.important === true && task.urgent === false)

	const eliminated = tasks.filter((task) => task.important === false && task.urgent === false)
  return { tasks, dumped, scheduled, delegated, eliminated }
}