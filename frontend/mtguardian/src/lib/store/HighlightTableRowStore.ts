import type { Modal } from "$lib/types/Modal";
import { writable } from "svelte/store";

const HighlightTableRowStore = writable({
  id: "",
  canHighlightAfterCreation: false,
  canHighlightAfterUpdate: false,
});

export const setHighlightTableRowStore = (data: any) => {
  HighlightTableRowStore.update((current) => {
    return { ...current, ...data };
  });
};

export default HighlightTableRowStore;
