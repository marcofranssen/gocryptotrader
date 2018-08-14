import React, { Component } from 'react';
import PropTypes from 'prop-types';

class CandlestickChart extends Component {
  constructor(props) {
    super(props);
    this.state = {
      data: props.data
    };
  }

  render() {
    const { data } = this.state;
    return <pre>{JSON.stringify(data, null, 2)}</pre>;
  }
}

CandlestickChart.propTypes = {
  data: PropTypes.array.isRequired,
  width: PropTypes.number.isRequired,
  type: PropTypes.oneOf(['hybrid', 'svg']).isRequired
};

CandlestickChart.defaultProps = {
  type: 'hybrid'
};

export default CandlestickChart;
