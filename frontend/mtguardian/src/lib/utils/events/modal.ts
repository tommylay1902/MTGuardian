import { updateModal } from "$lib/store/ActiveModalStore";
import FormStore from "$lib/store/Form";
import type { Prescription } from "$lib/types/Prescription";

export function createPrescriptionModal() {
  FormStore.update((current) => {
    return {
      ...current,
      formAction: "createPrescription",
      formMethod: "POST",
    };
  });

  updateModal({ isOpen: true, header: "Create Prescription", body: "form" });
}

export function deletePrescriptionModal(id: string) {
  updateModal({
    isOpen: true,
    header: "Delete Prescription",
    body: "Are you sure you want to delete this prescription?",
    id,
  });
}

export function editPrescriptionModal(p: Prescription) {
  FormStore.update((current) => {
    return {
      ...current,
      data: { ...p },
      formAction: "updatePrescription",
    };
  });
  updateModal({ isOpen: true, header: "Edit Prescription", body: "form" });
}
