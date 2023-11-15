import type { Prescription } from "$lib/types/Prescription";
import { writable } from "svelte/store";

const PrescriptionStore = writable<Prescription[]>([]);

export default PrescriptionStore;
