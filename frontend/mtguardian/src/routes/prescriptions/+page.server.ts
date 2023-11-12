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
