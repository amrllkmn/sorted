<script lang="ts">
	import DumpingGround from '$lib/components/ui/dumping-ground';
	import Heading from '$lib/components/ui/heading';
	import Quadrant from '$lib/components/ui/quadrant';
	import TaskForm from '$lib/components/ui/task-form';
	import type { TTask } from '../data.js';

	let { data } = $props();

	let currentTasks = $state(data);

	const addTaskHandler = (description: string) => {
		const newTask: TTask = {
			id: currentTasks.tasks.length + 1,
			description,
			urgent: null,
			important: null
		};

		// ✅ Replace the array reference so Svelte detects the change
		currentTasks = {
			...currentTasks,
			tasks: [...currentTasks.tasks, newTask],
			dumped: [...currentTasks.tasks, newTask].filter((task) => task.important === null && task.urgent === null),
			scheduled: [...currentTasks.tasks, newTask].filter((task) => task.important === false && task.urgent === true),
			delegated: [...currentTasks.tasks, newTask].filter((task) => task.important === true && task.urgent === false),
			eliminated: [...currentTasks.tasks, newTask].filter((task) => task.important === false && task.urgent === false)
		};
	};
</script>

<Heading />
<section class="space-y-8">
	<TaskForm handler={addTaskHandler}/>
	<div class="flex flex-col gap-6 lg:flex-row">
		<div class="lg:w-1/3">
			<DumpingGround items={currentTasks.dumped} />
		</div>
		<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:w-2/3">
			<Quadrant
				title="DO"
				subtitle="Urgent and Important"
				description="Do these immediately"
				items={[]}
			/>
			<Quadrant
				title="SCHEDULE"
				subtitle="Urgent but Not Important"
				description="Find a time to do these"
				items={currentTasks.scheduled}
			/>
			<Quadrant
				title="DELEGATE"
				subtitle="Not Urgent but Important"
				description="If possible, delegate these"
				items={currentTasks.delegated}
			/>
			<Quadrant
				title="ELIMINATE"
				subtitle="Not Urgent and Not Important"
				description="When possible, eliminate these"
				items={currentTasks.eliminated}
			/>
		</div>
	</div>
</section>
