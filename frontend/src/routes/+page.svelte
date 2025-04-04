<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import DumpingGround from '$lib/components/ui/dumping-ground';
	import Heading from '$lib/components/ui/heading';
	import Quadrant from '$lib/components/ui/quadrant';
	import TaskForm from '$lib/components/ui/task-form';
	import { CreateTask, DeleteTask } from '$lib/wailsjs/go/main/App.js';

	let { data } = $props();

	let currentTasks = $derived(data.tasks);
	let scheduledTasks = $derived(
		currentTasks.filter((task) => task.important === false && task.urgent === true)
	);
	let delegatedTasks = $derived(
		currentTasks.filter((task) => task.important === true && task.urgent === false)
	);
	let dumpedTasks = $derived(
		currentTasks.filter((task) => task.important === null && task.urgent === null)
	);
	let eliminatedTasks = $derived(
		currentTasks.filter((task) => task.important === false && task.urgent === false)
	);

	const addTaskHandler = async (description: string) => {
		await CreateTask(description);

		await invalidateAll();
	};

	const deleteTaskHandler = async (id: number) => {
		await DeleteTask(id);

		await invalidateAll();
	};

	$effect(() => {
		console.log('task list updated', data.tasks.length);
	});
</script>

<Heading />
<section class="space-y-8">
	<TaskForm handler={addTaskHandler} />
	<div class="flex flex-col gap-6 lg:flex-row">
		<div class="lg:w-1/3">
			<DumpingGround items={dumpedTasks} {deleteTaskHandler} />
		</div>
		<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:w-2/3">
			<Quadrant
				title="DO"
				subtitle="Urgent and Important"
				description="Do these immediately"
				items={[]}
				{deleteTaskHandler}
			/>
			<Quadrant
				title="SCHEDULE"
				subtitle="Urgent but Not Important"
				description="Find a time to do these"
				items={scheduledTasks}
				{deleteTaskHandler}
			/>
			<Quadrant
				title="DELEGATE"
				subtitle="Not Urgent but Important"
				description="If possible, delegate these"
				items={delegatedTasks}
				{deleteTaskHandler}
			/>
			<Quadrant
				title="ELIMINATE"
				subtitle="Not Urgent and Not Important"
				description="When possible, eliminate these"
				items={eliminatedTasks}
				{deleteTaskHandler}
			/>
		</div>
	</div>
</section>
