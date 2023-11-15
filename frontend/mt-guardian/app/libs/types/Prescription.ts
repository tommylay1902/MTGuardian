export interface Prescription {
  id: string;
  medication: string;
  dosage: string;
  notes: string;
  started: string;
  ended: string | null;
}
