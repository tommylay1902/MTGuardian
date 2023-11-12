import type { Modal } from "$lib/types/Modal";
import { writable } from "svelte/store";

const HighlightTableRowStore = writable({
  id: "",
  canHighlightAfterCreation: false,
  canHighlightAfterUpdate: false,
});

export default HighlightTableRowStore;
