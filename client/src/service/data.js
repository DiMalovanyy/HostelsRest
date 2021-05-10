import axios from 'axios'

export const getAllHousings = async () => {
   const res = await axios.get('https://pacific-escarpment-18341.herokuapp.com/faculty_hostels')
   return res.data
}
