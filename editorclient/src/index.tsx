import React from 'react';
import ReactDOM from 'react-dom';
import { register } from './serviceWorker';
import App from './App';

import 'normalize.css';
import 'src/theme/theme.scss'

ReactDOM.render(<App />, document.getElementById('root'));

register();
