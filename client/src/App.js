import { useState, useEffect } from 'react'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import axios from 'axios'
import Navbar from './components/Navbar'
import Home from './components/Home'
import Register from './components/Register'
import Login from './components/Login'
import MyHousing from './components/MyHousing'
import PrivateRoute from './components/routing/PrivateRoute'
import { getStudentStatus } from './service/data'

import './App.css'

const App = () => {
   const [auth, setAuth] = useState({
      loggedIn: false,
      loading: true
   })

   useEffect(() => {
      login()
   }, [])

   const login = async () => { 
      const status = await getStudentStatus()
      if (status !== null) {
         setAuth({
            loggedIn: true,
            loading: false
         })
      }
      else {
         setAuth({
            loggedIn: false,
            loading: false
         })
      }
   }

   const logout = async () => {
      try {
         await axios.get('http://localhost:8080/private/logout', { withCredentials: true })
      }
      catch(error) { console.error(error) }

      setAuth({
         loggedIn: false,
         loading: false
      })
   }

   return (
      <Router>
         <Navbar loggedIn={auth.loggedIn} logout={logout} />
         <Switch>
            <Route exact path="/" render={() => <Home />} />
            <Route path="/">
               <section className="container">
                  <Route exact path="/register" render={() => <Register auth={auth}/>} />
                  <Route exact path="/login" render={() => <Login onLogIn={login} auth={auth}/>} />
                  <PrivateRoute exact path="/myhousing" component={MyHousing} auth={auth}/>
               </section>
            </Route>
         </Switch>
      </Router>
   )
}

export default App
