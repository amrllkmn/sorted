<script lang="ts">
	let { children } = $props();
</script>

<section class="grid h-full w-full grid-cols-1 gap-6 md:grid-cols-2">
	{@render children()}
</section>
