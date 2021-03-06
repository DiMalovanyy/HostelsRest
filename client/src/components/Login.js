import { useState } from 'react'
import { Link, Redirect } from 'react-router-dom'
import axios from 'axios'

const Login = ({ onLogIn, auth: { loggedIn } }) => {
   const [formData, setFormData] = useState({
      email: '',
      password: ''
   })

   const [notification, setNotification] = useState({
      active: false,
      message: '',
      className: ''
   })

   const { email, password } = formData

   const clearNotification = () => {
      setNotification({ active: false, message: '', className: '' })
   }

   const showNotification = (message, className) => {
      clearNotification()
      setNotification({ active: true, message, className })
      setTimeout(() => clearNotification(), 4000)
   }

   const isValid = () => {
      if (!email){
         showNotification('Email field cannot be empty', 'notification-error')
         return false
      }
      if (!password) {
         showNotification('Password field cannot be empty', 'notification-error')
         return false
      }

      return true
   }

   const onChange = event => setFormData({ ...formData, [event.target.name]: event.target.value })

   const onSubmit =  async event => {
      event.preventDefault()
      
      if (isValid()) {
         try {
            const res = await axios.post('http://localhost:8080/login', { email, password }, {withCredentials: true})

            const { error } = await res.data
   
            if (error) {
               showNotification(error, 'notification-error')
               return
            }

            onLogIn()
         }
         catch (error) {
            if (error.response && error.response.data.error)
               showNotification(error.response.data.error, 'notification-error')
            else
               showNotification(error.message, 'notification-error')
         }
      }
   }

   if (loggedIn) return <Redirect to="/myhousing" />

   return (
      <>
         <h1 className="large text-primary">Sign In</h1>
         <p className="lead"><i className="fas fa-user"></i> Sign In to Your Account</p>
         {notification.active &&
            <div className={`notification ${notification.className}`}>
               {notification.message}
            </div>
         }
         <form className="form" onSubmit={e => onSubmit(e)}>
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
                  minLength="6"
               />
            </div>
            <input type="submit" className="btn btn-primary" value="Log In" />
         </form>
         <p className="my-1">
            Don't have an account? <Link to="/register">Sign Up</Link>
         </p>
      </>
   )
}

export default Login
