import React, { useState, useEffect } from 'react'
import { makeStyles } from '@material-ui/core/styles';
import ListSubheader from '@material-ui/core/ListSubheader';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Collapse from '@material-ui/core/Collapse';
import SendIcon from '@material-ui/icons/Send';
import ExpandLess from '@material-ui/icons/ExpandLess';
import ExpandMore from '@material-ui/icons/ExpandMore';
import StarBorder from '@material-ui/icons/StarBorder';

import { getAllHousings, getFaculties } from '../service/data'

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
   const [open, setOpen] = React.useState(true)

   const [items, setItems] = useState([])

   const handleClick = () => {
      setOpen(!open)
   }

   useEffect(() => {
      try {
         (async () => {
            const housings = await getAllHousings()
            setItems(housings)
         })()
      }
      catch (error) {
         console.log('couldnt get data')
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
            <ListItem key={i} button onClick={handleClick}>
               <ListItemIcon>
                  <SendIcon />
               </ListItemIcon>
               <ListItemText primary={item.faculty_name} />
            </ListItem>
         ))}
      </List>
     )
  }
  else {
     return <h2>No data</h2>
  }
}

export default HousingList
