import { fail, redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load = (async () => {
  return {};
}) satisfies PageServerLoad;

export const actions = {
  login: async ({ cookies, request, url }) => {
    const data = await request.formData();

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
    console.log(response.status);
    if (response.status === 404) {
      return fail(404, { message: "failed to login" });
    }

    const responseToken = await response.json();
    const token = responseToken["access"];

    cookies.set("access", token);

    const redirectTo = url.searchParams.get("redirectTo");

    if (redirectTo !== "null" && redirectTo) {
      throw redirect(302, `/${redirectTo.slice(1)}`);
    }

    throw redirect(302, "/");
  },

  register: async ({ cookies, request, url }) => {
    const data = await request.formData();
    const response = await fetch("http:localhost:8004/api/v1/auth/register", {
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
    if (!token) {
      return fail(409, { message: "failed to login" });
    }
    cookies.set("access", token);
    const redirectTo = url.searchParams.get("redirectTo");
    if (redirectTo && redirectTo !== "null") {
      throw redirect(302, `/${redirectTo.slice(1)}`);
    }
    throw redirect(302, "/");
  },
};
