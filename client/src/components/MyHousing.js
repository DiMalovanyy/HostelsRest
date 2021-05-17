import { useState, useEffect } from 'react'
import PropTypes from 'prop-types'
import axios from 'axios'
import { getFaculties } from '../service/data'

const MyHousing = ({ status }) => {
   const [formData, setFormData] = useState({
      degreeLevel: '',
      sex: '',
      facultyName: ''
   })
   
   const [faculties, setFaculties] = useState([])

   const { degreeLevel, sex, facultyName } = formData

   useEffect(() => {
      (async () => {
         const data = await getFaculties()
         setFaculties(data)
      })()
   }, [])

   const onChange = event => setFormData({ ...formData, [event.target.name]: event.target.value })

   const onSubmit = async event => {
      event.preventDefault()

      try {
         const res = await axios.post('https://pacific-escarpment-18341.herokuapp.com/upgrade_user',
         { degreeLevel, sex, facultyName })
         
         console.log('done', res)
      }
      catch (error) {
         console.log(error)
      }
   }

   return (
      <section id="my-housing">
         <h1 className="large text-primary">My Housing</h1>
         {status === 'new' && (
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
         )}
      </section>
   )
}

MyHousing.propTypes = {
   status: PropTypes.string.isRequired
}

export default MyHousing
