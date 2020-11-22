import React, { Suspense } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Header from 'src/components/app/Header';
import ContentWrapper from 'src/components/app/ContentWrapper';
import Editor from 'src/components/routes/Editor';

function App(): JSX.Element {
  return (
    <Router>
      <Header />
      <ContentWrapper>
        <Suspense fallback="loading....">
          <Switch>
            <Route exact path="/">
              ROOT
            </Route>
            <Route exact path="/editor">
              <Editor />
            </Route>
            <Route path="*">Not found</Route>
          </Switch>
        </Suspense>
      </ContentWrapper>
    </Router>
  );
}

export default App;
