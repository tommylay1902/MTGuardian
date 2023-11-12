<script lang="ts">
  import Form from "./Form.svelte";
  import ActiveModalStore, {
    resetModalStore,
  } from "$lib/store/ActiveModalStore";
  import PrescriptionStore from "$lib/store/PrescriptionStore";
  import { resetFormStore } from "$lib/store/Form";
  async function deletePrescription() {
    PrescriptionStore.update((currentData) => {
      currentData = currentData.filter(
        (curr) => curr.id !== $ActiveModalStore.id
      );
      return currentData;
    });

    await fetch(
      `http://0.0.0.0:8000/api/v1/prescription/${$ActiveModalStore.id}`,
      {
        method: "DELETE",
      }
    );
    resetFormStore();
    resetModalStore();
  }
</script>

<div class="modal" class:modal-open={$ActiveModalStore.isOpen}>
  <div class="modal-box">
    <h3 class="font-bold text-lg">{$ActiveModalStore.header}</h3>
    <p class="py-4">
      {#if $ActiveModalStore.body !== "form" && $ActiveModalStore.isOpen}
        <p>{$ActiveModalStore.body}</p>
        <div class="flex flex-row space-x-3 pt-4">
          <button class="btn btn-primary w-1/2" on:click={deletePrescription}
            >Delete</button
          >
          <button class="btn btn-secondary w-1/2" on:click={resetModalStore}
            >Cancel</button
          >
        </div>
      {:else if $ActiveModalStore.isOpen}
        <Form />
      {/if}
    </p>
  </div>
</div>
