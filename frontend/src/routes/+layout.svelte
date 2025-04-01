<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>

<div class="flex min-h-screen w-full flex-col gap-4 p-4">
	{@render children()}
</div>
