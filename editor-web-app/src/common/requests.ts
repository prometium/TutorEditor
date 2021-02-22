import { Script, ScriptInfo } from "./types";

export const API_ROOT = "http://localhost:9000";

type RequestPayload = {
  endpoint: string;
  method?: string;
  data?: string | null;
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

type ScriptsInfoResponse = {
  scripts: ScriptInfo[];
};

export function getScriptsInfo(): Promise<ScriptsInfoResponse> {
  return executeRequest({
    endpoint: "/scripts"
  });
}

type ScriptResponse = {
  script: Script;
};

export function getScript(uid: string): Promise<ScriptResponse> {
  return executeRequest({
    endpoint: `/scripts/${uid}`
  });
}
