interface IObjectKeys {
  [key: string]: string | number;
}
export interface Prescription extends IObjectKeys {
  id: string;
  medication: string;
  dosage: string;
  notes: string;
  started: string;
  ended: string;
}
