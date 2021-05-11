import { Link } from 'react-router-dom'

const Navbar = ({ view }) => (
   <nav className="navbar bg-dark">
      <h1>
         <Link to="/"><i className="fas fa-university"></i> Student Housing</Link>
      </h1>
      <ul>
         {view === 'user' && (<li><Link to="/myhousing">My Housing</Link></li>)}
         <li><Link to="/register">Register</Link></li>
         <li><Link to="/login">Log In</Link></li>
      </ul>
   </nav>
)

export default Navbar
