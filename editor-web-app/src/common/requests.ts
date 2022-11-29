import type { Script, ScriptInfo } from "./types";

export const API_ROOT = `${import.meta.env.VITE_SVC_URL || ""}/api`;

type RequestPayload = {
  endpoint: string;
  method?: "GET" | "PUT" | "POST" | "DELETE";
  data?: string | FormData | null;
  headers?: Headers;
};

function executeRequest<T>({
  endpoint = "",
  method = "GET",
  data = null,
  headers,
}: RequestPayload): Promise<T> {
  return new Promise((resolve, reject) => {
    fetch(API_ROOT + endpoint, {
      method,
      body: data,
      headers,
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error(response.statusText);
        }

        const contentType = response.headers.get("content-type");
        if (!contentType) return response.text();

        if (contentType.includes("application/json")) {
          return response.json();
        }
        if (contentType.includes("application/zip")) {
          return response.blob();
        }
        return response.text();
      })
      .then((data) => {
        resolve(data);
      })
      .catch((err) => {
        reject(err);
      });
  });
}

type GetScriptsInfoResponse = {
  scripts: ScriptInfo[];
};

export function getScriptsInfo(): Promise<GetScriptsInfoResponse> {
  return executeRequest({
    endpoint: "/scripts",
  });
}

type GetScriptResponse = {
  script: Script;
};

export function getScript(uid: string): Promise<GetScriptResponse> {
  return executeRequest({
    endpoint: `/scripts/${uid}`,
  });
}

type CreateScriptResponse = {
  uid: string;
};

export function createScript(
  scriptData: FormData
): Promise<CreateScriptResponse> {
  return executeRequest({
    endpoint: "/archive",
    method: "POST",
    data: scriptData,
  });
}

export function createScriptV2(
  scriptData: FormData
): Promise<CreateScriptResponse> {
  return executeRequest({
    endpoint: "/archiveV2",
    method: "POST",
    data: scriptData,
  });
}

type UpdateScriptResponse = {
  uids?: Record<string, string> | null;
};

export function updateScript(
  script: Script,
  {
    frameIdsToDel,
    actionIdsToDel,
  }: { frameIdsToDel?: string[]; actionIdsToDel?: string[] } = {}
): Promise<UpdateScriptResponse> {
  return executeRequest({
    endpoint: "/scripts",
    method: "PUT",
    data: JSON.stringify({ script, frameIdsToDel, actionIdsToDel }),
  });
}

type DeleteScriptResponse = Record<string, unknown>;

export function deleteScript(uid: string): Promise<DeleteScriptResponse> {
  return executeRequest({
    endpoint: `/scripts/${uid}`,
    method: "DELETE",
  });
}

export function downloadScriptArchive(uid: string): Promise<Blob> {
  return executeRequest({
    endpoint: `/archiveV2/${uid}`,
    method: "GET",
  });
}

type ReleaseScriptArchiveResponse = void;

export function releaseScriptArchive(
  uid: string
): Promise<ReleaseScriptArchiveResponse> {
  return executeRequest({
    endpoint: `/archiveV2`,
    method: "PUT",
    data: JSON.stringify({ uid }),
  });
}

type AddImageResponse = {
  link: string;
};

export function addImage(imageData: FormData): Promise<AddImageResponse> {
  return executeRequest({
    endpoint: "/images",
    method: "POST",
    data: imageData,
  });
}

type CopyScriptResponse = {
  uid: string;
};

export function copyScript(script: Script): Promise<CopyScriptResponse> {
  return executeRequest({
    endpoint: "/scripts",
    method: "POST",
    data: JSON.stringify({ script }),
  });
}
