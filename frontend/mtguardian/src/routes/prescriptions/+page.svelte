<script lang="ts">
  //imports
  import type { PageData } from "./$types";
  import type { Prescription } from "$lib/types/Prescription";
  import Modal from "$lib/components/Modal.svelte";
  import ActiveModalStore, { updateModal } from "$lib/store/ActiveModalStore";
  import FormStore from "$lib/store/Form";
  import PrescriptionStore from "$lib/store/PrescriptionStore";

  // load data
  export let data: PageData;
  PrescriptionStore.set(data.prescriptions);

  //page specific variables
  const tableHeaders: string[] = Object.keys($PrescriptionStore[0]);
  const ignoreHeaders: string[] = ["id"];

  function createPrescriptionModal() {
    FormStore.update((current) => {
      return {
        ...current,
        formAction: "?/createPrescription",
        formMethod: "post",
      };
    });
    updateModal({ isOpen: true, header: "Create Prescription", body: "form" });
  }
  function deletePrescriptionModal(id: string) {
    updateModal({
      isOpen: true,
      header: "Delete Prescription",
      body: "Are you sure you want to delete this prescription?",
      id,
    });
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
      <th class="text-3xl text-white">Edit/Delete</th>
    </tr>
  </thead>

  <tbody>
    {#each $PrescriptionStore as p}
      <tr>
        {#each tableHeaders as th}
          {#if !ignoreHeaders.includes(th)}
            <td class="text-white text-2xl"
              >{p[th] == null ? "present" : p[th]}</td
            >
          {/if}
        {/each}
        <td>
          <button class="btn btn-primary">Edit</button>
          <button
            class="btn btn-secondary"
            on:click={() => deletePrescriptionModal(p.id)}>Delete</button
          >
        </td>
      </tr>
    {/each}
  </tbody>
</table>

<Modal />
