import type { PageLoad } from "./$types";

export const load = (async () => {
  const res = await fetch(
    "http://0.0.0.0:8000/api/v1/prescription?present=true",
    {
      cache: "no-cache",
    }
  );
  const prescriptions = await res.json();

  return { prescriptions };
}) satisfies PageLoad;
