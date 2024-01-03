interface IObjectKeys {
  [key: string]: string | number | null;
}

export interface Prescription extends IObjectKeys {
  id: string;
  medication: string;
  dosage: string;
  notes: string;
  started: string | null;
  ended: string | null;
}

export const generatePrescriptionTemplate = () => {
  return {
    id: "",
    medication: "",
    dosage: "",
    notes: "",
    started: "",
    ended: "",
  };
};
