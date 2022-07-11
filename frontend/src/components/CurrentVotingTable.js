import React, {useState, useEffect} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import TablePagination from '@mui/material/TablePagination';
import ProgressBar from "@ramonak/react-progress-bar";

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

function endEvent(id) {
    fetch(
        "http://localhost:8080/admin/end_event",
        {
            method: 'POST',
            credentials: "include",
            headers: {
                'Content-Type': 'application/json; charset=utf-8',
            },
            body: JSON.stringify({
                event_id: id
            })
        })
        .then(resp => resp.json())
        .then(data => {
            console.log(data)
            window.location.reload(false);
        })
        .catch(error => alert(error.toString()))
}
let fetched = false

function CurrentVotingTable() {
    const [rows, setrows] = React.useState([]);

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
                if (data.in_voting_events) data.in_voting_events.forEach(({name, received_vote_num, total_vote_num, id})=>{
                    temp.push({eventName: name, completion: 100*received_vote_num/total_vote_num, id: id});
                })
                setrows(temp)
            })
            .catch(error => alert(error.toString()))
    })

    const classes = useStyles();
    const [page, setPage] = React.useState(0);
    const rowsPerPage = 5;
    const handleChangePage = (event, newPage) => {
        setPage(newPage);
    };
    return (
        <div>
            <Table className={classes.table}>
                <TableHead>
                    <TableRow>
                        <TableCell>Event Name</TableCell>
                        <TableCell>Completion</TableCell>
                        <TableCell>End Event</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {rows.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map((row) => (
                        <TableRow key={row.id}>
                        <TableCell style={{ width: 250 }}>{row.eventName}</TableCell>
                        <TableCell style={{ width: 300 }}>
                            <ProgressBar completed={row.completion} bgColor="#2596be"/>
                        </TableCell>
                        <TableCell >
                            <Button 
                                type="submit"
                                fullWidth
                                variant="outlined"
                                onClick={()=>endEvent(row.id)}
                                >
                                End Event
                            </Button>
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
  
export default CurrentVotingTable;