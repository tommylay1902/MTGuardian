"use client";
import React, { useState } from "react";
import { Prescription } from "../../prescriptions/page";
import Link from "next/link";
import { useSearchParams } from "next/navigation";
import PrescriptionTableHeader from "./PrescriptionTableHeader";
import PrescriptionTableBody from "./PrescriptionTableBody";
import EditPrescriptionModal from "../modal/EditPrescriptionModal";
import DeletePrescriptionModal from "../modal/DeletePrescriptionModal";
import AddPrescriptionModal from "../modal/AddPrescriptionModal";

type Props = {
  prescriptions: Prescription[];
};

const PrescriptionTableView: React.FC<Props> = ({ prescriptions }) => {
  const tableHeaders =
    prescriptions.length > 0 ? Object.keys(prescriptions[0]) : null;
  const tableHeaderExclusions = ["id"];

  const [showEditModal, setShowEditModal] = useState<boolean>(false);
  const [showDeleteModal, setShowDeleteModal] = useState<boolean>(false);
  const [showAddModal, setShowAddModal] = useState<boolean>(false);
  const [prescription, setPrescription] = useState<Prescription | null>(null);

  return (
    <>
      {tableHeaders == null ? (
        <div>no prescriptions found!</div>
      ) : (
        <div className="relative  sm:rounded-lg m-5">
          <div className="my-2 ">
            <button
              onClick={() => setShowAddModal(true)}
              className="rounded-lg bg-green-700 text-white py-2 px-3 hover:bg-green-500"
            >
              Create
            </button>
          </div>
          <table className="w-full text-left text-gray-500 dark:text-gray-400">
            <PrescriptionTableHeader
              tableHeaders={tableHeaders}
              tableHeaderExclusions={tableHeaderExclusions}
            />
            <PrescriptionTableBody
              prescriptions={prescriptions}
              setShowEditModal={setShowEditModal}
              setShowDeleteModal={setShowDeleteModal}
              setPrescription={setPrescription}
            />
          </table>
        </div>
      )}
      {showEditModal && prescription != null && (
        <EditPrescriptionModal
          prescription={prescription}
          setShowEditModal={setShowEditModal}
          setPrescription={setPrescription}
        />
      )}
      {showDeleteModal && prescription !== null && (
        <DeletePrescriptionModal
          setShowDeleteModal={setShowDeleteModal}
          prescription={prescription}
        />
      )}
      {showAddModal && (
        <AddPrescriptionModal
          prescription={{
            id: "",
            medication: "",
            dosage: "",
            notes: "",
            started: "",
          }}
          setPrescription={setPrescription}
          setShowAddModal={setShowAddModal}
        />
      )}
    </>
  );
};

export default PrescriptionTableView;
