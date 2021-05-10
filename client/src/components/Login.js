import { useState } from 'react'
import { Link } from 'react-router-dom'

const Login = ({ history }) => {
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
            const res = await fetch('https://pacific-escarpment-18341.herokuapp.com/login', {
               method: 'POST',
               headers: { 'Content-Type': 'application/json' },
               body: JSON.stringify({ email, password })
            })
   
            const { error } = await res.json()
   
            if (error) {
               showNotification(error.description, 'notification-error')
               return
            }
            
            // Navigate to the home page once logged in
            history.push('/')
         }
         catch (error) {
            showNotification('Network error', 'notification-error')
         }
      }
   }

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
