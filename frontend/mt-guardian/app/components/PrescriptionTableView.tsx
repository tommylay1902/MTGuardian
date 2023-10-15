import React from 'react'
import { Prescription } from '../prescriptions/page'

interface Props{
    prescriptions:Prescription[]
}
const PrescriptionTableView = (props: { prescriptions: Prescription[] }) => {
    const tableHeaders = props.prescriptions.length >=0  ? Object.keys(props.prescriptions[0]) : null
  return (
    <>
    {
        tableHeaders == null ? 
        <div>
            no prescriptions found!
        </div> :
        <table className="table-auto">
            <thead>
                <tr>
                   {tableHeaders?.map(header => <th key={header}>{header}</th>)}
                </tr>
            </thead>
            <tbody>
               {props.prescriptions.map(p => {
                    return (
                        <tr key={p.medication}>
                            <td>{p.medication}</td>
                            <td>{p.dosage}</td>
                            <td>{p.notes}</td>
                            <td>{p.started +""}</td>
                        </tr>)
                    })
                }
            </tbody>
        </table>
   
    }
         </>
  )
}

export default PrescriptionTableView