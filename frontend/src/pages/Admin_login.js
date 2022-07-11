import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import image from "../images/AdminLoginBackground.png"; 
import AdminLoginForm from '../components/AdminLoginForm';

const useStyles = makeStyles(theme => ({
  bg: {
    display: 'flex',
    backgroundImage:`url(${image})`,
    backgroundPosition: 'center',
    backgroundSize: 'cover',
    backgroundRepeat: 'no-repeat',
    width: '1530px',
    height: '710px'
  },
  center: {
    marginLeft: '40%'
  }
}));

function Admin_login() {
    const classes = useStyles();

    return (
        <div className={classes.bg}>
          <div className={classes.center}>
            <AdminLoginForm />
          </div>
        </div>
    );
}
  
export default Admin_login;