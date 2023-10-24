import { Prescription } from "@/app/prescriptions/page";
import { useRouter } from "next/navigation";
import React, { Dispatch, SetStateAction } from "react";

type Props = {
  prescriptions: Prescription[];
  setShowEditModal: Dispatch<SetStateAction<boolean>>;
  setShowDeleteModal: Dispatch<SetStateAction<boolean>>;
  activeModal: boolean;
  setActiveModal: Dispatch<SetStateAction<boolean>>;
  setPrescription: Dispatch<SetStateAction<Prescription | null>>;
};
const PrescriptionTableBody: React.FC<Props> = ({
  prescriptions,
  setShowEditModal,
  setShowDeleteModal,
  activeModal,
  setActiveModal,
  setPrescription,
}) => {
  const convertDate = (date: Date) => {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0"); // Month is 0-based

    const day = String(date.getDate() + 1).padStart(2, "0");
    return `${month}-${day}-${year}`;
  };
  return (
    <tbody>
      {prescriptions.map((p) => (
        <React.Fragment key={p.id}>
          <tr
            key={p.id}
            className="text-xl bg-white border-b dark:bg-gray-800 dark:border-gray-700 dark:hover-bg-gray-600 hover:bg-gray-700 hover:cursor-pointer"
          >
            <th scope="row" className="px-4 py-4 text-gray-900 dark:text-white">
              {p.medication}
            </th>{" "}
            {/* Decreased the padding */}
            <td className="px-4 py-4">{p.dosage}</td>{" "}
            {/* Decreased the padding */}
            <td className="px-4 py-4">{p.notes}</td>{" "}
            {/* Decreased the padding */}
            <td className="px-4 py-4">
              {convertDate(new Date(p.started))}
            </td>{" "}
            <td className="px-4 py-4">
              {p.ended !== null ? convertDate(new Date(p.ended)) : "Present"}
            </td>{" "}
            {/* Decreased the padding */}
            <td className="px-6 py-3 ">
              {" "}
              {/* Keep the last column with the original padding */}
              <button
                className="rounded-md bg-blue-700 text-white py-2 px-3 mr-3 hover:bg-blue-500"
                disabled={activeModal}
                onClick={() => {
                  setPrescription(p);
                  setActiveModal(true);
                  setShowEditModal(true);
                }}
              >
                Edit
              </button>
              <button
                className="rounded-md bg-red-700 text-white py-2 px-3 hover:bg-red-500"
                disabled={activeModal}
                onClick={() => {
                  setPrescription(p);
                  setActiveModal(true);
                  setShowDeleteModal(true);
                }}
              >
                Delete
              </button>
            </td>
          </tr>
        </React.Fragment>
      ))}
    </tbody>
  );
};

export default PrescriptionTableBody;
