<script lang="ts">
	import Grid from '$lib/components/ui/grid/grid.svelte';
	import Heading from '$lib/components/ui/heading';
	import Quadrant from '$lib/components/ui/quadrant';

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
<section class="grid h-screen w-full grid-cols-[1fr_2fr] gap-6">
	<Quadrant title="DUMP" description="All the stuff in your head goes here" items={items1} />
	<Grid>
		<Quadrant title="DO" description="Urgent and Important" items={[]} />
		<Quadrant title="SCHEDULE" description="Urgent but Not Important" items={items2} />
		<Quadrant title="DELEGATE" description="Not Urgent but Important" items={items3} />
		<Quadrant title="ELIMINATE" description="Not Urgent and Not Important" items={items4} />
	</Grid>
</section>
