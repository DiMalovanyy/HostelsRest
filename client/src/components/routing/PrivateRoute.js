import { Route, Redirect } from 'react-router-dom'
import PropTypes from 'prop-types'

const PrivateRoute = ({ component: Component, auth: { loggedIn }, ...rest }) => (
   <Route
      {...rest}
      render={routeProps =>
         !loggedIn ? (
            <Redirect to="/login" />
         ) : (
            <Component {...rest} {...routeProps} />
         )
      }
   />
)

PrivateRoute.propTypes = {
   auth: PropTypes.object.isRequired
}

export default PrivateRoute
