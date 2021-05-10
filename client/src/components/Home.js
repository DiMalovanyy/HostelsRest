import PropTypes from 'prop-types'
import HousingList from './HousingList'

const Home = ({ view }) => {
   return (
      <section id="home">
         <div className="home-left">
            <HousingList />
         </div>
         <div className="home-right">
            
         </div>
      </section>
   )
}

Home.propTypes = {
   view: PropTypes.string.isRequired
}

export default Home
