<script lang="ts">
	import DumpingGround from '$lib/components/ui/dumping-ground';
	import Heading from '$lib/components/ui/heading';
	import Quadrant from '$lib/components/ui/quadrant';
	import TaskForm from '$lib/components/ui/task-form';

	let items1 = [
		{ id: 1, description: 'item1' },
		{ id: 2, description: 'item2' },
		{ id: 3, description: 'item3' },
		{ id: 4, description: 'item4' }
	];

	let items2 = [
		{ id: 5, description: 'item5' },
		{ id: 6, description: 'item6' },
		{ id: 7, description: 'item7' },
		{ id: 8, description: 'item8' }
	];

	let items3 = [
		{ id: 9, description: 'item9' },
		{ id: 10, description: 'item10' },
		{ id: 11, description: 'item11' },
		{ id: 12, description: 'item12' }
	];

	let items4 = [
		{ id: 13, description: 'item13' },
		{ id: 14, description: 'item14' },
		{ id: 15, description: 'item15' },
		{ id: 16, description: 'item16' }
	];
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
