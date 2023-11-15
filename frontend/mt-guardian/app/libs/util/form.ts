import { Dispatch, MutableRefObject, SetStateAction } from "react";
import { Prescription } from "../types/Prescription";

export const handlePrescriptionFormChange = (
  e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  setPrescriptionForm: Dispatch<SetStateAction<Prescription>>
) => {
  let { name, value } = e.target;

  if (name === "started") {
    value = new Date(value).toISOString();
  }
  setPrescriptionForm((prevPrescription) => {
    if (prevPrescription === null) {
      return {
        id: "", // Provide default values for other properties
        medication: "",
        dosage: "",
        notes: "",
        started: "",
        ended: null,
        [name]: value,
      };
    } else {
      return {
        ...prevPrescription,
        [name]: value,
      };
    }
  });
};
