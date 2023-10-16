'use client'
import React from 'react'
import { Prescription } from '../prescriptions/page'

const editRow = (medication: string) => {

}

const PrescriptionEdit = (props:{prescription:Prescription}) => {
  return (
    <td className="font-medium text-blue-600 dark:text-blue-500 hover:underline hover:cursor-pointer" onClick={() => editRow(props.prescription.medication)}>Edit</td>
  )
}

export default PrescriptionEdit