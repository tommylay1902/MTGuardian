import React from 'react'
import { Prescription } from '../prescriptions/page'
import Link from 'next/link'
import PrescriptionEdit from './PrescriptionEdit'

const PrescriptionTableView = (props: { prescriptions: Prescription[] }) => {
    const tableHeaders = props.prescriptions.length >=0  ? Object.keys(props.prescriptions[0]) : null
    const tableHeaderExclusions =['id']
    return (
      <>
        {
            tableHeaders == null ? 
            <div>
                no prescriptions found!
            </div> :
            <div className="relative overflow-x-auto shadow-md sm:rounded-lg m-5">
                <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
               
                    <thead className="text-xl text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                        <tr>
                        {tableHeaders.map(h => {
                            if (!tableHeaderExclusions.includes(h)){
                                return (
                                        <th scope="col"className="px-6 py-3"  key={h}>
                                            {h}
                                        </th>
                                    )}
                            }
                        )}
                        <th scope="col" className="px-6 py-3">
                                <span className="sr-only">Edit</span>
                            </th>
                        </tr>
                    
                        
                    </thead>
                    <tbody>
                    
                            {
                                props.prescriptions.map(p => {
                                    return (
                                        <tr key={p.id} className="text-xl bg-white border-b dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-600">
                                            <th scope="row" className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">{p.medication}</th>
                                            <td className="px-6 py-4">{p.dosage}</td>
                                            <td className="px-6 py-4">{p.notes}</td>
                                            <td className="px-6 py-4">{p.started.toString()}</td>
                                            <PrescriptionEdit prescription={p}/>
                                        </tr>
                                    
                                    )
                                })
                            }
                        
                    </tbody>
            </table>
        </div>
           
    
        }
            </>
    )
}

export default PrescriptionTableView