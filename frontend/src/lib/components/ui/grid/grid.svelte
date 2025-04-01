<script lang="ts">
  let {children} = $props();
</script>
<section class="grid grid-cols-1 md:grid-cols-2 gap-6">
  {@render children()}
</section>