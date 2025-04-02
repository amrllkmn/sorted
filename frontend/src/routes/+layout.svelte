<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>

<main class="container mx-auto px-4 py-10">
	{@render children()}
</main>
