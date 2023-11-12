<script lang="ts">
  import { resetModalStore, updateModal } from "$lib/store/ActiveModalStore";
  import FormStore, { resetFormStore } from "$lib/store/Form";
  import HighlightTableRowStore from "$lib/store/HighlightTableRowStore";
  import PrescriptionStore from "$lib/store/PrescriptionStore";

  import { convertDateHtmlInputStringToISO8601 } from "$lib/utils/date";
  import PrescriptionInput from "./PrescriptionInput.svelte";

  async function createPrescription(event: Event) {
    try {
      const values = event.target as HTMLFormElement;
      const data = new FormData(values);

      const date = data.get("started");

      let formattedStartedDate = new Date().toDateString();

      if (date !== null) {
        formattedStartedDate = convertDateHtmlInputStringToISO8601(
          date.toString()
        );
      }

      const prescription = {
        medication: data.get("medication"),
        dosage: data.get("dosage"),
        notes: data.get("notes"),
        started: formattedStartedDate,
        ended: data.get("ended"),
      };

      const response = await fetch(`http://0.0.0.0:8000/api/v1/prescription`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ ...prescription }),
      });

      const responseId = await response.json();
      const id = responseId["success"];

      PrescriptionStore.update((currentData) => [
        ...currentData,
        {
          id,
          medication: data.get("medication")?.toString() || "",
          dosage: data.get("dosage")?.toString() || "",
          notes: data.get("notes")?.toString() || "",
          started: formattedStartedDate.toString() || "",
          ended: data.get("ended")?.toString() || "null",
        },
      ]);

      resetModalStore();
      resetFormStore();
    } catch (e) {
      console.log(e);
    }
  }

  async function updatePrescription(event: Event) {
    try {
      const values = event.target as HTMLFormElement;
      const data = new FormData(values);

      const date = data.get("started");

      let formattedStartedDate = new Date().toDateString();

      if (date !== null) {
        formattedStartedDate = convertDateHtmlInputStringToISO8601(
          date.toString()
        );
      }

      const prescription = {
        medication: data.get("medication"),
        dosage: data.get("dosage"),
        notes: data.get("notes"),
        started: formattedStartedDate,
        ended: data.get("ended"),
      };

      fetch(`http://0.0.0.0:8000/api/v1/prescription/${$FormStore.data.id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ ...prescription }),
      });

      $PrescriptionStore = $PrescriptionStore.map((obj) => {
        const id = $FormStore.data.id;
        if (id === obj.id) {
          return {
            id,
            medication: data.get("medication")?.toString() || "",
            dosage: data.get("dosage")?.toString() || "",
            notes: data.get("notes")?.toString() || "",
            started: formattedStartedDate.toString() || "",
            ended: data.get("ended")?.toString() || "null",
          };
        } else return obj;
      });

      HighlightTableRowStore.set({
        id: $FormStore.data.id,
        canHighlightAfterCreation: false,
        canHighlightAfterUpdate: true,
      });

      resetModalStore();
      resetFormStore();
    } catch (e) {
      console.log(e);
    }
  }
</script>

<form
  on:submit|preventDefault={$FormStore.formAction === "createPrescription"
    ? createPrescription
    : updatePrescription}
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
