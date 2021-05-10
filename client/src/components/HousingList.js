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
   const [open, setOpen] = React.useState(true)

   const [items, setItems] = useState([])

   const handleClick = () => {
      setOpen(!open)
   }

   useEffect(() => {
      try {
         (async () => {
            const housings = await getAllHousings()
            console.log(housings)
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
         <ListItem button>
            <ListItemIcon>
               <SendIcon />
            </ListItemIcon>
            <ListItemText primary="Sent mail" />
         </ListItem>
         <ListItem button>
            <ListItemIcon>
               <SendIcon />
            </ListItemIcon>
            <ListItemText primary="Drafts" />
         </ListItem>
         <ListItem button onClick={handleClick}>
            <ListItemIcon>
               <SendIcon />
            </ListItemIcon>
            <ListItemText primary="Inbox" />
            {open ? <ExpandLess /> : <ExpandMore />}
         </ListItem>
         <Collapse in={open} timeout="auto" unmountOnExit>
            <List component="div" disablePadding>
               <ListItem button className={classes.nested}>
                  <ListItemIcon>
                  <StarBorder />
                  </ListItemIcon>
                  <ListItemText primary="Starred" />
               </ListItem>
            </List>
         </Collapse>

      </List>
     )
  }
  else {
     return <h2>No data</h2>
  }
}

export default HousingList
