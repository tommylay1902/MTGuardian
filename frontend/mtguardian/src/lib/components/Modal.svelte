<script lang="ts">
  import Form from "./Form.svelte";
  import ActiveModalStore, { updateModal } from "$lib/store/ActiveModalStore";

  export let header: string;
</script>

<div class="modal" class:modal-open={$ActiveModalStore}>
  <div class="modal-box">
    <h3 class="font-bold text-lg">{header}</h3>
    <p class="py-4">
      <Form />
    </p>
    <div class="modal-action">
      <!-- 🔵 set false on click -->
      <button
        class="btn btn-secondary"
        type="button"
        on:click={() => updateModal(false)}
      >
        Cancel
      </button>
    </div>
  </div>
</div>
