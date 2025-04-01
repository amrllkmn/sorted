<script lang="ts">
  import { Input } from "../input";
  import { Button } from "../button";
  import { CirclePlus } from "@lucide/svelte";
</script>

<section class="mx-auto flex flex-col items-center justify-center w-full">
  <h1 class="text-4xl font-bold md:text-6xl m-4">Sorted.</h1>
  <p class="text-muted-foreground text-lg m-4"><i>Dump it. Do it. Sorted.</i></p>
  <form class="flex flex-row items-center justify-center gap-2 m-4 w-full">
      <Input placeholder="Dump your todos for the day..."/>
      <Button>
        <CirclePlus/>
        Add task
        </Button>
  </form>
</section>