import { Script, ScriptInfo } from "./types";

export const API_ROOT = "http://localhost:9000";

type RequestPayload = {
  endpoint: string;
  method?: "GET" | "PUT" | "POST";
  data?: string | FormData | null;
};

function executeRequest<T>({
  endpoint = "",
  method = "GET",
  data = null
}: RequestPayload): Promise<T> {
  return new Promise((resolve, reject) => {
    fetch(API_ROOT + endpoint, {
      method,
      body: data
    })
      .then(response => {
        if (!response.ok) {
          throw new Error(response.statusText);
        }
        return response.json();
      })
      .then(data => {
        resolve(data);
      })
      .catch(err => {
        reject(err);
      });
  });
}

type GetScriptsInfoResponse = {
  scripts: ScriptInfo[];
};

export function getScriptsInfo(): Promise<GetScriptsInfoResponse> {
  return executeRequest({
    endpoint: "/scripts"
  });
}

type GetScriptResponse = {
  script: Script;
};

export function getScript(uid: string): Promise<GetScriptResponse> {
  return executeRequest({
    endpoint: `/scripts/${uid}`
  });
}

type CreateScriptResponse = {
  uid: string
}

export function createScript(script: FormData): Promise<CreateScriptResponse> {
  return executeRequest({
    endpoint: '/raw',
    method: 'POST',
    data: script
  });
}