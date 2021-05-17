import React, { useState, useEffect } from 'react'
import { makeStyles } from '@material-ui/core/styles'
import ListSubheader from '@material-ui/core/ListSubheader'
import List from '@material-ui/core/List'
import ListItem from '@material-ui/core/ListItem'
import ListItemIcon from '@material-ui/core/ListItemIcon'
import ListItemText from '@material-ui/core/ListItemText'
import Collapse from '@material-ui/core/Collapse'
import SendIcon from '@material-ui/icons/Send'
import HouseIcon from '@material-ui/icons/House'
import ExpandLess from '@material-ui/icons/ExpandLess'
import ExpandMore from '@material-ui/icons/ExpandMore'

import { getAllHousings } from '../service/data'

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100%',
    maxWidth: 360,
    backgroundColor: theme.palette.background.paper,
  },
  nested: {
    paddingLeft: theme.spacing(4),
  },
}));

const HousingList = () => {
   const classes = useStyles()

   const [items, setItems] = useState([])

   const handleClick = (faculty) => {
      setItems(items.map(item => item.faculty_name === faculty.faculty_name ?
         ({...item, open: !item.open }) : item))
   }

   useEffect(() => {
      try {
         (async () => {
            const housings = await getAllHousings()
            setItems(housings.map(item => ({ ...item, open: false })))
         })()
      }
      catch (error) {
         console.log('Network error')
      }
   }, [])

  if (items && items.length > 0) {
     return (
      <List
         component="nav"
         aria-labelledby="nested-list-subheader"
         subheader={
         <ListSubheader component="div" id="nested-list-subheader">
            Faculties
         </ListSubheader>
         }
         className={classes.root}
      >
         {items.map((item, i) => (
            <div key={i}>
               <ListItem button onClick={() => handleClick(item)}>
                  <ListItemIcon>
                     <SendIcon />
                  </ListItemIcon>
                  <ListItemText primary={item.faculty_name} />
                  {item.open ? <ExpandLess /> : <ExpandMore />}
               </ListItem>
               <Collapse in={item.open} timeout="auto" unmountOnExit>
                  <List component="div" disablePadding>
                     {item.housings && item.housings.map((housing, j) => (
                        <ListItem key={j+10} button className={classes.nested}>
                           <ListItemIcon>
                              <HouseIcon />
                           </ListItemIcon>
                           <ListItemText primary={housing.hostel_name} />
                        </ListItem>
                     ))}
                  </List>
               </Collapse>
            </div>
         ))}
      </List>
     )
  }
  else {
     return <h2>No data</h2>
  }
}

export default HousingList
