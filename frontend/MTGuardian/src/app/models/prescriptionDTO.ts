export interface PrescriptionDTO {
  medication?: string;
  doseage?: string;
  notes?: 'MALE' | 'FEMALE' | 'OTHER';
  prescribedAt?: Date;
}
