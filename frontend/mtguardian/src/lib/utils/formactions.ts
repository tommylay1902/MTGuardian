import toast from "svelte-french-toast";
import { convertDateHtmlInputStringToISO8601 } from "./date";

export async function createPrescription(event: Event, access: string) {
  try {
    const values = event.target as HTMLFormElement;
    const data = new FormData(values);

    const startedDate = data.get("started");
    const endedDate = data.get("ended");

    let formattedStartedDate = new Date().toDateString();

    let formattedEndedDate = new Date().toDateString();

    let isPresent: boolean = false;

    if (startedDate !== null) {
      formattedStartedDate = convertDateHtmlInputStringToISO8601(
        startedDate.toString()
      );
    }

    if (endedDate !== null) {
      if (endedDate === "") {
        isPresent = true;
      } else {
        formattedEndedDate = convertDateHtmlInputStringToISO8601(
          endedDate.toString()
        );
      }
    }
    const prescription = {
      medication: data.get("medication"),
      dosage: data.get("dosage"),
      notes: data.get("notes"),
      started: formattedStartedDate,
      ended: isPresent ? null : formattedEndedDate,
    };

    // const response = await fetch(`http://0.0.0.0:8004/api/v1/prescription`, {
    //   method: "POST",
    //   headers: {
    //     "Content-Type": "application/json",
    //     Authorization: `Bearer ${access}`,
    //   },
    //   body: JSON.stringify({ ...prescription }),
    // });

    const fetchPromise = new Promise(async (resolve, reject) => {
      const res = await fetch(`http://0.0.0.0:8004/api/v1/prescription`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${access}`,
        },
        body: JSON.stringify({ ...prescription }),
      });
      if (res.status === 201) {
        const id = (await res.json())["success"];
        resolve(id);
      } else {
        const body = await res.json();

        reject(body);
      }
    });

    const prescriptionId = await toast.promise(
      fetchPromise,
      {
        loading: "Creating...",
        success: `Successfully created prescription:${prescription.medication}!`,
        error: ({ error }) => `${error}`,
      },
      {
        style: "color:#fff; background: #333;",
      }
    );

    // const responseId = await response.json();
    // const id = responseId["success"];

    return {
      id: prescriptionId,
      medication: data.get("medication")?.toString() || "",
      dosage: data.get("dosage")?.toString() || "",
      notes: data.get("notes")?.toString() || "",
      started: formattedStartedDate.toString() || "null",
      ended: data.get("ended")?.toString() || "null",
    };
  } catch (e) {
    console.log(e);
  }
}

export async function updatePrescription(
  event: Event,
  id: string,
  access: string
) {
  try {
    const values = event.target as HTMLFormElement;
    const data = new FormData(values);

    const startedDate = data.get("started");
    const endedDate = data.get("ended");

    let formattedStartedDate = new Date().toDateString();

    let formattedEndedDate = new Date().toDateString();

    let isPresent: boolean = false;

    if (startedDate !== null) {
      formattedStartedDate = convertDateHtmlInputStringToISO8601(
        startedDate.toString()
      );
    }

    if (endedDate !== null) {
      if (endedDate === "") {
        isPresent = true;
      } else {
        formattedEndedDate = convertDateHtmlInputStringToISO8601(
          endedDate.toString()
        );
      }
    }

    const prescription = {
      id,
      medication: data.get("medication")?.toString() || "",
      dosage: data.get("dosage")?.toString() || "",
      notes: data.get("notes")?.toString() || "",
      started: formattedStartedDate,
      ended: isPresent ? null : formattedEndedDate,
    };

    const fetchPromise = new Promise(async (resolve, reject) => {
      const res = await fetch(`http://0.0.0.0:8004/api/v1/prescription/${id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${access}`,
        },
        body: JSON.stringify({ ...prescription }),
      });
      if (res.status === 200) {
        resolve(prescription.medication);
      } else {
        const body = await res.json();

        reject(body);
      }
    });

    toast.promise(
      fetchPromise,
      {
        loading: "Saving...",
        success: (result) => `Successfully updated ${result}!`,
        error: ({ error }) => `${error}`,
      },
      {
        style: "color:#fff; background: #333;",
      }
    );
    return {
      id,
      medication: data.get("medication")?.toString() || "",
      dosage: data.get("dosage")?.toString() || "",
      notes: data.get("notes")?.toString() || "",
      started: formattedStartedDate.toString() || "null",
      ended: data.get("ended")?.toString() || "null",
    };
  } catch (e) {
    console.log("hello", e);
  }
}
