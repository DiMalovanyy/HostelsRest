import { useState } from 'react'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import Navbar from './components/Navbar'
import Home from './components/Home'
import Register from './components/Register'
import Login from './components/Login'
import MyHousing from './components/MyHousing'

import './App.css'

const App = () => {
   //const [view, setView] = useState('default')
   const [status, setStatus] = useState('new')
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

   return (
      <Router>
         <Navbar />
         <Switch>
            <Route exact path="/" render={() => <Home />} />
            <Route path="/">
               <section className="container">
                  <Route exact path="/register" component={Register} />
                  <Route exact path="/login" render={() => <Login onLogIn={login} />} />
                  <Route exact path="/myhousing" render={() => <MyHousing status={status} />} />
               </section>
            </Route>
         </Switch>
      </Router>
   )
}

export default App
