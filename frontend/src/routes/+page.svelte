<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import DumpingGround from '$lib/components/ui/dumping-ground';
	import Heading from '$lib/components/ui/heading';
	import Quadrant from '$lib/components/ui/quadrant';
	import TaskForm from '$lib/components/ui/task-form';
	import { CreateTask, DeleteTask, UpdateTaskStatus } from '$lib/wailsjs/go/main/App.js';

	let { data } = $props();

	let scheduledTasks = $derived(data.scheduled);
	let delegatedTasks = $derived(data.delegated);
	let dumpedTasks = $derived(data.dumped);
	let eliminatedTasks = $derived(data.eliminated);
	let tasksToBeDone = $derived(data.toBeDone);

	const addTaskHandler = async (description: string) => {
		await CreateTask(description);

		await invalidateAll();
	};

	const deleteTaskHandler = async (id: number) => {
		await DeleteTask(id);

		await invalidateAll();
	};

	const handleTaskToBeDone = async (id: number) => {
		await UpdateTaskStatus(id, true, true);
		await invalidateAll();
	};

	const handleTaskToBeScheduled = async (id: number) => {
		await UpdateTaskStatus(id, true, false);
		await invalidateAll();
	};
	const handleTaskToBeDelegated = async (id: number) => {
		await UpdateTaskStatus(id, false, true);
		await invalidateAll();
	};
	const handleTaskToBeEliminated = async (id: number) => {
		await UpdateTaskStatus(id, false, false);
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
		<!-- on drag finalize set urgent: true important: true-->
		<div class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:w-2/3">
			<Quadrant
				title="DO"
				subtitle="Urgent and Important"
				description="Do these immediately"
				items={tasksToBeDone}
				{deleteTaskHandler}
				updateTaskHandler={handleTaskToBeDone}
			/>
			<!-- on drag finalize set urgent: true important: false -->
			<Quadrant
				title="SCHEDULE"
				subtitle="Urgent but Not Important"
				description="Find a time to do these"
				items={scheduledTasks}
				{deleteTaskHandler}
				updateTaskHandler={handleTaskToBeScheduled}
			/>
			<!-- on drag finalize set urgent: false important: true -->
			<Quadrant
				title="DELEGATE"
				subtitle="Not Urgent but Important"
				description="If possible, delegate these"
				items={delegatedTasks}
				{deleteTaskHandler}
				updateTaskHandler={handleTaskToBeDelegated}
			/>
			<!-- on drag finalize set urgent: false important: false -->
			<Quadrant
				title="ELIMINATE"
				subtitle="Not Urgent and Not Important"
				description="When possible, eliminate these"
				items={eliminatedTasks}
				{deleteTaskHandler}
				updateTaskHandler={handleTaskToBeEliminated}
			/>
		</div>
	</div>
</section>
