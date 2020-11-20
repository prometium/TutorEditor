import React, { Suspense } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Login from 'src/components/routes/Login';

function App(): JSX.Element {
  return (
    <Router>
      <Suspense fallback="loading....">
        <Switch>
          <Route exact path="/">
            root
          </Route>
          <Route exact path="/login">
            <Login />
          </Route>
          <Route path="*">not found</Route>
        </Switch>
      </Suspense>
    </Router>
  );
}

export default App;
