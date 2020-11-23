import { Store, createStore, compose, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import rootReducer from 'src/store';
import { ApplicationState } from './store';
import { api } from './store/api/middleware';

declare global {
  interface Window {
    __REDUX_DEVTOOLS_EXTENSION_COMPOSE__?: typeof compose;
  }
}

const composeEnhancers =
  (typeof window !== 'undefined' &&
    window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__) ||
  compose;

export default function configureStore(
  preloadedState: ApplicationState
): Store<ApplicationState> {
  return createStore(
    rootReducer,
    preloadedState,
    composeEnhancers(applyMiddleware(thunk, api))
  );
}
