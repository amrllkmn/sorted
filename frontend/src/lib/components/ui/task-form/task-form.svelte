<script lang="ts">
	import { Input } from '$lib/components/ui/input';
	import { Button } from '$lib/components/ui/button';
	import { CirclePlus } from '@lucide/svelte';
	import { enhance } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { error } from '@sveltejs/kit';

	let originalDescription = $state('');

	let { handler } = $props();
</script>

<form
	class="flex w-full flex-col gap-4 sm:flex-row"
	method="POST"
	use:enhance={({ cancel, formData }) => {
		cancel();
		const description = formData.get('description') as string;

		if (!description || description.trim() === '') {
			error(400, {
				message: 'description cannot be empty'
			});
		}
		handler(description);
		originalDescription = '';
	}}
>
	<Input
		placeholder="Dump your todos for the day..."
		name="description"
		bind:value={originalDescription}
	/>
	<Button type="submit">
		<CirclePlus />
		Add task
	</Button>
</form>
