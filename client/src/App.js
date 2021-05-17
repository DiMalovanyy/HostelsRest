import { useState } from 'react'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import Navbar from './components/Navbar'
import Home from './components/Home'
import Register from './components/Register'
import Login from './components/Login'
import MyHousing from './components/MyHousing'
import PrivateRoute from './components/routing/PrivateRoute'

import './App.css'

const App = () => {
   const [auth, setAuth] = useState({
      loggedIn: false,
      loading: true
   })

   const login = () => { 
      setAuth({
         loggedIn: true,
         loading: false
      })
   }

   const logout = () => {
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
