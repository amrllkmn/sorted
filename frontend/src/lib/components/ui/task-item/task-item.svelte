<script lang="ts">
	import * as Card from '../card';
	import { Button } from '../button';
	import { TrashIcon } from '@lucide/svelte';
	let { description } = $props();
</script>

<Card.Root>
	<Card.Content class="flex items-center justify-between px-4 py-2">
		<p class="text-sm">{description}</p>
		<Button size="icon" variant="destructive">
			<TrashIcon />
		</Button>
	</Card.Content>
</Card.Root>
