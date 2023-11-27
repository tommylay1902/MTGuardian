import { writable } from "svelte/store";

const RegisterStore = writable<boolean>(false);

export default RegisterStore;
