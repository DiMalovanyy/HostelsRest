import { Link } from 'react-router-dom'

const Navbar = () => (
   <nav className="navbar bg-dark">
      <h1>
         <Link to="/"><i className="fas fa-university"></i> Dorms</Link>
      </h1>
      <ul>
         <li><Link to="/register">Register</Link></li>
         <li><Link to="/login">Log In</Link></li>
      </ul>
   </nav>
)

export default Navbar
