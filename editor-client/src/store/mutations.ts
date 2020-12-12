import { MutationTree } from "vuex";
import { MutationTypes } from "./mutation-types";
import { State } from "./state";
import { ScriptInfo, TraversableScript, PathItem } from "@/common/types";

export type Mutations<S = State> = {
  [MutationTypes.SET_SCRIPTS_INFO](
    state: S,
    scriptsInfo: Array<ScriptInfo>
  ): void;
  [MutationTypes.SET_SCRIPT](state: S, script: TraversableScript): void;
  [MutationTypes.SET_FRAME](state: S, uid: string): void;
  [MutationTypes.CONFIGURE_PATH](
    state: S,
    fork?: {
      frameUid: string;
      branchNum: number;
    }
  ): void;
};

export const mutations: MutationTree<State> & Mutations = {
  [MutationTypes.SET_SCRIPTS_INFO](state, scriptsInfo) {
    state.scriptsInfo = scriptsInfo;
  },
  [MutationTypes.SET_SCRIPT](state, script) {
    state.script = script;
  },
  [MutationTypes.SET_FRAME](state, uid) {
    state.frame = state.script.frameByUid[uid];
  },
  [MutationTypes.CONFIGURE_PATH](state, fork) {
    if (fork != null) {
      state.script.branchNumByUid[fork.frameUid] = fork.branchNum;
    }

    let frameUid = state.script.firstFrame.uid;
    const path: Array<PathItem> = [];
    while (path.length <= Object.keys(state.script.frameByUid).length) {
      const pathItem: PathItem = {
        frameUid,
        branchNumber: state.script.branchNumByUid[frameUid] || 0
      };
      path.push(pathItem);

      const actions = state.script.frameByUid[frameUid].actions;
      if (actions == null || actions.length == 0) {
        break;
      }
      frameUid = actions[pathItem.branchNumber].nextFrame.uid;
    }

    state.script.path = path;
  }
};
