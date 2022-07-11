import React, {useState} from "react";
import {makeStyles} from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import CssBaseline from '@material-ui/core/CssBaseline';
import Container from '@material-ui/core/Container';
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
        margin: theme.spacing(2, 4.5, 0),
        width: '30%'
    },
    error: {
        color: "red",
        textAlign: "center"
    }
}));

function AdminLoginForm() {
    const classes = useStyles();
    const [ID, setID] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    // check if the admin ID is valid
    const isValidID = true;
    const history = useHistory()

    function login() {
        fetch(
            "http://localhost:8080/admin_login",
            {
                method: 'POST',
                credentials: "include",
                headers: {
                    'Content-Type': 'application/json; charset=utf-8',
                },
                body: JSON.stringify({
                    name: ID,
                    password: password,
                })
            })
            .then(resp => resp.json())
            .then(data => {
                console.log(data)
                if (data.status_code === 1) {
                    history.push({
                        pathname: "/Admin_homepage",
                        search: "?id=" + data.admin_id,
                    });
                } else {
                    setError("Wrong username or password")
                }

            })
            .catch(error => setError("Wrong username or password"))

    }

    function register() {
        fetch(
            "http://localhost:8080/admin_register",
            {
                method: 'POST',
                credentials: "include",
                headers: {
                    'Content-Type': 'application/json; charset=utf-8',
                },
                body: JSON.stringify({
                    name: ID,
                    password: password,
                })
            })
            .then(resp => resp.json())
            .then(data => {
                console.log(data)
                if (data.is_success) {
                    login()
                } else {
                    setError("Registration failed")
                }
            })
            .catch(error => alert(error.toString()))

    }

    return (
        <Container component="main" maxWidth="xs">
            <CssBaseline/>
            <div className={classes.paper}>
                <form className={classes.form} noValidate>
                    {isValidID
                        ? <TextField
                            variant="outlined"
                            margin="normal"
                            required
                            fullWidth
                            label="Admin ID"
                            name="Admin ID"
                            autoFocus
                            onChange={e => setID(e.target.value)}
                        />
                        : <TextField
                            error
                            variant="outlined"
                            margin="normal"
                            required
                            fullWidth
                            label="Admin ID"
                            name="Admin ID"
                            autoFocus
                            onChange={e => setID(e.target.value)}
                            helperText="Invalid ID. Please enter again."
                        />
                    }
                    <TextField
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        label="Password"
                        name="Password"
                        autoFocus
                        onChange={e => setPassword(e.target.value)}
                    />
                    {/*<Link href='/Admin_homepage'>*/}
                    <Button
                        fullWidth
                        variant="outlined"
                        color="primary"
                        className={classes.submit}
                        onClick={login}
                    >
                        Login
                    </Button>
                    {/*</Link>*/}
                    {/*<Link href='/Admin_homepage'>*/}
                        <Button
                            fullWidth
                            variant="outlined"
                            color="primary"
                            className={classes.submit}
                            onClick={register}
                        >
                            Register
                        </Button>
                    {/*</Link>*/}
                    <p className={classes.error}>{error}</p>
                </form>
            </div>
        </Container>
    );
}

export default AdminLoginForm;
