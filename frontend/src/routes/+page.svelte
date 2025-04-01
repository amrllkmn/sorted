<script lang="ts">
	import Grid from '$lib/components/ui/grid/grid.svelte';
	import Heading from '$lib/components/ui/heading';
	import Quadrant from '$lib/components/ui/quadrant';
</script>

<Heading />
<Grid>
	<Quadrant title="DO" description="Urgent and Important">
		<div>Hello</div>
	</Quadrant>
	<Quadrant title="SCHEDULE" description="Urgent and Not Important">
		<div>Hello</div>
	</Quadrant>
	<Quadrant title="DELEGATE" description="Not Urgent and Important">
		<div>Hello</div>
	</Quadrant>
	<Quadrant title="ELIMINATE" description="Not Urgent and Not Important">
		<div>Hello</div>
	</Quadrant>
</Grid>
