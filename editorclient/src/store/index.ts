import { combineReducers } from 'redux';
import { scriptsReducer } from './scripts/reducer';
import { ScriptsState } from './scripts/types';

export interface ApplicationState {
  scripts: ScriptsState;
}

const rootReducer = combineReducers({
  scripts: scriptsReducer
});

export default rootReducer;
