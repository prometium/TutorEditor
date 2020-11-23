import { ActionCreator, Action } from 'redux';
import { ApiAction, APITypes } from './types';

const API_ROOT = 'http://localhost:9000';

export const api = () => (next: ActionCreator<Action>) => async (
  action: Action
): Promise<void> => {
  if (action.type !== APITypes.CALL_API) {
    next(action);
    return;
  }

  const {
    endpoint,
    method = 'GET',
    //withCredentials = true,
    data = null,
    contentType = 'application/json',
    resolve,
    reject
  } = (action as ApiAction).payload;

  const headers = new Headers();

  //   if (withCredentials) {
  //     const token = getState().auth.token;

  //     if (!token) return;

  //     headers.append('Authorization', `Bearer ${token}`);
  //   }

  if (data && contentType) {
    headers.append('Content-Type', contentType);
  }

  try {
    const response = await fetch(`${API_ROOT}${endpoint}`, {
      method,
      headers,
      body: data
    });

    if (!response.ok) {
      throw new HttpError(response.statusText, response.status);
    }

    const responseData = await response.json();
    resolve(responseData);
  } catch (error) {
    if (error instanceof HttpError) {
      //console.error(error.statusCode);
    }
    reject(error.message);
  }
};

class HttpError extends Error {
  statusCode: number;

  constructor(message: string, statusCode: number) {
    super(message);
    this.name = this.constructor.name;
    this.statusCode = statusCode;
  }
}
