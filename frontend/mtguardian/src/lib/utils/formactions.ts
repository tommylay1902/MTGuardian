import { convertDateHtmlInputStringToISO8601 } from "./date";

export async function createPrescription(event: Event) {
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
    console.log(JSON.stringify({ ...prescription }));

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
    console.log("hello", e);
  }
}
