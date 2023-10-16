import { Prescription } from "@/app/prescriptions/page";
import React, { Dispatch, SetStateAction } from "react";
type Props = {
  prescription: Prescription | null;
  setShowModal: Dispatch<SetStateAction<boolean>>;
  setPrescription: Dispatch<SetStateAction<Prescription | null>>;
};
const EditPrescriptionModal: React.FC<Props> = ({
  prescription,
  setShowModal,
  setPrescription,
}) => {
  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;
    setPrescription((prevPrescription) => {
      if (prevPrescription === null) {
        return {
          id: "", // Provide default values for other properties
          medication: "",
          dosage: "",
          notes: "",
          started: "",
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

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // Handle form submission here
    console.log("Submitted Prescription:", prescription);
  };

  const convertDate = (date: Date) => {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0"); // Month is 0-based
    const day = String(date.getDate()).padStart(2, "0");
    return `${year}-${month}-${day}`;
  };

  return (
    <div
      aria-hidden="true"
      className="fixed top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 z-50 w-full p-4 overflow-x-hidden overflow-y-auto h-[calc(100% - 1rem)] max-h-full md:w-1/2 md:h-auto sm:w-full sm:h-auto"
    >
      <div className="relative w-full max-w-2xl max-h-full">
        <div className="relative bg-white rounded-lg shadow dark:bg-gray-700">
          <div className="flex items-start justify-between p-4 border-b rounded-t dark:border-gray-600">
            <h1 className="text-xl font-semibold text-gray-900 dark:text-white">
              Edit {prescription?.medication}
            </h1>
            <button
              type="button"
              className="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ml-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white"
              onClick={() => setShowModal(false)}
            >
              <svg
                className="w-3 h-3"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 14 14"
              >
                <path
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
                />
              </svg>
              <span className="sr-only">Close modal</span>
            </button>
          </div>
          <div className="p-6 space-y-6">
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
                  value={prescription?.medication}
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
                  value={prescription?.dosage}
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
                  value={prescription?.notes}
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
                    prescription !== null
                      ? convertDate(new Date(prescription?.started))
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
                  onClick={() => setShowModal(false)}
                  type="button"
                  className="px-4 py-2 dark:bg-red-600 rounded-md text-white hover:bg-red-800"
                >
                  Cancel
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};

export default EditPrescriptionModal;
