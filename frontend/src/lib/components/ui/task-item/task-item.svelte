<script lang="ts">
	import * as Card from '../card';
	import { Button } from '../button';
	import { TrashIcon } from '@lucide/svelte';
	let { description, handleDeleteTask, id } = $props();
</script>

<Card.Root>
	<Card.Content class="flex items-center justify-between px-4 py-2">
		<p class="text-sm">{description}</p>
		<Button size="icon" variant="destructive" onclick={() => handleDeleteTask(id)}>
			<TrashIcon />
		</Button>
	</Card.Content>
</Card.Root>
