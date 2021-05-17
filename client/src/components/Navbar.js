import { Link, withRouter } from 'react-router-dom'

const Navbar = ({ loggedIn, logout, history }) => {
   const onLogOut = () => {
      logout()
      history.push('/')
   }

   return (
      <nav className="navbar bg-dark">
         <h1>
            <Link to="/"><i className="fas fa-university"></i> Student Housing</Link>
         </h1>
         <ul>
            <li><Link to="/myhousing">My Housing</Link></li>
            {loggedIn ? (
               <li><a onClick={onLogOut} style={{cursor: 'pointer'}}>Log Out</a></li>
            ) : (
               <>
                  <li><Link to="/register">Register</Link></li>
                  <li><Link to="/login">Log In</Link></li>
               </>)
            }
            
         </ul>
      </nav>
   )
}

export default withRouter(Navbar)
