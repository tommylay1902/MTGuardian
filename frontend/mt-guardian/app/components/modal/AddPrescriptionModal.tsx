import { Prescription } from "@/app/libs/types/Prescription";
import { createPrescriptionWithBody } from "@/app/libs/api/prescriptions";
import { useRouter } from "next/navigation";
import React, { Dispatch, SetStateAction } from "react";
import PrescriptionForm from "../form/PrescriptionForm";
type Props = {
  createPrescription: Prescription | null;
  setShowAddModal: Dispatch<SetStateAction<boolean>>;
  setActiveModal: Dispatch<SetStateAction<boolean>>;
  setCreatePrescription: Dispatch<SetStateAction<Prescription | null>>;
};
const AddPrescriptionModal: React.FC<Props> = ({
  setShowAddModal,
  setCreatePrescription,
  setActiveModal,
  createPrescription,
}) => {
  const router = useRouter();

  const handleSubmit = async (
    e: React.FormEvent,
    createPrescription: Prescription
  ) => {
    e.preventDefault();
    try {
      if (createPrescription != null && createPrescription?.id !== null) {
        await createPrescriptionWithBody(createPrescription);
        setActiveModal(false);
        setShowAddModal(false);
        setCreatePrescription({
          id: "",
          medication: "",
          dosage: "",
          notes: "",
          started: "",
          ended: null,
        });
        router.refresh();
      }
    } catch (e) {
      console.log(e);
    }
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
              Add Prescription
            </h1>
            <button
              type="button"
              className="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ml-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white"
              onClick={() => {
                setActiveModal(false);
                setShowAddModal(false);
              }}
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
            <PrescriptionForm
              prescription={createPrescription}
              setPrescription={setCreatePrescription}
              setShowAddModal={setShowAddModal}
              setActiveModal={setActiveModal}
              handleSubmit={handleSubmit}
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default AddPrescriptionModal;
