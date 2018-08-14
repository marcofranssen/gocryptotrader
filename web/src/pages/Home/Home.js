import React, { Component, Fragment } from 'react';
import PropTypes from 'prop-types';
import { Paper, Typography } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import { pageStyles } from '../styles';
// import logo from './logo.svg';
import './Home.css';
import logo from '../../assets/images/gctlogo-diff.svg';
import { EmptyState, CandlestickChart, withFetching } from '../../components';
import { tsvParse } from 'd3-dsv';
import { timeParse } from 'd3-time-format';

const styles = theme => ({
  ...pageStyles(theme),
  logo: {
    animation: 'App-logo-wiggle infinite 5s linear',
    height: '200px'
  }
});

const dateParser = timeParse('%Y-%m-%d');

const parseData = parse => d => {
  d.date = parse(d.date);
  d.open = +d.open;
  d.high = +d.high;
  d.low = +d.low;
  d.close = +d.close;
  d.volume = +d.volume;

  return d;
};

const tsvParser = async response => {
  const data = await response.text();
  return tsvParse(data, parseData(dateParser));
};

class Home extends Component {
  render() {
    const { classes, data, error, isLoading } = this.props;

    return (
      <div className={classes.root}>
        <Typography variant="title" gutterBottom>
          Dashboard
        </Typography>
        <Paper className={classes.paper}>
          {!data || error || isLoading ? (
            <Fragment>
              <Typography>
                Crypto Trader is a crypto trading bot with back testing support
                and support for a bunch of popular exchanges.
              </Typography>
              <img src={logo} className={classes.logo} alt="logo" />
              <EmptyState data={data} error={error} isLoading={isLoading} />
            </Fragment>
          ) : (
            <CandlestickChart type="hybrid" data={data} />
          )}
        </Paper>
      </div>
    );
  }
}

Home.propTypes = {
  classes: PropTypes.object.isRequired,
  theme: PropTypes.object.isRequired,
  isLoading: PropTypes.bool.isRequired,
  data: PropTypes.array,
  error: PropTypes.object
};

export default withFetching(
  'https://rrag.github.io/react-stockcharts/data/MSFT.tsv',
  { parser: tsvParser }
)(withStyles(styles, { withTheme: true })(Home));
