import type { Modal } from "$lib/types/Modal";
import { writable } from "svelte/store";

const ActiveModalStore = writable<Modal>({
  isOpen: false,
  header: "",
  body: "",
  id: "",
});

export function updateModal(newData: any) {
  ActiveModalStore.update((current) => {
    console.log({
      ...current,
      ...newData,
    });
    return {
      ...current,
      ...newData,
    };
  });
}

export default ActiveModalStore;
