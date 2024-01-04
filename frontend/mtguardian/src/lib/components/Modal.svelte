<script lang="ts">
  import Form from "./Form.svelte";
  import ActiveModalStore, {
    resetModalStore,
  } from "$lib/store/ActiveModalStore";
  import PrescriptionStore from "$lib/store/PrescriptionStore";
  import { resetFormStore } from "$lib/store/Form";
  import { getContext } from "svelte";
  import toast from "svelte-french-toast";

  const access = getContext("access");

  async function deletePrescriptionEvent() {
    try {
      const toastFetch: Promise<boolean> = new Promise(
        async (resolve, reject) => {
          try {
            const data = await fetch(
              `http://0.0.0.0:8004/api/v1/prescription/${$ActiveModalStore.id}`,
              {
                method: "DELETE",
                headers: {
                  Authorization: `Bearer ${access}`,
                },
              }
            );
            if (data.status === 200) {
              resolve(true);
            } else {
              const error = await data.json();
              reject(false);
            }
          } catch (e) {
            reject(false);
          }
        }
      );

      const isSuccess = await toast.promise(
        toastFetch,
        {
          loading: "Deleting...",
          success: "Successfully deleted prescription",
          error: "Error deleting prescription",
        },
        {
          style: "color:#fff; background: #333;",
        }
      );
      if (isSuccess) {
        PrescriptionStore.update((currentData) => {
          currentData = currentData.filter(
            (curr) => curr.id !== $ActiveModalStore.id
          );
          return currentData;
        });
      }

      resetFormStore();
      resetModalStore();
    } catch (e) {
      resetFormStore();
      resetModalStore();
    }
  }
</script>

<div class="modal" class:modal-open={$ActiveModalStore.isOpen}>
  <div class="modal-box">
    <h3 class="font-bold text-lg">{$ActiveModalStore.header}</h3>
    <p class="py-4">
      {#if $ActiveModalStore.body !== "form" && $ActiveModalStore.isOpen}
        <p>{$ActiveModalStore.body}</p>
        <div class="flex flex-row space-x-3 pt-4">
          <button
            class="btn btn-primary w-1/2"
            on:click={deletePrescriptionEvent}>Delete</button
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
