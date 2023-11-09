<script lang="ts">
  //imports
  import type { PageData } from "./$types";
  import type { Prescription } from "$lib/types/Prescription";
  import Modal from "$lib/components/Modal.svelte";
  import { updateModal } from "$lib/store/ActiveModalStore";
  import FormStore from "$lib/store/Form";

  // load data
  export let data: PageData;
  let prescriptions: Prescription[] = data.prescriptions;

  //page specific variables
  const tableHeaders: string[] = Object.keys(prescriptions[0]);
  const ignoreHeaders: string[] = ["id"];

  function createPrescriptionModal() {
    FormStore.update((current) => {
      return {
        ...current,
        formAction: "?/createPrescription",
        formMethod: "post",
      };
    });
    updateModal(true);
  }
</script>

<button class="btn btn-primary" on:click={createPrescriptionModal}
  >Create</button
>
<table class="table table-lg">
  <thead>
    <tr>
      {#each tableHeaders as th}
        {#if !ignoreHeaders.includes(th)}
          <th class="text-3xl text-white"
            ><strong>{th.toUpperCase()}</strong></th
          >
        {/if}
      {/each}
    </tr>
  </thead>

  <tbody>
    {#each data.prescriptions as p}
      <tr>
        {#each tableHeaders as th}
          {#if !ignoreHeaders.includes(th)}
            <td class="text-white text-2xl"
              >{p[th] == null ? "present" : p[th]}</td
            >
          {/if}
        {/each}
      </tr>
    {/each}
  </tbody>
</table>

<Modal header="Create Prescription" />
