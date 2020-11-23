import { Reducer } from 'redux';
import { ScriptsState } from './types';

export const initialState: ScriptsState = {
  all: [1]
};

const scriptsReducer: Reducer<ScriptsState> = function (
  state = initialState,
  action
) {
  return state;
};

export { scriptsReducer };
