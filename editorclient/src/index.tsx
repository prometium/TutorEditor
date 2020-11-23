import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { register } from './serviceWorker';
import App from './App';
import { loadState } from 'src/localStorage';
import configureStore from './configureStore';
import { ApplicationState } from './store';

import 'normalize.css';
import 'src/theme/theme.scss';

const preloadedState: ApplicationState = (loadState() as unknown) as ApplicationState;
const store = configureStore(preloadedState);

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);

register();
