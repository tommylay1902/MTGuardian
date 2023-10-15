import React from 'react'
import PrescriptionTableView from '../components/PrescriptionTableView'

export interface Prescription{
  medication: string,
  dosage: string,
  notes: string,
  started: Date
}

const PrescriptionPage = async () => {
  const res = await fetch("http://0.0.0.0:8000/api/v1/prescription", {
    cache:"no-cache"
  })
  const prescriptions:Prescription[] = await res.json();
  
  return (
    <PrescriptionTableView prescriptions={prescriptions}/>
  )
}

export default PrescriptionPage
