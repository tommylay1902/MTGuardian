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
import { create } from "domain";

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
  const [activeModal, setActiveModal] = useState<boolean>(false);
  const [prescription, setPrescription] = useState<Prescription | null>({
    id: "",
    medication: "",
    dosage: "",
    notes: "",
    started: "",
  });
  const [createPrescription, setCreatePrescription] =
    useState<Prescription | null>({
      id: "",
      medication: "",
      dosage: "",
      notes: "",
      started: "",
    });

  return (
    <>
      {tableHeaders == null ? (
        <div>no prescriptions found!</div>
      ) : (
        <div className="relative  sm:rounded-lg m-5">
          <div className="my-2 ">
            <button
              onClick={() => {
                setActiveModal(true);
                setShowAddModal(true);
              }}
              disabled={activeModal}
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
              setActiveModal={setActiveModal}
              activeModal={activeModal}
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
          setActiveModal={setActiveModal}
        />
      )}
      {showDeleteModal && prescription !== null && (
        <DeletePrescriptionModal
          setShowDeleteModal={setShowDeleteModal}
          setActiveModal={setActiveModal}
          prescription={prescription}
        />
      )}
      {showAddModal && (
        <AddPrescriptionModal
          createPrescription={createPrescription}
          setCreatePrescription={setCreatePrescription}
          setShowAddModal={setShowAddModal}
          setActiveModal={setActiveModal}
        />
      )}
    </>
  );
};

export default PrescriptionTableView;
