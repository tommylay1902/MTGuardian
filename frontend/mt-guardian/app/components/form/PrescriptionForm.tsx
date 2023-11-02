import convertDate from "@/app/libs/util/date";
import { Prescription } from "@/app/libs/types/Prescription";

import React, { Dispatch, SetStateAction } from "react";
type Props = {
  prescription: Prescription | null;
  setPrescription: Dispatch<SetStateAction<Prescription | null>>;
  setShowEditModal?: Dispatch<SetStateAction<boolean>>;
  setShowAddModal?: Dispatch<SetStateAction<boolean>>;
  setActiveModal: Dispatch<SetStateAction<boolean>>;

  handleSubmit: (e: React.FormEvent) => Promise<void>;
};
const PrescriptionForm: React.FC<Props> = ({
  prescription,
  setPrescription,
  setShowEditModal,
  setShowAddModal,
  setActiveModal,
  handleSubmit,
}) => {
  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    let { name, value } = e.target;

    if (name === "started") {
      value = new Date(value).toISOString();
    }
    setPrescription((prevPrescription) => {
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

  return (
    <form onSubmit={handleSubmit}>
      <div className="mb-4">
        <label
          className="block text-sm font-medium text-white"
          htmlFor="medication"
        >
          Medication
        </label>
        <input
          type="text"
          id="medication"
          name="medication"
          value={prescription?.medication}
          onChange={handleChange}
          className="w-full px-3 py-2 border rounded-md shadow-sm"
        />
      </div>

      <div className="mb-4">
        <label
          className="block text-sm font-medium text-white"
          htmlFor="dosage"
        >
          Dosage
        </label>
        <input
          type="text"
          id="dosage"
          name="dosage"
          value={prescription?.dosage}
          onChange={handleChange}
          className="w-full px-3 py-2 border rounded-md shadow-sm"
        />
      </div>

      <div className="mb-4">
        <label className="block text-sm font-medium text-white" htmlFor="notes">
          Notes
        </label>
        <textarea
          id="notes"
          name="notes"
          value={prescription?.notes}
          onChange={handleChange}
          className="w-full px-3 py-2 border rounded-md shadow-sm"
        />
      </div>

      <div className="mb-4">
        <label
          className="block text-sm font-medium text-white"
          htmlFor="started"
        >
          Started
        </label>
        <input
          type="date"
          id="started"
          name="started"
          value={
            prescription
              ? convertDate(prescription.started)
              : convertDate(new Date(0).toDateString())
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
            setActiveModal(false);
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
