<script lang="ts">
	import { Input } from '$lib/components/ui/input';
	import { Button } from '$lib/components/ui/button';
	import { CirclePlus } from '@lucide/svelte';
</script>

<form class="flex w-full flex-col gap-4 sm:flex-row">
	<Input placeholder="Dump your todos for the day..." />
	<Button>
		<CirclePlus />
		Add task
	</Button>
</form>
