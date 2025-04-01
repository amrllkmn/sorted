<script lang="ts">
	import { Input } from '../input';
	import { Button } from '../button';
	import { CirclePlus } from '@lucide/svelte';
</script>

<section class="mx-auto flex w-full flex-col items-center justify-center">
	<h1 class="m-4 text-4xl font-bold md:text-6xl">Sorted.</h1>
	<p class="m-4 text-lg text-muted-foreground"><i>Dump it. Do it. Sorted.</i></p>
	<form class="m-4 flex w-full flex-row items-center justify-center gap-2">
		<Input placeholder="Dump your todos for the day..." />
		<Button>
			<CirclePlus />
			Add task
		</Button>
	</form>
</section>
