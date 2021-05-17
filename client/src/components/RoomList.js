import React, { useState, useEffect } from 'react'
import { makeStyles } from '@material-ui/core/styles'
import ListSubheader from '@material-ui/core/ListSubheader'
import List from '@material-ui/core/List'
import ListItem from '@material-ui/core/ListItem'
import ListItemIcon from '@material-ui/core/ListItemIcon'
import ListItemText from '@material-ui/core/ListItemText'
import Collapse from '@material-ui/core/Collapse'
import SendIcon from '@material-ui/icons/Send'
import ExpandLess from '@material-ui/icons/ExpandLess'
import ExpandMore from '@material-ui/icons/ExpandMore'
import FaceIcon from '@material-ui/icons/Face'

import { getHousingRooms } from '../service/data'

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

const RoomList = () => {
   const classes = useStyles()

   const [items, setItems] = useState({})
   const [loading, setLoading] = useState(true)

   const handleClick = (room) => {
      let rooms = items.rooms.map(item => item.num === room.num ? ({...item, open: !item.open }) : item)
      setItems({ ...items, rooms})
   }

   useEffect(() => {
      try {
         (async () => { 
            try {
               const hostel = await getHousingRooms()
               setItems(hostel)
               setLoading(false)
            }
            catch (error) {
               console.log('Network error')
               setLoading(false)
            }
         })()
      }
      catch (error) {
         console.log('Network error')
         setLoading(false)
      }
   }, [])

   if (loading) return <h2>Loading...</h2>

   if (items && items.rooms.length > 0) {
      return (
         <List
            component="nav"
            aria-labelledby="nested-list-subheader"
            subheader={
            <ListSubheader component="div" id="nested-list-subheader">
               Hostel {items.hostel_name} rooms:
            </ListSubheader>
            }
            className={classes.root}
         >
            {items.rooms.map((item, i) => (
               <div key={i}>
                  <ListItem button onClick={() => handleClick(item)}>
                     <ListItemIcon>
                        <SendIcon />
                     </ListItemIcon>
                     <ListItemText primary={item.num} />
                     {item.open ? <ExpandLess /> : <ExpandMore />}
                  </ListItem>
                  <Collapse in={item.open} timeout="auto" unmountOnExit>
                     <List component="div" disablePadding>
                        {item.names && item.names.map((name, j) => (
                           <ListItem key={j+10} button className={classes.nested}>
                              <ListItemIcon>
                                 <FaceIcon />
                              </ListItemIcon>
                              <ListItemText primary={name} />
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

export default RoomList
