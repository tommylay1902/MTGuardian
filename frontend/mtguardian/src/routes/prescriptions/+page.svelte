<script lang="ts">
  //imports
  import type { PageData } from "./$types";

  import Modal from "$lib/components/Modal.svelte";

  import PrescriptionStore from "$lib/store/PrescriptionStore";

  import PrescriptionTable from "$lib/components/PrescriptionTable.svelte";
  import { showCreatePrescriptionModal } from "$lib/utils/events/modal";
  import PrescriptionMenu from "$lib/components/PrescriptionMenu.svelte";
  import { setContext } from "svelte";

  // load data
  export let data: PageData;
  PrescriptionStore.set(data.prescriptions);

  setContext("access", data.access);
</script>

<div class="m-3 flex flex-col">
  <div>
    <button
      class="btn btn-primary mb-3 mr-4"
      on:click={showCreatePrescriptionModal}>Create Prescription</button
    >
  </div>
  <div>
    <PrescriptionMenu />
  </div>
  <div>
    <PrescriptionTable />
  </div>

  <Modal />
</div>
