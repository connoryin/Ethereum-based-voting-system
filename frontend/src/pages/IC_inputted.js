import React, {useState, useEffect } from 'react';
import Checkbox from '@mui/material/Checkbox';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import {makeStyles} from '@material-ui/core/styles';
import image from "../images/BeforeVoteBackground.png";
import image2 from "../images/VotedBackground.png";
import Button from '@material-ui/core/Button';
import CanvasJSReact from '../charts/canvasjs.react.js';
import qs from "qs";

const useStyles = makeStyles((theme) => ({
    main: {
        display: 'flex',
        backgroundImage: `url(${image})`,
        backgroundPosition: 'center',
        backgroundSize: 'cover',
        backgroundRepeat: 'no-repeat',
        width: '1530px',
        height: '710px'
    },
    descriptionBox: {
        marginTop: '10%',
        marginLeft: '20%',
        marginRight: '10%',
        width: '600px',
        height: '400px',
    },
    optionBox: {
        marginTop: '15%',
        width: '330px',
        height: '350px',
    },
    title: {
        textAlign: "center",
        marginTop: '20%',
    },
    main2: {
        display: 'flex',
        backgroundImage: `url(${image2})`,
        backgroundPosition: 'center',
        backgroundSize: 'cover',
        backgroundRepeat: 'no-repeat',
        width: '1530px',
        height: '710px'
    },
    myVoteBox: {
        marginTop: '12%',
        marginLeft: '15%',
        marginRight: '10%',
        width: '400px',
        height: '400px',
        textAlign: "center"
    },
    pieChartBox: {
        position: 'absolute',
        left: '830px',
        top: '200px',
        width: '500px',
        height: '400px',
    }
}));

function IC_inputted(props) {
    const classes = useStyles();
    const [eventName, seteventName] = useState('');
    const [description, setdescription] = useState("");


    const [maxVoteNum, setmaxVoteNum] = useState(3);
    const [isVoted, setIsVoted] = useState('');
    const [candidateList, setcandidateList] = useState([]);

    // let selfVote = [];
    const [selfVote, setselfVote] = useState([]);

    const [eventEnded, seteventEnded] = useState(false);
    const [error, seterror] = useState('');

    var CanvasJSChart = CanvasJSReact.CanvasJSChart;
    let data = [];
    const [pieChart, setpieChart] = useState({
        animationEnabled: true,
        backgroundColor: null,
        data: [{
            type: "pie",
            indexLabelFontSize: 10,
            radius: 110,
            toolTipContent: '{label}: #percent%',
            dataPoints: data
        }]
    });

    let checkedList = []
    let params = qs.parse(props.location.search)
    const eventId = params['?eventId']
    const IC = params.invcode

    useEffect(() => {
        if (description) return
        fetch(
            "http://localhost:8080/user/vote-details/",
            {
                method: 'POST',
                credentials: "same-origin",
                headers: {
                    'Content-Type': 'application/json; charset=utf-8',
                },
                body: JSON.stringify({
                    inv_code: IC,
                })
            })
            .then(resp => {
                return resp.json()
            })
            .then(data => {
                console.log(data)
                let globalResults = [];
                let candidates = []
                // pieData = []
                setdescription(data.event.description)
                seteventName(data.event.name)
                seteventEnded(data.event.is_end)

                for (let candidate of data.event.candidates) {
                    candidates.push(candidate.name)
                }
                setcandidateList(candidates)
                setmaxVoteNum(data.event.max_vote_num_per_person)

                setIsVoted(data.is_voted)
                if (data.is_voted) {
                    let temp = []
                    data.self_vote.map((cand)=>temp.push(cand.name))
                    setselfVote(temp)
                    for (let candidate of data.event.candidates) {
                        globalResults.push({
                            name: candidate.name,
                            num: candidate.vote_num_get
                        })
                    }
                    let tempPie = []
                    for (var i = 0; i < globalResults.length; i++) {
                        tempPie.push({y: globalResults[i].num, label: globalResults[i].name});
                    }
                    console.log(tempPie)
                    setpieChart({
                        animationEnabled: true,
                        backgroundColor: null,
                        data: [{
                            type: "pie",
                            indexLabelFontSize: 10,
                            radius: 110,
                            toolTipContent: '{label}: #percent%',
                            dataPoints: tempPie
                        }]
                    })
                }

            })
            .catch(error => alert(error.toString()))

    });

    function refreshPage() {
        setIsVoted('true');
        setselfVote(checkedList)

        fetch(
            "http://localhost:8080/user/vote/",
            {
                method: 'POST',
                credentials: "same-origin",
                headers: {
                    'Content-Type': 'application/json; charset=utf-8',
                },
                body: JSON.stringify({
                    inv_code: IC,
                    event_id: parseInt(eventId),
                    voted_candidate_names: checkedList
                })
            })
            .then(resp => {
                return resp.json()
            })
            .then(data => {
                console.log(data)
                if (!data.is_success) seterror("Vote submission failed")
            })
            .catch(error => alert(error.toString()))

    }

    function handler(e) {
        // console.log(e)
        if (e.target.checked) {
            if (checkedList.length >= maxVoteNum) {
                e.target.checked = false;
            } else {
                checkedList.push(e.target.name);
            }  
        } else {
            checkedList = checkedList.filter(function (value) {
                return value !== e.target.name;
            });
        }
    }

    if (isVoted) {
        return (
            <div className={classes.main2}>
                <Box className={classes.myVoteBox}>
                    <Typography component="h1" variant="h4">
                        Your Vote(s)
                    </Typography>
                    {selfVote.map((name) => {
                        return (
                            <Typography component="h1" variant="h5" sx={{m: 2}}>
                                {name}
                            </Typography>
                        );
                    })}
                </Box>
                <Box className={classes.pieChartBox} >
                    {eventEnded
                        ?
                        <div>
                            <CanvasJSChart options = {pieChart} />
                        </div>
                        :
                        <Typography component="h1" variant="h5" sx={{m: 2}}>
                            Result will display after the event is ended.
                        </Typography>
                    }
                </Box>
            </div>
        );
    } else {
        return (
            <div className={classes.main}>
                <Box className={classes.descriptionBox}>
                    <Typography component="h1" variant="h3" className={classes.title}>
                        {eventName}
                    </Typography>
                    <hr/>
                    {description}
                </Box>
                <Box className={classes.optionBox}>
                    {eventEnded
                        ?
                        <div>
                            <Typography component="h1" variant="h7" className={classes.title} >
                                The election is ended
                            </Typography>
                            <Typography component="h1" variant="h5" className={classes.title} color="#3a5062">
                                Please select {maxVoteNum} candidate(s).
                            </Typography>
                        </div>
                        : 
                        <Typography component="h1" variant="h5" className={classes.title}>
                            Please select {maxVoteNum} candidate(s).
                        </Typography>
                    }
                    {candidateList.map((candidateName, index) => {
                        return (
                            <label>
                                <Checkbox 
                                    name={candidateName} 
                                    value={candidateName} 
                                    onClick={e => handler(e)}
                                    disabled={eventEnded}
                                />
                                    {candidateName}
                            </label>
                        );
                    })}
                    <Button
                        type="submit"
                        fullWidth
                        variant="outlined"
                        className={classes.submit}
                        onClick={refreshPage}
                        disabled={eventEnded}
                    >
                        Submit
                    </Button>
                    <p>{error}</p>
                </Box>
            </div>
        );
    }
}

export default IC_inputted;
