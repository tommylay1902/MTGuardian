<script lang="ts">
  import {
    generatePrescriptionTemplate,
    type Prescription,
  } from "$lib/types/Prescription";
  import { convertStringISO8601ToShortDate } from "$lib/utils/date";

  import PrescriptionStore from "$lib/store/PrescriptionStore";
  import HighlightTableRowStore from "$lib/store/HighlightTableRowStore";
  import { updateModal } from "$lib/store/ActiveModalStore";
  import FormStore from "$lib/store/Form";
  import {
    deletePrescriptionModal,
    editPrescriptionModal,
  } from "$lib/utils/events/modal";

  const tableHeaders: string[] = Object.keys(generatePrescriptionTemplate());
  const ignoreHeaders: string[] = ["id"];
</script>

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
    {#each $PrescriptionStore as p (p.id)}
      <tr
        class:highlight={$HighlightTableRowStore.id === p.id}
        on:animationend={() => {
          HighlightTableRowStore.set({
            id: "",
            canHighlightAfterCreation: false,
            canHighlightAfterUpdate: false,
          });
        }}
      >
        {#each tableHeaders as th}
          {#if !ignoreHeaders.includes(th)}
            {#if th === "started" || th === "ended"}
              <td class="text-white text-2xl">
                {p[th] == null || p[th] === "null" || typeof p[th] !== "string"
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

<style>
  .highlight {
    animation: highlight 1.4s ease-in-out;
  }

  @keyframes highlight {
    50% {
      background-color: green;
    }
  }
</style>
