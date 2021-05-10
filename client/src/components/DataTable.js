import Paper from '@material-ui/core/Paper'
import { Grid, Table, TableHeaderRow } from '@devexpress/dx-react-grid-material-ui'

const columns = [
  { name: 'id', title: 'ID' },
  { name: 'name', title: 'Name' }
]

const rows = [
  { id: 0, name: 'Computer Science' },
  { id: 1, name: 'Psychology' },
]

const DataTable = () => {
   return (
      <Paper>
         <Grid rows={rows} columns={columns}>
            <Table />
            <TableHeaderRow />
         </Grid>
      </Paper>
   )
}
  

export default DataTable
