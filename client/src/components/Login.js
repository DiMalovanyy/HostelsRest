import { useState } from 'react'
import { Link } from 'react-router-dom'

const Login = () => {
   const [formData, setFormData] = useState({
      email: '',
      password: ''
   })

   const { email, password } = formData

   const onChange = event => setFormData({ ...formData, [event.target.name]: event.target.value })

   const onSubmit =  async event => {
      event.preventDefault()
      console.log('Success')
   }

   return (
      <>
         <h1 className="large text-primary">Sign In</h1>
         <p className="lead"><i className="fas fa-user"></i> Sign Into Your Account</p>
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
