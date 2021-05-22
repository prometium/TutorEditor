import { MutationTypes } from "./mutation-types";
import { State } from "./state";
import { ScriptInfo, TraversableScript, Frame, Script } from "@/common/types";

export type Mutations<S = State> = {
  [MutationTypes.SET_SCRIPTS_INFO](state: S, scriptsInfo: ScriptInfo[]): void;
  [MutationTypes.SET_SCRIPT](state: S, script: TraversableScript): void;
  [MutationTypes.UPDATE_SCRIPT](
    state: S,
    data: {
      script?: Script;
      frames?: Frame[];
      uids?: Record<string, string> | null;
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
    { script = {}, frames = [], uids = {} }
  ) {
    state.script = {
      ...state.script,
      ...script
    };

    const framesWithCorrectUids = !uids
      ? frames
      : frames.map(frame => {
          const actionsWithCorrectUids = frame.actions?.map(action => {
            const newNextFrameUid = uids?.[action.nextFrame?.uid.slice(2)];
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
      } else {
        state.script.frameByUid[frame.uid] = frame;
      }
    });
  },
  [MutationTypes.SELECT_FRAME](state, uid) {
    state.frameUid = uid;
  },
  [MutationTypes.CONFIGURE_PATH](state, fork) {
    state.script.branchNumByUid[fork.frameUid] = fork.branchNum;
  }
};
