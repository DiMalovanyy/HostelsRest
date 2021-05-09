import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import Navbar from './components/Navbar'
import Register from './components/Register'
import Login from './components/Login'

import './App.css'

const App = () => {
   return (
      <Router>
         <Navbar />
         <section className="container">
            <Switch>
               <Route exact path="/register" component={Register} />
               <Route exact path="/login" component={Login} />
            </Switch>
         </section>
      </Router>
   )
}

export default App
