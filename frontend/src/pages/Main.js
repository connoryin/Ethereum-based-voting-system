import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import LoginForm from '../components/LoginForm';
import image from "../images/LoginBackground.png"; 

const useStyles = makeStyles(theme => ({
  main: {
    display: 'flex',
    backgroundImage:`url(${image})`,
    backgroundPosition: 'center',
    backgroundSize: 'cover',
    backgroundRepeat: 'no-repeat',
    width: '1530px',
    height: '710px'
  },
}));

function Main() {
    const classes = useStyles();
  
    return (
      <div className={classes.main}>
        <LoginForm />
      </div>
    );
  }
  
  export default Main;