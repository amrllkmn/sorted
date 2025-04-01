<script lang="ts">
	let { children } = $props();
</script>

<section class="grid grid-cols-1 gap-6 md:grid-cols-2">
	{@render children()}
</section>
