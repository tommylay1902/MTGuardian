import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load = (async ({ cookies, url }) => {
  const access = cookies.get("access");
  if (!access) {
    //create this so we can be redirected back to the specific page that we want after logging in
    const fromURL = url.pathname + url.search;
    throw redirect(302, `/auth?redirectTo=${fromURL}`);
  }

  const res = await fetch(
    "http://0.0.0.0:8000/api/v1/prescription?present=true",
    {
      cache: "no-cache",
    }
  );

  const prescriptions = await res.json();

  return { prescriptions };
}) satisfies PageServerLoad;
