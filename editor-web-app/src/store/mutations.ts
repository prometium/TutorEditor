import { MutationTypes } from "./mutation-types";
import { State } from "./state";
import { ScriptInfo, TraversableScript, Frame, Script } from "@/common/types";

export type Mutations<S = State> = {
  [MutationTypes.SET_SCRIPTS_INFO](state: S, scriptsInfo: ScriptInfo[]): void;
  [MutationTypes.SET_SCRIPT](state: S, script: TraversableScript): void;
  [MutationTypes.UPDATE_SCRIPT](
    state: S,
    data: { script?: Script; frames?: Frame[] }
  ): void;
  [MutationTypes.SELECT_FRAME](state: S, uid?: string): void;
  [MutationTypes.CONFIGURE_PATH](
    state: S,
    fork: {
      frameUid: string;
      branchNum: number;
    }
  ): void;
};

export const mutations: Mutations = {
  [MutationTypes.SET_SCRIPTS_INFO](state, scriptsInfo) {
    state.scriptsInfo = scriptsInfo;
  },
  [MutationTypes.SET_SCRIPT](state, script) {
    state.script = script;
  },
  [MutationTypes.UPDATE_SCRIPT](state, { script = {}, frames = [] }) {
    state.script = {
      ...state.script,
      ...script
    };
    frames.forEach(frame => {
      const currentFrame = state.script?.frameByUid[frame.uid];
      state.script.frameByUid[frame.uid] = {
        ...currentFrame,
        ...frame,
        actions: currentFrame.actions?.map(currentAction => {
          const newAppropriateAction = frame.actions?.find(
            action => action.uid === currentAction.uid
          );
          return { ...currentAction, ...newAppropriateAction };
        })
      };
    });
  },
  [MutationTypes.SELECT_FRAME](state, uid) {
    state.frameUid = uid;
  },
  [MutationTypes.CONFIGURE_PATH](state, fork) {
    state.script.branchNumByUid[fork.frameUid] = fork.branchNum;
  }
};
