interface IObjectKeys {
  [key: string]: string | number | boolean;
}
export interface Modal extends IObjectKeys {
  isOpen: boolean;
  header: string;
  body: string;
  id: string;
}
