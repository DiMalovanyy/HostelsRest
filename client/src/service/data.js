import axios from 'axios'

export const getAllHousings = async () => {
   const res = await axios.get('http://localhost:8080/faculty_hostels')
   return res.data
}

export const getFaculties = async () => {
   const res = await axios.get('http://localhost:8080/faculties')
   return res.data
}

export const getStudentStatus = async () => {
   try {
      const res = await axios.get('http://localhost:8080/user_status', { withCredentials: true })
      
      if (res.status !== 200) return null
      
      return res.data
   }
   catch(error) { return null }
}

export const getHousingRooms = async () => {
   try {
      const res = await axios.get('http://localhost:8080/private/hostel_room_members',
      {withCredentials: true})

      if (res.status !== 200) return null

      return res.data
   }
   catch (error) { return null }
}
