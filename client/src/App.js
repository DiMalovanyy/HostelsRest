import { useState } from 'react'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import Navbar from './components/Navbar'
import Home from './components/Home'
import Register from './components/Register'
import Login from './components/Login'

import './App.css'

const App = () => {
   const [cookie, setCookie] = useState({})
   const [view, setView] = useState('default')

   const logIn = () => {
      setView('user')
   }

   return (
      <Router>
         <Navbar />
         <Switch>
            <Route exact path="/" render={() => <Home />} />
            <Route path="/">
               <section className="container">
                  <Route exact path="/register" component={Register} />
                  <Route exact path="/login" render={() => <Login onLogIn={logIn} />} />
               </section>
            </Route>
         </Switch>
      </Router>
   )
}

export default App
