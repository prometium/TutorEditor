import { MutationTree } from "vuex";
import { MutationTypes } from "./mutation-types";
import { State } from "./state";
import { ScriptInfo, TraversableScript, PathItem, Frame } from "@/common/types";

export type Mutations<S = State> = {
  [MutationTypes.SET_SCRIPTS_INFO](
    state: S,
    scriptsInfo: ScriptInfo[]
  ): void;
  [MutationTypes.SET_SCRIPT](state: S, script: TraversableScript): void;
  [MutationTypes.UPDATE_FRAMES](state: S, frame: Frame[]): void;
  [MutationTypes.SELECT_FRAME](state: S, uid: string): void;
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
  [MutationTypes.UPDATE_FRAMES](state, frames) {
    frames.forEach(frame => {
      const currentFrame = state.script?.frameByUid[frame.uid];
      state.script.frameByUid[frame.uid] = {
        ...currentFrame, ...frame, actions: currentFrame.actions?.map(currentAction => {
          const newAppropriateAction = frame.actions?.find(action => action.uid === currentAction.uid);
          return { ...currentAction, ...newAppropriateAction };
        })
      };
    })
  },
  [MutationTypes.SELECT_FRAME](state, uid) {
    state.frameUid = uid;
  },
  [MutationTypes.CONFIGURE_PATH](state, fork) {
    if (fork != null) {
      state.script.branchNumByUid[fork.frameUid] = fork.branchNum;
    }

    let frameUid = state.script.firstFrame.uid;
    const path: PathItem[] = [];
    while (path.length <= Object.keys(state.script.frameByUid).length) {
      const pathItem: PathItem = {
        frameUid,
        branchNum: state.script.branchNumByUid[frameUid] || 0
      };
      path.push(pathItem);

      const actions = state.script.frameByUid[frameUid].actions;
      if (actions == null || !actions.length) break;

      const nextFrame = actions[pathItem.branchNum].nextFrame;
      if (!nextFrame) break;

      frameUid = nextFrame.uid;
    }

    state.script.path = path;
  }
};
