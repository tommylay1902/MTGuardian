<script lang="ts">
  import { generatePrescriptionTemplate } from "$lib/types/Prescription";
  import { convertStringISO8601ToShortDate } from "$lib/utils/date";
  import { fly } from "svelte/transition";
  import { circOut } from "svelte/easing";

  import PrescriptionStore from "$lib/store/PrescriptionStore";
  import HighlightTableRowStore from "$lib/store/HighlightTableRowStore";

  import {
    showDeletePrescriptionModal,
    showEditPrescriptionModal,
  } from "$lib/utils/events/modal";

  const tableHeaders: string[] = Object.keys(generatePrescriptionTemplate());
  const ignoreHeaders: string[] = ["id", "notes"];
</script>

<table class="table table-lg table-fixed top-5">
  <thead class="sticky top-0 bg-slate-800 rounded-lg">
    <tr class="border-white">
      {#each tableHeaders as th}
        {#if !ignoreHeaders.includes(th)}
          <th class={`text-3xl text-white py-5`}
            ><strong>{th.toUpperCase()}</strong></th
          >
        {/if}
      {/each}
      <th class="text-3xl text-white py-5">Edit/Delete</th>
    </tr>
  </thead>

  <tbody>
    {#each $PrescriptionStore as p (p.id)}
      <tr
        class="border-white"
        class:highlight={$HighlightTableRowStore.id === p.id}
        transition:fly={{ x: 400, duration: 800, easing: circOut }}
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
            {#if th === "ended"}
              <td class={`text-white text-2xl`} id="date">
                {p[th] == null || p[th] === "null" || typeof p[th] !== "string"
                  ? "Present"
                  : convertStringISO8601ToShortDate(p[th])}
              </td>
            {:else if th === "started"}
              <td class={`text-white text-2xl`} id="date">
                {p[th] == null || p[th] === "null" || typeof p[th] !== "string"
                  ? "Unkown"
                  : convertStringISO8601ToShortDate(p[th])}
              </td>
            {:else}
              <td class="text-white text-2xl break-words">{p[th]}</td>
            {/if}
          {/if}
        {/each}
        <td>
          <button
            class="btn btn-primary"
            on:click={() => showEditPrescriptionModal(p)}>Edit</button
          >
          <button
            class="btn btn-secondary"
            on:click={() => showDeletePrescriptionModal(p.id)}>Delete</button
          >
        </td>
      </tr>
    {/each}
  </tbody>
</table>

<style>
  .highlight {
    animation: highlight 1s ease-in-out;
  }

  @keyframes highlight {
    50% {
      background-color: green;
    }
  }

  td {
    max-width: 15vw;
  }
</style>
