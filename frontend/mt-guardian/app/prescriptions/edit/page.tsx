import React, { FormEvent } from 'react'
import { Prescription } from '../page'

const EditPrescription = (props:{prescription:Prescription}) => {
    async function onSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
     
        const formData = new FormData(event.currentTarget)
        const response = await fetch('0.0.0.0:8000/api/v1/prescription', {
          method: 'POST',
          body: formData,
        })
     
        // Handle response if necessary
        const data = await response.json()
        
      }
     
      return (
        <form onSubmit={onSubmit}>
          <input type="text" name="medication" />
          <button type="submit">Submit</button>
        </form>
      )
  
}

export default EditPrescription