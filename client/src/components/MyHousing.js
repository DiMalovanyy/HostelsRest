import { useState, useEffect } from 'react'
import axios from 'axios'
import { getFaculties, getStudentStatus } from '../service/data'

const MyHousing = () => {
   const [formData, setFormData] = useState({
      degreeLevel: '',
      sex: '',
      facultyName: ''
   })

   const [status, setStatus] = useState(false)

   const [hostel, setHostel] = useState({
      hostelName: '',
      rooms: []
   })
   
   const [faculties, setFaculties] = useState([])

   const { degreeLevel, sex, facultyName } = formData

   useEffect(() => {
      (async () => {
         const data = await getFaculties()
         setFaculties(data)
         const studStatus = await getStudentStatus()
         if (studStatus) setStatus(studStatus)
      })()
   }, [])

   const onChange = event => setFormData({ ...formData, [event.target.name]: event.target.value })

   const onSubmit = async event => {
      event.preventDefault()

      try {
         const res = await axios.post('http://localhost:8080/private/upgrade_user',
         { degreeLevel: parseInt(degreeLevel), sex, facultyName })
         
         console.log('done', res)
      }
      catch (error) {
         console.log(error)
      }
   }

   return (
      <section id="my-housing">
         <h1 className="large text-primary">My Housing</h1>
         {!status ? (
            <>
               <h2 className="lead text-primary">Apply for Housing</h2>
               <form className="form" onSubmit={e => onSubmit(e)}>
                  <div className="form-group">
                     <label htmlFor="degreeLevel">Degree Level:</label>
                     <select id="degreeLevel" name="degreeLevel" onChange={e => onChange(e)} value={degreeLevel}>
                        <option value="1">1</option>
                        <option value="2">2</option>
                        <option value="3">3</option>
                        <option value="4">4</option>
                     </select>
                  </div>
                  <div className="form-group">
                     <label htmlFor="sex">Sex:</label>
                     <select id="sex" name="sex" onChange={e => onChange(e)} value={sex}>
                        <option value="male">Male</option>
                        <option value="female">Female</option>
                     </select>
                  </div>
                  <div className="form-group">
                     <label htmlFor="facultyName">Faculty:</label>
                     {faculties.length > 0 && (
                        <select id="facultyName" name="facultyName" onChange={e => onChange(e)} value={facultyName}>
                           {faculties.map((item, i) => (<option key={i} value={item}>{item}</option>))}
                        </select>
                     )}
                  </div>
                  <input type="submit" className="btn btn-primary" value="Submit" />
               </form>
            </>
         ) : (
            <>
               <h2 className="lead text-primary">List of Rooms</h2>

            </>
         )
      }
      </section>
   )
}

export default MyHousing
