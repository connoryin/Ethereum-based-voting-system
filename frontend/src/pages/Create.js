import React, {useState} from 'react';
import {makeStyles} from '@material-ui/core/styles';
import image from "../images/CreateEventBackground.png";
import Box from '@mui/material/Box';
import MenuItem from '@mui/material/MenuItem';
import TextField from '@material-ui/core/TextField';
import {IoIosAddCircle} from "react-icons/io";
import Button from '@material-ui/core/Button';
import Link from '@material-ui/core/Link';
import qs from "qs";
import {useHistory} from "react-router-dom";

const papa = require('papaparse');

const useStyles = makeStyles(theme => ({
    bg: {
        display: 'flex',
        backgroundImage: `url(${image})`,
        backgroundPosition: 'center',
        backgroundSize: 'cover',
        backgroundRepeat: 'no-repeat',
        width: '1530px',
        height: '710px'
    },
    addBtn: {
        position: 'absolute',
        left: '780px',
        top: '570px',
        width: 50,
        height: 50,
    },
    candidateBox: {
        position: 'absolute',
        left: '720px',
        top: '210px',
        width: '220px',
        height: '350px',
        overflowX: 'hidden',
        overflowY: 'scroll'
    },
    optionsText: {
        position: 'absolute',
        left: '760px',
        top: '170px',
        fontSize: '30px',
        color: '#3e637d',
    },
    text: {
        fontSize: '30px',
        color: '#e88b6a',
    },
    box: {
        width: '310px',
        height: '260px',
        lineHeight: "100px",
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    p1: {
        position: 'absolute',
        left: '220px',
        top: '110px',
    },
    p2: {
        position: 'absolute',
        left: '220px',
        top: '390px',
    },
    p3: {
        position: 'absolute',
        left: '1075px',
        top: '110px',
    },
    p4: {
        position: 'absolute',
        left: '1075px',
        top: '390px',
    },
    input: {
        width: '100px'
    },
    createBtn: {
        position: 'absolute',
        left: '1270px',
        top: '610px',
        width: '120px',
        height: '60px',
        textAlign: 'center',
        fontSize: '27px',
        color: '#3e637d'
    },
}));

const Input = () => {
    return <TextField margin="normal"/>;
};

let candidates = []
let voters = []

function Create(props) {
    const classes = useStyles();
    let history = useHistory()

    const handleOptionList = (e) => {
        let index = parseInt(e.target.name)

        while (candidates.length <= index) candidates.push({
            "name": "",
            "description": "niu"
        })

        candidates[index].name = e.target.value
        candidates[index].id = index
        candidates[index].vote_num_get = 0
        candidates[index].is_winner = false

        console.log(candidates)
    }

    // for add options
    const [inputList, setInputList] = useState([<TextField margin="normal" name={0} onInput={handleOptionList}/>]);

    const onAddBtnClick = event => {
        setInputList(inputList.concat(<TextField margin="normal" name={inputList.length} onInput={handleOptionList}/>));
    };

    // for max votes selection
    const maxVotesList = [];
    for (var i = 1; i <= inputList.length; i++) {
        maxVotesList.push(i);
    }
    const [maxVotes, setMaxVotes] = React.useState(1);
    const handleChange = (event) => {
        setMaxVotes(event.target.value);
    };

    let params = qs.parse(props.location.search)

    const [name, setname] = React.useState("");
    const handleNameChange = (e) => {
        setname(e.target.value)
    }

    const [description, setdescription] = React.useState("");
    const handleDescriptionChange = (e) => {
        setdescription(e.target.value)
    }


    function handleFileSelected(e) {
        papa.parse(e.target.files[0], {
            complete: function (results) {
                voters = results.data[0]
                console.log(voters)
            }
        });
    }

    function handleSubmit(e) {
        fetch(
            "http://localhost:8080/admin/create_event",
            {
                method: 'POST',
                credentials: "include",
                headers: {
                    'Content-Type': 'application/json; charset=utf-8',
                },
                body: JSON.stringify({
                    "event": {
                        "admin_id": parseInt(params['?id']),
                        "name": name,
                        "description": description,
                        "max_vote_num_per_person": parseInt(maxVotes),
                        "candidates": candidates,
                        "total_vote_num": voters.length
                    },
                    "voters": voters
                })
            })
            .then(resp => {
                return resp.json()
            })
            .then(res => {
                console.log(res)
                if (res.is_success) {
                    alert("Successfully created the event")
                    history.push({
                        pathname: "/Admin_homepage",
                        search: "?id=" + params['?id'],
                    });
                }
                else alert("Event creation fails")
            })
            .catch(error => alert(error.toString()))
    }

    return (
        <div className={classes.bg}>
            <div className={classes.p1}>
                <Box className={classes.box}>
                    <div className={classes.text}>
                        Event Name
                    </div>
                    <TextField
                        fullWidth
                        onInput={handleNameChange}
                    />
                </Box>
            </div>
            <div className={classes.p2}>
                <Box className={classes.box}>
                    <div className={classes.text}>
                        Event Description
                    </div>
                    <TextField
                        fullWidth
                        multiline
                        maxRows={5}
                        onInput={handleDescriptionChange}
                    />
                </Box>
            </div>
            <div className={classes.p3}>
                <Box className={classes.box}>
                    <div className={classes.text}>
                        Voter List
                    </div>
                    <input type="file" onChange={handleFileSelected}/>
                </Box>
            </div>
            <div className={classes.p4}>
                <Box className={classes.box}>
                    <div className={classes.text}>
                        Max Votes
                    </div>
                    <TextField
                        select
                        value={maxVotes}
                        onChange={handleChange}
                    >
                        {maxVotesList.map((option) => (
                            <MenuItem value={option}>
                                {option}
                            </MenuItem>
                        ))}
                    </TextField>
                </Box>
            </div>
            <div className={classes.optionsText}>
                Options
            </div>
            <div className={classes.candidateBox}>
                {inputList}
            </div>
            <IoIosAddCircle className={classes.addBtn} color="#9dc1d0" onClick={onAddBtnClick}/>
            {/*<Link href='/Admin_homepage'>*/}
                <Button
                    // type="submit"
                    fullWidth
                    variant="text"
                    className={classes.createBtn}
                    onClick={handleSubmit}
                >
                    Create
                </Button>
            {/*</Link>*/}
        </div>
    );
}

export default Create;