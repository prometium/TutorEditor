export const API_ROOT = "http://localhost:9000";

type RequestPayload = {
  endpoint: string;
  method?: string;
  data?: string | null;
};

function executeRequest({
  endpoint = "",
  method = "GET",
  data = null
}: RequestPayload) {
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

export function getScriptsInfo() {
  return executeRequest({
    endpoint: "/scripts"
  });
}

export function getScript(uid: string) {
  return executeRequest({
    endpoint: `/scripts/${uid}`
  });
}
