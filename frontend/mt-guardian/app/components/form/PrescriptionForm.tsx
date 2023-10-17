import { Prescription } from "@/app/prescriptions/page";
import { useRouter } from "next/navigation";
import React, { Dispatch, SetStateAction } from "react";
type Props = {
  prescription?: Prescription;
  setPrescription: Dispatch<SetStateAction<Prescription | null>>;
  setShowEditModal?: Dispatch<SetStateAction<boolean>>;
  setShowAddModal?: Dispatch<SetStateAction<boolean>>;
  handleSubmit: (e: React.FormEvent) => Promise<void>;
  handleChange: (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => void;
};
const PrescriptionForm: React.FC<Props> = ({
  prescription,
  setPrescription,
  setShowEditModal,
  setShowAddModal,
  handleSubmit,
  handleChange,
}) => {
  const convertDate = (date: Date) => {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0"); // Month is 0-based
    const day = String(date.getDate()).padStart(2, "0");
    return `${year}-${month}-${day}`;
  };
  return (
    <form onSubmit={handleSubmit}>
      <div className="mb-4">
        <label
          className="block text-sm font-medium text-gray-700"
          htmlFor="medication"
        >
          Medication
        </label>
        <input
          type="text"
          id="medication"
          name="medication"
          value={prescription ? prescription.medication : ""}
          onChange={handleChange}
          className="w-full px-3 py-2 border rounded-md shadow-sm"
        />
      </div>

      <div className="mb-4">
        <label
          className="block text-sm font-medium text-gray-700"
          htmlFor="dosage"
        >
          Dosage
        </label>
        <input
          type="text"
          id="dosage"
          name="dosage"
          value={prescription ? prescription.dosage : ""}
          onChange={handleChange}
          className="w-full px-3 py-2 border rounded-md shadow-sm"
        />
      </div>

      <div className="mb-4">
        <label
          className="block text-sm font-medium text-gray-700"
          htmlFor="notes"
        >
          Notes
        </label>
        <textarea
          id="notes"
          name="notes"
          value={prescription ? prescription.notes : ""}
          onChange={handleChange}
          className="w-full px-3 py-2 border rounded-md shadow-sm"
        />
      </div>

      <div className="mb-4">
        <label
          className="block text-sm font-medium text-gray-700"
          htmlFor="started"
        >
          Started
        </label>
        <input
          type="date"
          id="started"
          name="started"
          value={
            prescription && prescription !== null
              ? convertDate(new Date(prescription.started))
              : convertDate(new Date(0))
          }
          onChange={handleChange}
          className="w-full px-3 py-2 border rounded-md shadow-sm"
        />
      </div>
      <div className="flex items-center p-6 space-x-2 border-t border-gray-200 rounded-b dark:border-gray-600">
        <button
          type="submit"
          onSubmit={handleSubmit}
          className="px-4 py-2 text-white bg-blue-500 rounded-md hover:bg-blue-600 focus:outline-none"
        >
          Submit
        </button>
        <button
          onClick={() => {
            if (setShowEditModal) {
              setShowEditModal(false);
            }
            if (setShowAddModal) {
              setShowAddModal(false);
            }
          }}
          type="button"
          className="px-4 py-2 dark:bg-red-600 rounded-md text-white hover:bg-red-800"
        >
          Cancel
        </button>
      </div>
    </form>
  );
};

export default PrescriptionForm;
