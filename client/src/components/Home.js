import HousingList from './HousingList'
import { ReactComponent as HouseSvg } from './house.svg';

const Home = () => {
   return (
      <section id="home">
         <HousingList className="left" />
         <HouseSvg className="right" style={{height: '300px', width: '600px'}} />
      </section>
   )
}

export default Home
