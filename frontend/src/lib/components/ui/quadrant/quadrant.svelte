<script lang="ts">
	import * as Card from '../card';
	let { title, description = '', subtitle, items } = $props();
	import { dndzone } from 'svelte-dnd-action';
	import TaskItem from '../task-item/task-item.svelte';
	import { flip } from 'svelte/animate';

	const flipDurationMs = 300;

	const handleDndAction = (e: { detail: { items: any } }) => {
		items = e.detail.items;
	};
</script>

<Card.Root class="h-full">
	<Card.Header class="pb-2">
		<Card.Title>{title}</Card.Title>
		<Card.Description class="space-y-2">
			<p class="text-sm font-medium">{subtitle}</p>
			<p class="text-xs">{description}</p>
		</Card.Description>
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
			class="flex max-h-[200px] min-h-[100px] flex-col space-y-2 overflow-y-auto"
		>
			{#each items as item (item.id)}
				<div animate:flip={{ duration: flipDurationMs }}>
					<TaskItem description={item.description} />
				</div>
			{/each}
		</div>
	</Card.Content>
</Card.Root>
