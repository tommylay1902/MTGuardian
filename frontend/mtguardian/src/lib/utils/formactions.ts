import { resetModalStore } from "$lib/store/ActiveModalStore";
import { resetFormStore } from "$lib/store/Form";
import HighlightTableRowStore from "$lib/store/HighlightTableRowStore";
import PrescriptionStore from "$lib/store/PrescriptionStore";
import type { Writable } from "svelte/store";
import { convertDateHtmlInputStringToISO8601 } from "./date";
import type { Prescription } from "$lib/types/Prescription";
import type { Form } from "$lib/types/Form";

export async function createPrescription(event: Event) {
  try {
    const values = event.target as HTMLFormElement;
    const data = new FormData(values);

    const date = data.get("started");

    let formattedStartedDate = new Date().toDateString();

    if (date !== null) {
      formattedStartedDate = convertDateHtmlInputStringToISO8601(
        date.toString()
      );
    }

    const prescription = {
      medication: data.get("medication"),
      dosage: data.get("dosage"),
      notes: data.get("notes"),
      started: formattedStartedDate,
      ended: data.get("ended"),
    };

    const response = await fetch(`http://0.0.0.0:8000/api/v1/prescription`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ...prescription }),
    });

    const responseId = await response.json();
    const id = responseId["success"];

    return {
      id,
      medication: data.get("medication")?.toString() || "",
      dosage: data.get("dosage")?.toString() || "",
      notes: data.get("notes")?.toString() || "",
      started: formattedStartedDate.toString() || "",
      ended: data.get("ended")?.toString() || "null",
    };
  } catch (e) {
    console.log(e);
  }
}

export async function updatePrescription(event: Event, id: string) {
  try {
    const values = event.target as HTMLFormElement;
    const data = new FormData(values);

    const date = data.get("started");

    let formattedStartedDate = new Date().toDateString();

    if (date !== null) {
      formattedStartedDate = convertDateHtmlInputStringToISO8601(
        date.toString()
      );
    }

    const prescription = {
      medication: data.get("medication"),
      dosage: data.get("dosage"),
      notes: data.get("notes"),
      started: formattedStartedDate,
      ended: data.get("ended"),
    };

    fetch(`http://0.0.0.0:8000/api/v1/prescription/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ...prescription }),
    });

    return {
      id,
      medication: data.get("medication")?.toString() || "",
      dosage: data.get("dosage")?.toString() || "",
      notes: data.get("notes")?.toString() || "",
      started: formattedStartedDate.toString() || "",
      ended: data.get("ended")?.toString() || "null",
    };
  } catch (e) {
    console.log(e);
  }
}
