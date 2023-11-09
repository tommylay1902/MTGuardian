<script lang="ts">
  import { updateModal } from "$lib/store/ActiveModalStore";
  import FormStore from "$lib/store/Form";
  import type { Prescription } from "$lib/types/Prescription";
  import { DateTime } from "luxon";

  let prescription: Prescription;
  if (Object.keys($FormStore.data).length === 0) {
    prescription = {
      id: "",
      medication: "",
      dosage: "",
      notes: "",
      started: "",
      ended: "",
    };
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

      await fetch(`http://0.0.0.0:8000/api/v1/prescription`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ ...prescription }),
      });
    } catch (e) {
      console.log(e);
    }
  }
</script>

<form method="POST" action="?/createPrescription">
  <div class="mb-4">
    <label class="block text-sm font-medium text-white" for="medication">
      Medication
    </label>
    <input
      type="text"
      id="medication"
      name="medication"
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
      class="w-full px-3 py-2 border rounded-md shadow-sm"
    />
  </div>
  <button class="btn btn-primary" type="submit"> Submit </button>
  <button
    class="btn btn-secondary"
    type="button"
    on:click={() => updateModal({ isOpen: false })}
  >
    Cancel
  </button>
</form>
