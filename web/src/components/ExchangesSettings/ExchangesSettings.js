import React, { Component } from 'react';
import PropTypes from 'prop-types';
import classNames from 'classnames';
import { withStyles } from '@material-ui/core/styles';
import {
  Table, TableRow, TableHead, TableCell, TableBody,
  Switch,
} from '@material-ui/core';
import { SecretInput } from '../';

const styles = theme => ({
  root: {
    overflowX: 'scroll',
  },
  table: {
    minWidth: 700,
  },
  tableCell: {
    padding: theme.spacing.unit / 2,
  },
  small: {
    width: 75,
  }
});

const tableCellType = (type) => {
  switch (type) {
    case 'number':
      return { numeric: true };
    default:
      return {};
  }
};
const renderValue = (column, value, rowIndex) => {
  const valueType = typeof value;

  if (['apiKey', 'apiSecret', 'clientID'].includes(column)) {
    return (
      <SecretInput
        id={column + rowIndex}
        value={value}
      />
    );
  }

  if (valueType === 'boolean') {
    return (
      <Switch
        checked={value}
        // onChange={this.handleChange('WebsocketAllowInsecureOrigin')}
      />
    );
  }

  if (['enabledPairs', 'baseCurrencies'].includes(column)) {
    return value.replace(/,/g, ', ');
  }

  return valueType === 'object' ? JSON.stringify(value, null, 2) : value;
};

const splitPascalCase = (str) => str.replace(/([A-Z]+)([A-Z][a-z])/g, ' $1 $2')
  .replace(/([a-z\d])([A-Z])/g, '$1 $2')
  .replace(/([a-zA-Z])(\d)/g, '$1 $2')
  .replace(/^./, function(str){ return str.toUpperCase(); })
  .trim();

const ExchangeTableRow = ({ classes, rowData, rowIndex, ...props }) => (
  <TableRow>
    {
      rowData.map((column, columnIndex) => (
        <TableCell
          key={columnIndex} {...tableCellType(typeof column[1])}
          className={classNames(classes.tableCell, classes.small)}
        >
          {renderValue(column[0], column[1], rowIndex)}
        </TableCell>
      ))
    }
  </TableRow>
);

ExchangeTableRow.propTypes = {
  classes: PropTypes.object.isRequired,
  rowData: PropTypes.array.isRequired,
  rowIndex: PropTypes.number.isRequired,
};

const ExchangeTable = ({ classes, columns, data, ...props }) => (
  <Table className={classes.table}>
    <TableHead>
      <TableRow>
        { columns.map((column, index) => <TableCell key={index}>{column}</TableCell>) }
      </TableRow>
    </TableHead>
    <TableBody>
      {data.map((row, rowIndex) => <ExchangeTableRow key={rowIndex} rowIndex={rowIndex} rowData={row} classes={classes} />)}
    </TableBody>
  </Table>
);

class ExchangesSettings extends Component {
  state = {
    headers: [],
    body: []
  }

  componentDidMount() {
    const { data } = this.props;
    if (!data || data.length === 0) {
      return;
    }

    const filterColumns = column => !['availablePairs'].includes(column[0]);
    const body = data.map(exchange => {
      const entries = Object.entries(exchange).filter(filterColumns);
      if (!exchange.hasOwnProperty('clientID')) {
        entries.splice(10, 0, ['clientID', 'N.A.']);
      }
      return entries;
    });

    const headers = body[0].map(column => splitPascalCase(column[0]));

    this.setState({ headers, body });
  }

  render() {
    const { classes } = this.props;
    const { headers, body } = this.state;

    return (
      <div className={classes.root}>
        <ExchangeTable classes={classes} columns={headers} data={body} />
      </div>
    );
  }
}

ExchangesSettings.propTypes = {
  classes: PropTypes.object.isRequired,
  theme: PropTypes.object.isRequired,
  data: PropTypes.array.isRequired
};

export default withStyles(styles, { withTheme: true })(ExchangesSettings);
