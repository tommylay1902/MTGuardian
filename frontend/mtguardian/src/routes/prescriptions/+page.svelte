<script lang="ts">
  //imports
  import type { PageData } from "./$types";
  import type { Prescription } from "$lib/types/Prescription";
  import Modal from "$lib/components/Modal.svelte";
  import ActiveModalStore, { updateModal } from "$lib/store/ActiveModalStore";

  // load data
  export let data: PageData;
  let prescriptions: Prescription[] = data.prescriptions;

  //page specific variables
  let isModalOpen: boolean = false;
  const tableHeaders: string[] = Object.keys(prescriptions[0]);
  const ignoreHeaders: string[] = ["id"];
</script>

<!-- svelte-ignore missing-declaration -->
<button class="btn btn-primary" on:click={() => updateModal(true)}
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
    <tr>
      {#each prescriptions as p}
        {#each tableHeaders as th}
          {#if !ignoreHeaders.includes(th)}
            <td class="text-white text-2xl"
              >{p[th] == null ? "present" : p[th]}</td
            >
          {/if}
        {/each}
      {/each}
    </tr>
  </tbody>
</table>

<Modal header="Create Prescription" />
