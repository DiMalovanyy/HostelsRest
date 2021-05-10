import { useState } from 'react'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import Navbar from './components/Navbar'
import Home from './components/Home'
import Register from './components/Register'
import Login from './components/Login'

import './App.css'

const App = () => {
   const [user, setUser] = useState({
      loggedIn: false,
      name: '',
      hasAccommodation: false
   })

   return (
      <Router>
         <Navbar />
         <Switch>
            <Route exact path="/" component={Home} />
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
