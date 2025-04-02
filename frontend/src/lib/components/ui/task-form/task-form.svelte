<script lang="ts">
	import { Input } from '$lib/components/ui/input';
	import { Button } from '$lib/components/ui/button';
	import { CirclePlus } from '@lucide/svelte';
	import { enhance } from '$app/forms';
  let originalDescription = $state('')

  let { handler } = $props();
</script>

<form
	class="flex w-full flex-col gap-4 sm:flex-row"
	method="POST"
	use:enhance={({ cancel, formData }) => {
		cancel();
		const description = formData.get('description') as string;
		handler(description);
		originalDescription = '';
	}}
>
	<Input placeholder="Dump your todos for the day..." name="description" bind:value={originalDescription} />
	<Button>
		<CirclePlus />
		Add task
	</Button>
</form>
