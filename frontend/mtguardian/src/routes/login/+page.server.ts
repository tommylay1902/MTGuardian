import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load = (async () => {
  return {};
}) satisfies PageServerLoad;

export const actions = {
  default: async ({ cookies, request, url }) => {
    const data = await request.formData();
    console.log(data.get("email"), data.get("password"));
    const response = await fetch("http:localhost:8004/api/v1/auth/login", {
      method: "POST",
      headers: {
        "content-type": "application/json",
      },
      body: JSON.stringify({
        email: data.get("email"),
        password: data.get("password"),
      }),
    });

    const responseToken = await response.json();
    const token = responseToken["access"];
    console.log("GETTING TOKEN", token);

    cookies.set("access", "test");
    const redirectTo = url.searchParams.get("redirectTo");
    if (redirectTo) {
      throw redirect(302, `/${redirectTo.slice(1)}`);
    }
    throw redirect(302, "/");
  },
};
