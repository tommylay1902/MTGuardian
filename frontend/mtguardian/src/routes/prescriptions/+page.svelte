<script lang="ts">
  //imports
  import type { PageData } from "./$types";
  import {
    generatePrescriptionTemplate,
    type Prescription,
  } from "$lib/types/Prescription";
  import Modal from "$lib/components/Modal.svelte";
  import { updateModal } from "$lib/store/ActiveModalStore";
  import FormStore from "$lib/store/Form";
  import PrescriptionStore from "$lib/store/PrescriptionStore";
  import { convertStringISO8601ToShortDate } from "$lib/utils/date";
  import { PrescriptionViewHistoryStore } from "$lib/store/PrescriptionViewHistoryStore";

  // load data
  export let data: PageData;
  PrescriptionStore.set(data.prescriptions);

  //page specific variables
  const tableHeaders: string[] = Object.keys(generatePrescriptionTemplate());
  const ignoreHeaders: string[] = ["id"];
  let prescriptionHistory: string = "present";

  function createPrescriptionModal() {
    FormStore.update((current) => {
      return {
        ...current,
        formAction: "createPrescription",
        formMethod: "POST",
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

  function editPrescriptionModal(p: Prescription) {
    FormStore.update((current) => {
      return {
        ...current,
        data: { ...p },
        formAction: "updatePrescription",
      };
    });
    updateModal({ isOpen: true, header: "Edit Prescription", body: "form" });
  }

  async function presentMedication() {
    switch (prescriptionHistory) {
      case "present":
        PrescriptionViewHistoryStore.set(
          "http://0.0.0.0:8000/api/v1/prescription?present=true"
        );
        break;
      case "past":
        PrescriptionViewHistoryStore.set(
          "http://0.0.0.0:8000/api/v1/prescription?present=false"
        );
        break;
      default:
        PrescriptionViewHistoryStore.set(
          "http://0.0.0.0:8000/api/v1/prescription"
        );
    }

    const response = await fetch(`${$PrescriptionViewHistoryStore}`);

    const prescriptions = await response.json();

    $PrescriptionStore = prescriptions;
  }
</script>

<div class="m-3 flex flex-col">
  <div>
    <button class="btn btn-primary mb-3 mr-4" on:click={createPrescriptionModal}
      >Create Prescription</button
    >
    <select
      class="select select-bordered w-full max-w-xs text-white"
      bind:value={prescriptionHistory}
      on:change={presentMedication}
    >
      <option selected value="present">Current Prescriptions</option>
      <option value="past">Past Prescriptions</option>
      <option value="all">All Prescriptions</option>
    </select>
  </div>
  <div>
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
                {#if th === "started" || th === "ended"}
                  <td class="text-white text-2xl">
                    {p[th] == null ||
                    p[th] === "null" ||
                    typeof p[th] !== "string"
                      ? "Present"
                      : convertStringISO8601ToShortDate(p[th])}
                  </td>
                {:else}
                  <td class="text-white text-2xl">{p[th]}</td>
                {/if}
              {/if}
            {/each}
            <td>
              <button
                class="btn btn-primary"
                on:click={() => editPrescriptionModal(p)}>Edit</button
              >
              <button
                class="btn btn-secondary"
                on:click={() => deletePrescriptionModal(p.id)}>Delete</button
              >
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>

  <Modal />
</div>
