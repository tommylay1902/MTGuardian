import type { Action } from "@sveltejs/kit";
import type { PageServerLoad, RequestEvent } from "./$types";
import { DateTime } from "luxon";
import { updateModal } from "$lib/store/ActiveModalStore";

export const load = (async () => {
  const res = await fetch(
    "http://0.0.0.0:8000/api/v1/prescription?present=true",
    {
      cache: "no-cache",
    }
  );
  const prescriptions = await res.json();

  return { prescriptions };
}) satisfies PageServerLoad;

export const actions = {
  createPrescription: async ({ request }) => {
    const data = await request.formData();

    const date = data.get("started");

    let formattedStartedDate = new Date().toDateString();

    if (date !== null) {
      formattedStartedDate = DateTime.fromFormat(
        date.toString(),
        "yyyy-MM-dd"
      ).toFormat("yyyy-MM-dd'T'HH:mm:ssZZ");
    }

    const prescription = {
      medication: data.get("medication"),
      dosage: data.get("dosage"),
      notes: data.get("notes"),
      started: formattedStartedDate,
      ended: data.get("ended"),
    };

    await fetch(`http://0.0.0.0:8000/api/v1/prescription`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ...prescription }),
    });
    updateModal(false);
  },
};
