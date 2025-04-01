<script lang="ts">
	import * as Card from '../card';
	import { Button } from '../button';
	import { TrashIcon } from '@lucide/svelte';
	let { description } = $props();
</script>

<Card.Root>
	<Card.Content>
		<div class="flex flex-row items-center justify-between">
			<p>{description}</p>
			<Button size="icon" variant="destructive">
				<TrashIcon />
			</Button>
		</div>
	</Card.Content>
</Card.Root>
