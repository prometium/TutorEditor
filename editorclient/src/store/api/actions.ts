import { Dispatch } from 'redux';
import { APITypes } from './types';

export const getScripts = () => (dispatch: Dispatch): Promise<void> =>
  new Promise((resolve, reject) =>
    dispatch({
      type: APITypes.CALL_API,
      payload: {
        endpoint: '/scripts',
        resolve,
        reject
      }
    })
  );
