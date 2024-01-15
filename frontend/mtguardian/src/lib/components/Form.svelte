<script lang="ts">
  import { resetModalStore } from "$lib/store/ActiveModalStore";
  import FormStore, { resetFormStore } from "$lib/store/Form";
  import HighlightTableRowStore from "$lib/store/HighlightTableRowStore";
  import PrescriptionStore from "$lib/store/PrescriptionStore";
  import { PrescriptionViewHistoryStore } from "$lib/store/PrescriptionViewHistoryStore";

  import {
    createPrescription,
    updatePrescription,
  } from "$lib/utils/formactions";
  import {
    allViewHistory,
    pastViewHistory,
    presentViewHistory,
  } from "$lib/utils/static";
  import { getContext } from "svelte";
  import PrescriptionInput from "./PrescriptionInput.svelte";

  const access: string = getContext("access");

  function determineUpdate(
    data:
      | {
          id: string;
          medication: string;
          dosage: string;
          notes: string;
          started: string;
          ended: string;
        }
      | undefined
  ): boolean {
    if (
      $PrescriptionViewHistoryStore === presentViewHistory &&
      data?.ended === "null"
    ) {
      return true;
    } else if (
      $PrescriptionViewHistoryStore === pastViewHistory &&
      data?.ended !== "null"
    ) {
      return true;
    } else if ($PrescriptionViewHistoryStore === allViewHistory) {
      return true;
    }
    return false;
  }

  async function updatePrescriptionEvent(e: Event) {
    const data = await updatePrescription(e, $FormStore.data.id, access);

    const canStay = determineUpdate(data);

    if (!canStay && data !== undefined) {
      $PrescriptionStore = $PrescriptionStore.filter((obj) => {
        return obj.id !== data.id;
      });
    } else {
      $PrescriptionStore = $PrescriptionStore.map((obj) => {
        const id = $FormStore.data.id;
        if (id === obj.id && data !== undefined) {
          return { ...data };
        } else return obj;
      });

      HighlightTableRowStore.set({
        id: $FormStore.data.id,
        canHighlightAfterCreation: false,
        canHighlightAfterUpdate: true,
      });
    }

    resetModalStore();
    resetFormStore();
  }

  async function createPrescriptionEvent(e: Event) {
    const data = await createPrescription(e, access);

    if (data !== undefined && determineUpdate(data)) {
      PrescriptionStore.update((currentData) => {
        const newData = [
          ...currentData,
          {
            ...data,
          },
        ];

        const sortedData = newData.sort((a, b) => {
          const dateA = a.started !== null ? new Date(a.started).getTime() : 0;
          const dateB = b.started !== null ? new Date(b.started).getTime() : 0;

          return dateB - dateA;
        });

        return sortedData;
      });
    }

    resetModalStore();
    resetFormStore();
  }
</script>

<form
  on:submit|preventDefault={$FormStore.formAction === "createPrescription"
    ? createPrescriptionEvent
    : updatePrescriptionEvent}
>
  <div class="mb-4">
    <PrescriptionInput name="medication" />
  </div>

  <div class="mb-4">
    <PrescriptionInput name="dosage" />
  </div>

  <div class="mb-4">
    <PrescriptionInput name="notes" valueType="textarea" />
  </div>

  <div class="mb-4">
    <PrescriptionInput
      name="started"
      valueType="date"
      mode={$FormStore.formAction === "createPrescription"
        ? "create"
        : "update"}
    />
  </div>

  <div class="mb-4">
    <PrescriptionInput
      name="ended"
      valueType="date"
      mode={$FormStore.formAction === "createPrescription"
        ? "create"
        : "update"}
    />
  </div>
  <div class="flex space-x-1 content-between">
    <button class="btn btn-primary w-1/2" type="submit"> Submit </button>
    <button
      class="btn btn-secondary w-1/2"
      type="button"
      on:click={() => {
        resetModalStore();
        resetFormStore();
      }}
    >
      Cancel
    </button>
  </div>
</form>
