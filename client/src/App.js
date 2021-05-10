import { useState, useEffect } from 'react'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import Navbar from './components/Navbar'
import Home from './components/Home'
import Register from './components/Register'
import Login from './components/Login'

import './App.css'

const App = () => {
   const [cookie, setCookie] = useState({})
   const [view, setView] = useState('default')

   useEffect(() => {
      
   })

   return (
      <Router>
         <Navbar />
         <Switch>
            <Route exact path="/" render={() => <Home view={view} />} />
            <Route path="/">
               <section className="container">
                  <Route exact path="/register" component={Register} />
                  <Route exact path="/login" component={Login} />
               </section>
            </Route>
         </Switch>
      </Router>
   )
}

export default App
