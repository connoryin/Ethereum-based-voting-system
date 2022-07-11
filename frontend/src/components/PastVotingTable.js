import React, {useEffect, useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import TablePagination from '@mui/material/TablePagination';
import Link from '@material-ui/core/Link';

const useStyles = makeStyles(theme => ({
    table: {
        borderWidth: 0,
    },
    pagination: {
        position: 'absolute',
        top: '450px',
        left: '350px',
    },
}));

// const votingList = [ // get this from backend
//   {name: 'Name 1', result: ''},
//   {name: 'Name 2', result: ''},
//   {name: 'Name 3', result: ''},
//   {name: 'Name 4', result: ''},
// ];

// votingList.forEach(({name, result})=>{
//     rows.push({eventName: name, id: result});
// })

let fetched = false

function PastVotingTable() {
    const classes = useStyles();
    const [page, setPage] = React.useState(0);
    const [rows, setrows] = React.useState([]);
    const rowsPerPage = 5;
    const handleChangePage = (event, newPage) => {
        setPage(newPage);
    };

    useEffect(()=>{
        if (fetched && rows.length) return
        fetch(
            "http://localhost:8080/admin/detail",
            {
                method: 'POST',
                credentials: "include",
                headers: {
                    'Content-Type': 'application/json; charset=utf-8',
                },
            })
            .then(resp => resp.json())
            .then(data => {
                console.log(data)
                fetched = true

                let temp = []
                if (data.ended_voting_events) data.ended_voting_events.forEach(({name, id})=>{
                    temp.push({eventName: name, id: id});
                })
                setrows(temp)
            })
            .catch(error => alert(error.toString()))
    })

    return (
        <div>
            <Table className={classes.table}>
                <TableHead>
                    <TableRow>
                        <TableCell>Event Name</TableCell>
                        <TableCell>Result</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {rows.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map((row) => (
                        <TableRow key={row.id}>
                        <TableCell style={{ width: 250 }}>{row.eventName}</TableCell>
                        <TableCell >
                            <Link href={`/Event_result?eventId=${row.id}`}>
                                <Button 
                                    type="submit"
                                    fullWidth
                                    variant="outlined"
                                    >
                                    View Result
                                </Button>
                            </Link>
                        </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
            <TablePagination
                count={rows.length}
                rowsPerPage={rowsPerPage}
                rowsPerPageOptions={rowsPerPage}
                page={page}
                onPageChange={handleChangePage}
                className={classes.pagination}
            />
        </div>
    );
}
  
export default PastVotingTable;