import { Prescription } from "@/app/prescriptions/page";
import React from "react";

type Props = {
  prescriptions: Prescription[];
};
const PrescriptionTableBody: React.FC<Props> = ({ prescriptions }) => {
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
              {new Date(p.started).toDateString()}
            </td>{" "}
            {/* Decreased the padding */}
            <td className="px-6 py-3">
              {" "}
              {/* Keep the last column with the original padding */}
              <button className="rounded-md bg-blue-400 text-white py-2 px-3">
                Edit
              </button>
            </td>
          </tr>
        </React.Fragment>
      ))}
    </tbody>
  );
};

export default PrescriptionTableBody;
