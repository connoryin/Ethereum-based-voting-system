import React, {useState} from "react";
import {makeStyles} from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import CssBaseline from '@material-ui/core/CssBaseline';
import Container from '@material-ui/core/Container';
import Typography from '@mui/material/Typography';
import {useHistory} from "react-router-dom";

const useStyles = makeStyles((theme) => ({
    paper: {
        marginTop: '50%',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    form: {
        width: '100%',
        marginTop: theme.spacing(1),
    },
    submit: {
        margin: theme.spacing(2, 0, 3),
    },
    error: {
        color: "red",
        textAlign: "center"
    }
}));

function LoginForm() {
    const classes = useStyles();
    const [IC, setIC] = useState("");
    const [error, setError] = useState("");
    const isValidIC = true;
    const history = useHistory()

    // check if invite code is valid
    function checkIC() {
        fetch(
            "http://localhost:8080/",
            {
                method: 'POST',
                credentials: "same-origin",
                headers: {
                    'Content-Type': 'application/json; charset=utf-8',
                },
                body: JSON.stringify({
                    inv_code: IC
                })
            })
            .then(resp => {
                return resp.json()
            })
            .then(data => {
                switch (data.user_status) {
                    case 0:
                        setError("Invalid Invitation code")
                        break;
                    default:
                        history.push({
                            pathname: "/IC_inputted",
                            search: "?eventId="+data.event_id+"&invcode="+IC,
                        });
                        break;

                }
            })
            .catch(error => setError("Invalid Invitation code"))
    }

    return (
        <Container component="main" maxWidth="xs">
            <CssBaseline/>
            <div className={classes.paper}>
                <Typography component="h1" variant="h6">
                    CS6675 Voting App
                </Typography>
                <form className={classes.form} noValidate>
                    {isValidIC
                        ? <TextField
                            variant="outlined"
                            margin="normal"
                            required
                            fullWidth
                            label="Invite Code"
                            autoFocus
                            onChange={e => setIC(e.target.value)}
                        />
                        : <TextField
                            error
                            variant="outlined"
                            margin="normal"
                            required
                            fullWidth
                            label="Invite Code"
                            autoFocus
                            onChange={e => setIC(e.target.value)}
                            helperText="Invalid invite code. Please enter again."
                        />
                    }
                    {/*<Link href='/IC_inputted' variant="body2">*/}
                    <Button
                        fullWidth
                        variant="contained"
                        color="primary"
                        className={classes.submit}
                        onClick={checkIC}
                    >
                        Login
                    </Button>
                    {/*</Link>*/}
                    <Link href="/Admin_login" variant="body2">
                        Login As Admin
                    </Link>
                    <p className={classes.error}>{error}</p>
                </form>
            </div>
        </Container>
    );
}

export default LoginForm;
