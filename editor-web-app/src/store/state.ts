import { ScriptInfo, TraversableScript, Frame } from "@/common/types";

export interface State {
  scriptsInfo: Array<ScriptInfo>;
  script: TraversableScript;
  frameUid: string;
}

export const state: State = {
  scriptsInfo: [],
  script: {
    uid: "",
    name: "",
    version: "",
    firstFrame: {
      uid: ""
    },
    frameByUid: {},
    path: [],
    branchNumByUid: {}
  },
  frameUid: "",
};
