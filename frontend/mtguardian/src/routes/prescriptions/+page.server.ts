import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load = (async ({ cookies, url }) => {
  const access = cookies.get("access");
  const fromURL = url.pathname + url.search;
  if (!access) {
    //create this so we can be redirected back to the specific page that we want after logging in

    throw redirect(302, `/auth?redirectTo=${fromURL}`);
  }

  const res = await fetch(
    "http://0.0.0.0:8004/api/v1/prescription?present=true",
    {
      cache: "no-cache",
      headers: {
        Authorization: `Bearer ${access}`,
      },
    }
  );

  //rewrite or refactor
  if (res.status === 401) {
    const response = await fetch("http://0.0.0.0:8004/api/v1/auth/refresh", {
      method: "POST",
      body: JSON.stringify({ access: `${access}` }),
    });

    if (response.status === 401 || response.status === 500) {
      throw redirect(302, `/auth?redirectTo=${fromURL}`);
    }

    const responseToken = await response.json();

    const token = responseToken["access"];

    cookies.set("access", token);

    const res = await fetch(
      "http://0.0.0.0:8004/api/v1/prescription?present=true",
      {
        cache: "no-cache",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    const prescriptions = await res.json();

    return { prescriptions, token };
  }

  const prescriptions = await res.json();

  return { prescriptions, access };
}) satisfies PageServerLoad;
