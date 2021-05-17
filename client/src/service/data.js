import axios from 'axios'

export const getAllHousings = async () => {
   const res = await axios.get('https://pacific-escarpment-18341.herokuapp.com/faculty_hostels')
   return res.data
}

export const getFaculties = async () => {
   const res = await axios.get('https://pacific-escarpment-18341.herokuapp.com/faculties')
   return res.data
}

export const getStudentStatus = async () => {
   try {
      const res = await axios.get('https://pacific-escarpment-18341.herokuapp.com/user_status')
      
      if (res.status !== 200) return null
      
      return res.data
   }
   catch(error) { return null }
}
