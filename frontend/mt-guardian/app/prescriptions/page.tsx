import React from "react";
import PrescriptionTableView from "../components/table/PrescriptionTableView";

import { Prescription } from "../libs/types/Prescription";
import { getPresentPrescriptions } from "../libs/api/prescriptions";

const PrescriptionPage = async () => {
  const prescriptions: Prescription[] = await getPresentPrescriptions();

  return <PrescriptionTableView prescriptions={prescriptions} />;
};

export default PrescriptionPage;
