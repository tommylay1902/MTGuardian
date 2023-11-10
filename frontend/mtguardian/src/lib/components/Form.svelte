<script lang="ts">
  import { updateModal } from "$lib/store/ActiveModalStore";
  import FormStore from "$lib/store/Form";
  import PrescriptionStore from "$lib/store/PrescriptionStore";
  import {
    generatePrescriptionTemplate,
    type Prescription,
  } from "$lib/types/Prescription";
  import { DateTime } from "luxon";

  let prescriptionFormData: Prescription;

  if (Object.keys($FormStore.data).length === 0) {
    prescriptionFormData = generatePrescriptionTemplate();
  } else {
    prescriptionFormData = $FormStore.data;
  }

  function convertDate(date: string) {
    const parsedDate = DateTime.fromISO(date, { zone: "utc" }).toFormat(
      "yyyy-MM-dd"
    );
    return parsedDate;
  }

  async function createPrescription(event: Event) {
    try {
      const values = event.target as HTMLFormElement;
      const data = new FormData(values);

      const date = data.get("started");

      let formattedStartedDate = new Date().toDateString();

      if (date !== null) {
        formattedStartedDate = DateTime.fromFormat(
          date.toString(),
          "yyyy-MM-dd"
        ).toFormat("yyyy-MM-dd'T'HH:mm:ssZZ");
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

      const id = await response.json();

      PrescriptionStore.update((currentData) => [
        ...currentData,
        {
          id: id["success"],
          medication: data.get("medication")?.toString() || "",
          dosage: data.get("dosage")?.toString() || "",
          notes: data.get("notes")?.toString() || "",
          started: formattedStartedDate.toString() || "",
          ended: data.get("ended")?.toString() || "null",
        },
      ]);

      $FormStore.data = generatePrescriptionTemplate();
      updateModal({ isOpen: false });
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
        formattedStartedDate = DateTime.fromFormat(
          date.toString(),
          "yyyy-MM-dd"
        ).toFormat("yyyy-MM-dd'T'HH:mm:ssZZ");
      }

      const prescription = {
        medication: data.get("medication"),
        dosage: data.get("dosage"),
        notes: data.get("notes"),
        started: formattedStartedDate,
        ended: data.get("ended"),
      };

      await fetch(
        `http://0.0.0.0:8000/api/v1/prescription/${$FormStore.data.id}`,
        {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ ...prescription }),
        }
      );

      const response = await fetch("http://0.0.0.0:8000/api/v1/prescription", {
        cache: "no-cache",
      });
      const prescriptions = await response.json();

      $PrescriptionStore = [...prescriptions];
      $PrescriptionStore = $PrescriptionStore;

      $FormStore.data = generatePrescriptionTemplate();
      updateModal({ isOpen: false });
    } catch (e) {
      console.log(e);
    }
  }
</script>

<!-- method="POST" action={`?/${$FormStore.formAction}`} -->
<form
  on:submit|preventDefault={$FormStore.formAction === "createPrescription"
    ? createPrescription
    : updatePrescription}
>
  <div class="mb-4">
    <label class="block text-sm font-medium text-white" for="medication">
      Medication
    </label>
    <input
      type="text"
      id="medication"
      name="medication"
      value={$FormStore.data.medication}
      class="w-full px-3 py-2 border rounded-md shadow-sm"
    />
  </div>

  <div class="mb-4">
    <label class="block text-sm font-medium text-white" for="dosage">
      Dosage
    </label>
    <input
      type="text"
      id="dosage"
      name="dosage"
      value={$FormStore.data.dosage}
      class="w-full px-3 py-2 border rounded-md shadow-sm"
    />
  </div>

  <div class="mb-4">
    <label class="block text-sm font-medium text-white" for="notes">
      Notes
    </label>
    <textarea
      id="notes"
      name="notes"
      value={$FormStore.data.notes}
      class="w-full px-3 py-2 border rounded-md shadow-sm"
    />
  </div>

  <div class="mb-4">
    <label class="block text-sm font-medium text-white" for="started">
      Started
    </label>
    <input
      type="date"
      id="started"
      name="started"
      value={convertDate($FormStore.data.started)}
      class="w-full px-3 py-2 border rounded-md shadow-sm"
    />
  </div>
  <button class="btn btn-primary" type="submit"> Submit </button>
  <button
    class="btn btn-secondary"
    type="button"
    on:click={() => {
      ($FormStore.data = generatePrescriptionTemplate()),
        updateModal({
          isOpen: false,
        });
    }}
  >
    Cancel
  </button>
</form>
