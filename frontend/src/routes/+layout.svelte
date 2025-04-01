<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>
<div class="flex min-h-screen w-full flex-col p-4 gap-4">
	{@render children()}
</div>
