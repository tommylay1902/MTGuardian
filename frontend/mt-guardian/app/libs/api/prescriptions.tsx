import { Prescription } from "../types/Prescription";

export async function getPresentPrescriptions() {
  const res = await fetch(
    "http://0.0.0.0:8000/api/v1/prescription?present=true",
    {
      cache: "no-cache",
    }
  );
  const prescriptions = await res.json();

  return prescriptions;
}

export async function createPrescriptionWithBody(prescription: Prescription) {
  await fetch(`http://0.0.0.0:8000/api/v1/prescription`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ ...prescription }),
  });
}

export async function deletePrescriptionWithId(id: string) {
  await fetch(`http://0.0.0.0:8000/api/v1/prescription/${id}`, {
    method: "DELETE",
  });
}
