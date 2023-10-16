import React from "react";
import PrescriptionTableView from "../components/table/PrescriptionTableView";

export interface Prescription {
  id: string;
  medication: string;
  dosage: string;
  notes: string;
  started: string;
}

const PrescriptionPage = async () => {
  const res = await fetch("http://0.0.0.0:8000/api/v1/prescription", {
    cache: "no-cache",
  });
  const prescriptions: Prescription[] = await res.json();
  console.log("printing from server");

  return <PrescriptionTableView prescriptions={prescriptions} />;
};

export default PrescriptionPage;
