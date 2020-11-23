export enum APITypes {
  CALL_API = '@@CALL_API'
}

export type ApiAction = {
  type: typeof APITypes.CALL_API;
  payload: {
    endpoint: string;
    method?: string;
    data?: string;
    contentType?: string;
    resolve: (x: void | PromiseLike<void> | undefined) => void;
    reject: (x: unknown) => void;
  };
};
