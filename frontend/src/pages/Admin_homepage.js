import React, {useState, useEffect} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import image from "../images/AdminHomepageBackground.png";
import Box from '@mui/material/Box';
import Button from '@material-ui/core/Button';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import PropTypes from 'prop-types';
import CurrentVotingTable from '../components/CurrentVotingTable';
import PastVotingTable from '../components/PastVotingTable';
import Link from '@material-ui/core/Link';
import {useHistory} from "react-router-dom";
import qs from "qs";

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
  tabPanelBox: {
    position: 'absolute',
    left: '530px',
    top: '100px',
    width: '850px',
    height: '420px',
  },
  tabBox: {
    position: 'absolute',
    top: '270px',
    left: '150px',
  },
  tab: {
    height: '300px',
    textAlign: 'center',
    fontSize: '28px',
    color: '#3e637d',
    minHeight: '80px',
    marginBottom: '40px',
    borderTopWidth: 0
  },
  addVotebtn: {
    position: 'absolute',
    left: '1120px',
    top: '595px',
    width: '350px',
    height: '65px',
    textAlign: 'center',
    fontSize: '27px',
    color: '#3e637d'
  }, 
}));

// For tabs
function TabPanel(props) {
  const { children, value, index, ...other } = props;
 
  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`tabpanel-${index}`}
      aria-labelledby={`tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box >
          {children}
        </Box>
      )}
    </div>
  );
}

TabPanel.propTypes = {
  children: PropTypes.node,
  index: PropTypes.any.isRequired,
  value: PropTypes.any.isRequired,
};

function a11yProps(index) {
  return {
    id: `tab-${index}`,
    'aria-controls': `tabpanel-${index}`,
  };
}

function Admin_homepage(props) {
  const classes = useStyles();
  const history = useHistory()

  const [value, setValue] = React.useState(0);
  const handleChange = (event, newValue) => {
    setValue(newValue);
  };

    return (
      <div className={classes.bg}>
        <Box className={classes.tabPanelBox}>
          <TabPanel value={value} index={0}>
            <CurrentVotingTable/>
          </TabPanel>
          <TabPanel value={value} index={1}>
            <PastVotingTable />
          </TabPanel>
        </Box>
        <Box className={classes.tabBox} >
          <Tabs 
            value={value} 
            variant="fullWidth"
            onChange={handleChange} 
            orientation="vertical"
            TabIndicatorProps={{
              style: {
                  display: "none",
              },
            }}
          >
            <Tab
              label="View Ongoing Elections"
              className={classes.tab}
              {...a11yProps(0)}
              >
            </Tab>       
            <Tab
              label="View Past Elections"
              className={classes.tab}
              {...a11yProps(1)}
              >
            </Tab>
          </Tabs>
        </Box>
        {/*<Link href='/Create'>*/}
          <Button
            // type="submit"
            fullWidth
            variant="text"
            className={classes.addVotebtn}
            onClick={()=>{
              let params = qs.parse(props.location.search)
              history.push({
                pathname: "/Create",
                search: "?id=" + params['?id'],
              });
            }}
            >
            Start A New Election
          </Button>
        {/*</Link>*/}
      </div>
    );
  }
  
  export default Admin_homepage;