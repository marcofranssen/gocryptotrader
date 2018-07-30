import React from 'react';
import ReactDOM from 'react-dom';
import ExchangesSettings from './ExchangesSettings';

it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDOM.render(<ExchangesSettings data={[]} />, div);
  ReactDOM.unmountComponentAtNode(div);
});
