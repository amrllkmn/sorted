<script lang="ts">
	import * as Card from '../card';
	let { title, description, items } = $props();
	import { dndzone } from 'svelte-dnd-action';
	import TaskItem from '../task-item/task-item.svelte';
	import { flip } from 'svelte/animate';

	const flipDurationMs = 300;

	const handleDndAction = (e: { detail: { items: any } }) => {
		items = e.detail.items;
	};
</script>

<Card.Root class="h-full w-full">
	<Card.Header>
		<Card.Title>{title}</Card.Title>
		<Card.Description>{description}</Card.Description>
	</Card.Header>
	<Card.Content>
		<div
			use:dndzone={{
				items: items,
				flipDurationMs: flipDurationMs,
				dropTargetStyle: {
					outline: 'none'
				},
				dropTargetClasses: ['bg-secondary', 'rounded', 'border-xl']
			}}
			onconsider={handleDndAction}
			onfinalize={handleDndAction}
			class="flex max-h-[200px] min-h-[100px] flex-col gap-2 overflow-y-auto"
		>
			{#each items as item (item.id)}
				<div animate:flip={{ duration: flipDurationMs }}>
					<TaskItem description={item.description} />
				</div>
			{/each}
		</div>
	</Card.Content>
</Card.Root>
