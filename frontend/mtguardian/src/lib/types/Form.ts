import type { Prescription } from "./Prescription";
interface IObjectKeys {
  [key: string]: string | number | Prescription;
}
export interface Form extends IObjectKeys {
  formAction: string;
  formMethod: string;
  data: Prescription;
}
