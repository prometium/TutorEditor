import React from 'react';
import ReactDOM from 'react-dom';
import { register } from './serviceWorker';
import App from './App';

import 'normalize.css';

ReactDOM.render(<App />, document.getElementById('root'));

register();
