<script lang="ts">
	import * as Card from '../card';
	let { items } = $props();
	import { dndzone } from 'svelte-dnd-action';
	import TaskItem from '../task-item/task-item.svelte';
	import { flip } from 'svelte/animate';

	const flipDurationMs = 300;

	const handleDndAction = (e: { detail: { items: any } }) => {
		items = e.detail.items;
	};
</script>

<Card.Root class="h-[600px] flex flex-col overflow-hidden">
	<Card.Header class="pb-2 flex-shrink-0 flex flex-row items-center justify-between">
		<Card.Title>DUMP</Card.Title>
		<Card.Description class="space-y-2">
			<p class="text-sm font-medium">All the tasks in your head go here</p>
		</Card.Description>
	</Card.Header>
	<Card.Content class="flex-1 overflow-hidden">
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
			class="flex flex-col space-y-2 h-full overflow-y-auto"
		>
			{#each items as item (item.id)}
				<div animate:flip={{ duration: flipDurationMs }}>
					<TaskItem description={item.description} />
				</div>
			{/each}
		</div>
	</Card.Content>
</Card.Root>
