import type { Form } from "$lib/types/Form";
import { writable } from "svelte/store";

const FormStore = writable<Form>({
  formAction: "POST",
  formMethod: "",
  data: {
    id: "",
    medication: "",
    dosage: "",
    notes: "",
    started: "",
    ended: "",
  },
});

export default FormStore;
