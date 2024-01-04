<script lang="ts">
  import { PrescriptionViewHistoryStore } from "$lib/store/PrescriptionViewHistoryStore";
  import PrescriptionStore from "$lib/store/PrescriptionStore";
  import {
    allViewHistory,
    pastViewHistory,
    presentViewHistory,
  } from "$lib/utils/static";
  import { getContext } from "svelte";

  const token = getContext("access");
  $: activeTable = $PrescriptionViewHistoryStore;

  async function reloadTable(viewHistory: string) {
    if (viewHistory === $PrescriptionViewHistoryStore) return;
    PrescriptionViewHistoryStore.set(viewHistory);
    const response = await fetch(`${$PrescriptionViewHistoryStore}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    const prescriptions = await response.json();

    $PrescriptionStore = prescriptions;
  }
</script>

<div class="flex justify-center">
  <button
    class="btn btn-md text-white"
    class:btn-info={activeTable === pastViewHistory}
    on:click={() => reloadTable(pastViewHistory)}>Past Medication</button
  >
  <button
    class="btn btn-md text-white btn-info"
    class:btn-info={activeTable === presentViewHistory}
    on:click={() => reloadTable(presentViewHistory)}
  >
    Present</button
  >
  <button
    class="btn btn-md text-white"
    class:btn-info={activeTable === allViewHistory}
    on:click={() => reloadTable(allViewHistory)}>All Medication</button
  >
</div>
