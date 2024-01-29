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
  refills: string | null;
  owner: string;
}

export const generatePrescriptionTemplate = () => {
  return {
    id: '',
    medication: '',
    dosage: '',
    notes: '',
    started: '',
    ended: '',
    refills: '',
    owner: '',
  };
};
