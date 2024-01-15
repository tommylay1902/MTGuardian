import { writable } from "svelte/store";

export const PrescriptionViewHistoryStore = writable<string>(
  "http://0.0.0.0:8004/api/v1/prescription?present=true"
);
