<script lang="ts">
	import * as Card from '../card';
	let { title, description = '', subtitle, items, deleteTaskHandler, updateTaskHandler } = $props();
	import { dndzone } from 'svelte-dnd-action';
	import TaskItem from '../task-item/task-item.svelte';
	import { flip } from 'svelte/animate';

	const flipDurationMs = 300;

	const handleDndAction = (e: { detail: { items: any } }) => {
		items = e.detail.items;
	};

	const handleUpdateTask = (e: { detail: { items: any } }) => {
		items = e.detail.items;
		items.forEach((item: { id: any }) => {
			updateTaskHandler(item.id);
		});
	};
</script>

<Card.Root class="flex h-[300px] flex-col overflow-hidden">
	<Card.Header class="flex-shrink-0 pb-2">
		<Card.Title>{title}</Card.Title>
		<Card.Description class="space-y-2">
			<p class="text-sm font-medium">{subtitle}</p>
			<p class="text-xs">{description}</p>
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
			onfinalize={handleUpdateTask}
			class="flex h-full flex-col space-y-2 overflow-y-auto"
		>
			{#each items as item (item.id)}
				<div animate:flip={{ duration: flipDurationMs }}>
					<TaskItem
						description={item.description}
						handleDeleteTask={deleteTaskHandler}
						id={item.id}
					/>
				</div>
			{/each}
		</div>
	</Card.Content>
</Card.Root>
