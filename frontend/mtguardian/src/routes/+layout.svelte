<script>
  import Menu from "$lib/components/Menu.svelte";
  import "../app.postcss";

  import { Toaster } from "svelte-french-toast";
</script>

<Toaster />
<Menu />
<slot />
