import Vue from "vue";
import { ScriptInfo, TraversableScript, Frame, Script } from "@/common/types";
import { MutationTypes } from "./mutation-types";
import { State } from "./state";

export type Mutations<S = State> = {
  [MutationTypes.SET_SCRIPTS_INFO](state: S, scriptsInfo: ScriptInfo[]): void;
  [MutationTypes.SET_SCRIPT](state: S, script: TraversableScript): void;
  [MutationTypes.UPDATE_SCRIPT](
    state: S,
    data: {
      script?: Script;
      frames?: Frame[];
      uids?: Record<string, string> | null;
      frameIdsToDel?: string[];
      actionIdsToDel?: string[];
    }
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
  [MutationTypes.UPDATE_SCRIPT](
    state,
    {
      script = {},
      frames = [],
      uids = {},
      frameIdsToDel = [],
      actionIdsToDel = []
    }
  ) {
    state.script = {
      ...state.script,
      ...script
    };

    const framesWithCorrectUids = !uids
      ? frames
      : frames.map(frame => {
          const actionsWithCorrectUids = frame.actions?.map(action => {
            const newNextFrameUid = action.nextFrame
              ? uids?.[action.nextFrame?.uid.slice(2)]
              : null;
            const nextFrameWithCorrectUid = newNextFrameUid
              ? { uid: newNextFrameUid }
              : action.nextFrame;

            const newActionUid = uids?.[action.uid.slice(2)];
            return newActionUid
              ? {
                  ...action,
                  uid: newActionUid,
                  nextFrame: nextFrameWithCorrectUid
                }
              : { ...action, nextFrame: nextFrameWithCorrectUid };
          });

          const newFrameUid = uids?.[frame.uid.slice(2)];
          return newFrameUid
            ? {
                ...frame,
                uid: newFrameUid,
                actions: actionsWithCorrectUids
              }
            : { ...frame, actions: actionsWithCorrectUids };
        });

    framesWithCorrectUids.forEach(frame => {
      const currentFrame = state.script?.frameByUid[frame.uid];
      if (currentFrame) {
        const currentActions =
          currentFrame.actions?.map(currentAction => {
            const newAppropriateAction = frame.actions?.find(
              action => action.uid === currentAction.uid
            );
            return { ...currentAction, ...newAppropriateAction };
          }) || [];

        const newActions =
          frame.actions?.filter(
            action =>
              !currentActions.find(
                currentAction => currentAction.uid === action.uid
              )
          ) || [];

        state.script.frameByUid[frame.uid] = {
          ...currentFrame,
          ...frame,
          actions: [...currentActions, ...newActions]
        };
      } else {
        state.script.frameByUid[frame.uid] = frame;
      }
    });

    frameIdsToDel.forEach(frameId => {
      Vue.delete(state.script.frameByUid, frameId);
    });

    Object.values(state.script.frameByUid).forEach(frame => {
      if (frame.actions) {
        Vue.set(
          frame,
          "actions",
          frame.actions?.filter(action => !actionIdsToDel.includes(action.uid))
        );
      }
    });
  },
  [MutationTypes.SELECT_FRAME](state, uid) {
    state.frameUid = uid;
  },
  [MutationTypes.CONFIGURE_PATH](state, fork) {
    Vue.set(state.script.branchNumByUid, fork.frameUid, fork.branchNum);
  }
};
