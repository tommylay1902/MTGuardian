import { Prescription } from "@/app/prescriptions/page";
import { useRouter } from "next/navigation";
import React, { Dispatch, SetStateAction } from "react";
type Props = {
  setShowDeleteModal: Dispatch<SetStateAction<boolean>>;
  setActiveModal: Dispatch<SetStateAction<boolean>>;
  prescription: Prescription;
};
const DeletePrescriptionModal: React.FC<Props> = ({
  setShowDeleteModal,
  setActiveModal,
  prescription,
}) => {
  const router = useRouter();
  const deletePrescription = async (id: string) => {
    await fetch(`http://0.0.0.0:8000/api/v1/prescription/${id}`, {
      method: "DELETE",
    });
    router.refresh();
  };
  return (
    <div
      aria-hidden="true"
      className="fixed top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 z-50 w-full p-4 overflow-x-hidden overflow-y-auto h-[calc(100% - 1rem)] max-h-full md:w-1/2 md:h-auto sm:w-full sm:h-auto"
    >
      <div className="relative w-full max-w-2xl max-h-full">
        <div className="relative bg-white rounded-lg shadow dark:bg-gray-700">
          <div className="flex items-start justify-between p-4 border-b rounded-t dark:border-gray-600 text-white">
            <h3>Are you sure you want to delete this prescription?</h3>
          </div>
          <div className="flex items-start justify-between p-4 border-b rounded-t dark:border-gray-600">
            <button
              type="button"
              onClick={() => {
                setShowDeleteModal(false);
                setActiveModal(false);
                deletePrescription(prescription.id);
              }}
              className="px-4 py-2 text-white bg-blue-500 rounded-md hover:bg-blue-600 focus:outline-none"
            >
              Yes
            </button>
            <button
              onClick={() => {
                setActiveModal(false);
                setShowDeleteModal(false);
              }}
              type="button"
              className="px-4 py-2 dark:bg-red-600 rounded-md text-white hover-bg-red-800"
            >
              Cancel
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default DeletePrescriptionModal;
