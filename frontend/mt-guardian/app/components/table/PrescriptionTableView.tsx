"use client";
import React, { useState } from "react";
import { Prescription } from "../../prescriptions/page";
import Link from "next/link";
import { useSearchParams } from "next/navigation";
import PrescriptionTableHeader from "./PrescriptionTableHeader";
import PrescriptionTableBody from "./PrescriptionTableBody";
import EditPrescriptionModal from "../modal/EditPrescriptionModal";

type Props = {
  prescriptions: Prescription[];
};

const PrescriptionTableView: React.FC<Props> = ({ prescriptions }) => {
  const tableHeaders =
    prescriptions.length > 0 ? Object.keys(prescriptions[0]) : null;
  const tableHeaderExclusions = ["id"];

  const [showModal, setShowModal] = useState<boolean>(false);
  const [prescription, setPrescription] = useState<Prescription | null>(null);

  return (
    <>
      {tableHeaders == null ? (
        <div>no prescriptions found!</div>
      ) : (
        <div className="relative shadow-md sm:rounded-lg m-5">
          <table className="w-full text-left text-gray-500 dark:text-gray-400">
            <PrescriptionTableHeader
              tableHeaders={tableHeaders}
              tableHeaderExclusions={tableHeaderExclusions}
            />
            <PrescriptionTableBody
              prescriptions={prescriptions}
              setShowModal={setShowModal}
              setPrescription={setPrescription}
            />
          </table>
        </div>
      )}
      {showModal && (
        <EditPrescriptionModal
          prescription={prescription}
          setShowModal={setShowModal}
        />
      )}
    </>
  );
};

export default PrescriptionTableView;
