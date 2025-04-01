<script lang="ts">
	import * as Card from '../card';
	let { title, description, children, color } = $props();
</script>

<Card.Root class="w-full">
	<Card.Header>
		<Card.Title>{title}</Card.Title>
		<Card.Description>{description}</Card.Description>
	</Card.Header>
	<Card.Content>
		{@render children()}
	</Card.Content>
</Card.Root>
