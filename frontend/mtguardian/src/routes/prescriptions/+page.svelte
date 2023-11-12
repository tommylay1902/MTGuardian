<script lang="ts">
  //imports
  import type { PageData } from "./$types";

  import Modal from "$lib/components/Modal.svelte";

  import PrescriptionStore from "$lib/store/PrescriptionStore";

  import { PrescriptionViewHistoryStore } from "$lib/store/PrescriptionViewHistoryStore";

  import { onMount } from "svelte";
  import PrescriptionTable from "$lib/components/PrescriptionTable.svelte";
  import { showCreatePrescriptionModal } from "$lib/utils/events/modal";

  // load data
  export let data: PageData;
  PrescriptionStore.set(data.prescriptions);

  //page specific variables
  let prescriptionHistory: string = "present";

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

  let mounted: boolean = false;
  onMount(() => {
    mounted = true;
  });
</script>

<div class="m-3 flex flex-col">
  <div>
    <button
      class="btn btn-primary mb-3 mr-4"
      on:click={showCreatePrescriptionModal}>Create Prescription</button
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
    <PrescriptionTable />
  </div>

  <Modal />
</div>
