import type { Form } from "$lib/types/Form";
import { generatePrescriptionTemplate } from "$lib/types/Prescription";
import { writable } from "svelte/store";

const FormStore = writable<Form>({
  formAction: "POST",
  formMethod: "",
  data: generatePrescriptionTemplate(),
});

export const resetFormStore = () => {
  FormStore.set({
    formAction: "",
    formMethod: "",
    data: generatePrescriptionTemplate(),
  });
};

export default FormStore;
