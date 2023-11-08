<script lang="ts">
  import Form from "./Form.svelte";
  import ActiveModalStore, { updateModal } from "$lib/store/ActiveModalStore";

  export let header;
</script>

<div class="modal" class:modal-open={$ActiveModalStore}>
  <div class="modal-box">
    <h3 class="font-bold text-lg">{header}</h3>
    <p class="py-4">
      <Form />
    </p>
    <div class="modal-action">
      <!-- 🔵 set false on click -->
      <button class="btn" on:click={() => updateModal(false)}>Yay!</button>
    </div>
  </div>
</div>
