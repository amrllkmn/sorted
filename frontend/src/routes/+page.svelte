<script lang="ts">
	import DumpingGround from '$lib/components/ui/dumping-ground';
	import Heading from '$lib/components/ui/heading';
	import Quadrant from '$lib/components/ui/quadrant';
	import TaskForm from '$lib/components/ui/task-form';

	let {data} = $props();

	let currentTasks = $state(data.tasks)

	let items1 = $derived.by(() => currentTasks.filter((task) => task.important === null && task.urgent === null))
	let items2 = $derived.by(() => currentTasks.filter((task) => task.important === false && task.urgent === true))

	let items3 = $derived.by(() => currentTasks.filter((task) => task.important === true && task.urgent === false))

	let items4 = $derived.by(() => currentTasks.filter((task) => task.important === false && task.urgent === false))
</script>

<Heading />
<section class="space-y-8">
	<TaskForm />
	<div class="flex flex-col gap-6 lg:flex-row">
		<div class="lg:w-1/3">
			<DumpingGround items={items1}/>
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
				items={items2}
			/>
			<Quadrant
				title="DELEGATE"
				subtitle="Not Urgent but Important"
				description="If possible, delegate these"
				items={items3}
			/>
			<Quadrant
				title="ELIMINATE"
				subtitle="Not Urgent and Not Important"
				description="When possible, eliminate these"
				items={items4}
			/>
		</div>
	</div>
</section>
