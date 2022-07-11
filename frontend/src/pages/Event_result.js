import React, {useEffect, useState} from 'react';
import {makeStyles} from '@material-ui/core/styles';
import image from "../images/EventResultBackground.png";
import Button from '@material-ui/core/Button';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import CanvasJSReact from '../charts/canvasjs.react.js';

var qs = require('qs');


var CanvasJS = CanvasJSReact.CanvasJS;
var CanvasJSChart = CanvasJSReact.CanvasJSChart;

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
    title: {
        marginBottom: '10px',
        fontSize: '20px',
        color: '#e88b6a',
    },
    box: {
        width: '340px',
        height: '450px',
        lineHeight: "100px",
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    p1: {
        position: 'absolute',
        left: '160px',
        top: '130px',
    },
    downloadBtn: {
        position: 'absolute',
        left: '1180px',
        top: '620px',
        width: '220px',
        height: '70px',
        textAlign: 'center',
        fontSize: '27px',
        color: '#3e637d'
    },
    pieChartPosition: {
        position: 'absolute',
        left: '450px',
        top: '150px',
    },
    barChartPosition: {
        position: 'absolute',
        left: '950px',
        top: '150px',
        width: '400px',
    },
    winnerTitle: {
        marginTop: '20px',
        fontSize: '20px',
        color: '#e88b6a',
    },
    description: {
        fontSize: '15px',
        color: "#000000"
    }
}));


let winnerList = [];

let candidateList = [];
let votesList = [];
let data = [];

function Event_result(props) {
    let params = qs.parse(props.location.search)
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

    const [barChart, setbarChart] = useState({
        animationEnabled: true,
        backgroundColor: null,
        axisX: {
            labelFontSize: 10,
        },
        axisY: {
            labelFontSize: 10,
        },
        data: [{
            indexLabelFontSize: 10,
            toolTipContent: '{label}: {y} votes',
            indexLabel: "{y}",
            dataPoints: data
        }]
    });

    const [eventName, seteventName] = useState("");
    const [description, setdescription] = useState("");

    function handleDownload(e) {
        fetch(
            "http://localhost:8080/admin/download",
            {
                method: 'POST',
                credentials: "include",
                headers: {
                    'Content-Type': 'application/json; charset=utf-8',
                },
                body: JSON.stringify({
                    event_id: parseInt(params["?eventId"]),
                })
            })
            .then(resp => {
                return resp.json()
            })
            .then(res => {
                console.log(res.invitation_code_2_details)
                const a = document.createElement("a");
                const file = new Blob([JSON.stringify(res.invitation_code_2_details, null, "\t")], {type: 'application/json'});                a.href = URL.createObjectURL(file);
                a.download = eventName+ "_results.txt";
                a.click();
            })
            .catch(error => alert(error.toString()))
    }

    useEffect(() => {
        winnerList = []
        candidateList = [];
        votesList = [];
        data = [];

        fetch(
            "http://localhost:8080/admin/get_event",
            {
                method: 'POST',
                credentials: "include",
                headers: {
                    'Content-Type': 'application/json; charset=utf-8',
                },
                body: JSON.stringify({
                    event_id: parseInt(params["?eventId"]),
                })
            })
            .then(resp => {
                return resp.json()
            })
            .then(res => {
                console.log(res)
                seteventName(res.event.name)
                setdescription(res.event.description)
                for (let candidate of res.event.candidates) {
                    candidateList.push(candidate.name)
                    votesList.push(candidate.vote_num_get)
                    if (candidate.is_winner) winnerList.push(candidate.name)
                }
                for (var i = 0; i < votesList.length; i++) {
                    data.push({y: votesList[i], label: candidateList[i]});
                }
                console.log(data)
                // pie chart
                setpieChart({
                    animationEnabled: true,
                    backgroundColor: null,
                    data: [{
                        type: "pie",
                        indexLabelFontSize: 10,
                        radius: 110,
                        toolTipContent: '{label}: #percent%',
                        dataPoints: data
                    }]
                })

// bar chart
                setbarChart({
                    animationEnabled: true,
                    backgroundColor: null,
                    axisX: {
                        labelFontSize: 10,
                    },
                    axisY: {
                        labelFontSize: 10,
                    },
                    data: [{
                        indexLabelFontSize: 10,
                        toolTipContent: '{label}: {y} votes',
                        indexLabel: "{y}",
                        dataPoints: data
                    }]
                })
            })
            .catch(error => alert(error.toString()))
    }, [])

    const classes = useStyles();

    return (
        <div className={classes.bg}>
            <div className={classes.p1}>
                <Box className={classes.box}>
                    <div className={classes.title}>
                        {eventName}
                    </div>
                    <div className={classes.description}>
                        {description}
                    </div>
                    <div className={classes.winnerTitle}>
                        Winner
                    </div>
                    <div>
                        {winnerList.map((option) => (
                            <Typography variant="body1">
                                {option}
                            </Typography>
                        ))}
                    </div>
                </Box>
            </div>
            <div className={classes.pieChartPosition}>
                <CanvasJSChart options={pieChart}/>
            </div>
            <div className={classes.barChartPosition}>
                <CanvasJSChart options={barChart}/>
            </div>
            <Button
                // type="submit"
                fullWidth
                variant="text"
                className={classes.downloadBtn}
                onClick={handleDownload}
            >
                Download
            </Button>
        </div>
    );
}

export default Event_result;
