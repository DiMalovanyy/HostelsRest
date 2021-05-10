import { useState } from 'react'
import { Link } from 'react-router-dom'

import axios from 'axios'

const Register = () => {
   const [formData, setFormData] = useState({
      name: '',
      email: '',
      password: '',
      password2: ''
   })

   const [notification, setNotification] = useState({
      active: false,
      message: '',
      className: ''
   })

   const { name, email, password, password2 } = formData

   const clearNotification = () => {
      setNotification({ active: false, message: '', className: '' })
   }

   const showNotification = (message, className) => {
      clearNotification()
      setNotification({ active: true, message, className })
      setTimeout(() => clearNotification(), 4000)
   }

   const isValid = () => {
      if (!name){
         showNotification('Name field cannot be empty', 'notification-error')
         return false
      }
      if (!email){
         showNotification('Email field cannot be empty', 'notification-error')
         return false
      }
      if (!password || !password2) {
         showNotification('Password field cannot be empty', 'notification-error')
         return false
      }
      if (password !== password2) {
         showNotification('Passwords do not match', 'notification-error')
         return false
      }

      return true
   }

   const onChange = event => setFormData({ ...formData, [event.target.name]: event.target.value })

   const onSubmit = async event => {
      event.preventDefault()
      
      if (isValid()) {
         const data = { name: name, email: email, password: password }

         try {
            const res = await axios.post('https://pacific-escarpment-18341.herokuapp.com/register', data)
   
            const { error } = res.data
   
            if (error) {
               showNotification(error.description, 'notification-error')
               return
            }
            
            showNotification('Profile created', 'notification-success')
         }
         catch (error) {
            console.log(error)
            showNotification('Network error', 'notification-error')
         }
      }
   }

   return (
      <>
         <h1 className="large text-primary">Sign Up</h1>
         <p className="lead"><i className="fas fa-user"></i> Create Your Account</p>
         {notification.active &&
            <div className={`notification ${notification.className}`}>
               {notification.message}
            </div>
         }
         <form className="form" onSubmit={e => onSubmit(e)}>
            <div className="form-group">
               <input type="text" placeholder="Name" name="name" value={name} onChange={e => onChange(e)} />
            </div>
            <div className="form-group">
               <input type="email" placeholder="Email Address" name="email" value={email} onChange={e => onChange(e)} />
            </div>
            <div className="form-group">
               <input
                  type="password"
                  placeholder="Password"
                  name="password"
                  value={password}
                  onChange={e => onChange(e)}
               />
            </div>
            <div className="form-group">
               <input
                  type="password"
                  placeholder="Confirm Password"
                  name="password2"
                  value={password2}
                  onChange={e => onChange(e)}
               />
            </div>
            <input type="submit" className="btn btn-primary" value="Register" />
         </form>
         <p className="my-1">
            Already have an account? <Link to="/login">Sign In</Link>
         </p>
      </>
   )
}

export default Register
