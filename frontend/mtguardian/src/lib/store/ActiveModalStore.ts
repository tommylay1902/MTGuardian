import { writable } from "svelte/store";

const ActiveModalStore = writable(false);

export const updateModal = (isOpen: boolean) => {
  ActiveModalStore.set(isOpen);
};

export default ActiveModalStore;
