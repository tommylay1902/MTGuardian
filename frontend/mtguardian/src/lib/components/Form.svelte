<script lang="ts">
</script>

<form>
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
    <!--  <label className="block text-sm font-medium text-white" htmlFor="dosage">
      Dosage
    </label>
    <input
      type="text"
      id="dosage"
      name="dosage"
      value={prescriptionForm.dosage}
      onChange={(e) => handlePrescriptionFormChange(e, setPrescriptionForm)}
      className="w-full px-3 py-2 border rounded-md shadow-sm"
    />
  </div>

  <div className="mb-4">
    <label className="block text-sm font-medium text-white" htmlFor="notes">
      Notes
    </label>
    <textarea
      id="notes"
      name="notes"
      value={prescriptionForm.notes}
      onChange={(e) => handlePrescriptionFormChange(e, setPrescriptionForm)}
      className="w-full px-3 py-2 border rounded-md shadow-sm"
    />
  </div>

  <div className="mb-4">
    <label className="block text-sm font-medium text-white" htmlFor="started">
      Started
    </label>
    <input
      type="date"
      id="started"
      name="started"
      value={prescriptionForm.started
        ? convertDate(prescriptionForm.started)
        : convertDate(new Date(0).toDateString())}
      onChange={(e) => handlePrescriptionFormChange(e, setPrescriptionForm)}
      className="w-full px-3 py-2 border rounded-md shadow-sm"
    />
  </div>
  <div
    className="flex items-center p-6 space-x-2 border-t border-gray-200 rounded-b dark:border-gray-600"
  > -->
    <!-- <button
      type="submit"
      onSubmit={(e) => handleSubmit(e, prescriptionForm)}
      className="px-4 py-2 text-white bg-blue-500 rounded-md hover:bg-blue-600 focus:outline-none"
    >
      Submit
    </button> -->
    <button
      type="button"
      class="px-4 py-2 dark:bg-red-600 rounded-md text-white hover:bg-red-800"
    >
      Cancel
    </button>
  </div>
</form>
