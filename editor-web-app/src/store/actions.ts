import { ActionContext } from "vuex";
import { State } from "./state";
import { ActionTypes } from "./action-types";
import { Mutations } from "./mutations";
import { MutationTypes } from "./mutation-types";
import { Script, TraversableScript, Frame } from "@/common/types";
import {
  getScriptsInfo,
  getScript,
  updateScript
} from "@/common/requests";

type AugmentedActionContext = {
  commit<K extends keyof Mutations>(
    key: K,
    payload?: Parameters<Mutations[K]>[1]
  ): ReturnType<Mutations[K]>;
} & Omit<ActionContext<State, State>, "commit">;

type Actions = {
  [ActionTypes.LOAD_SCRIPTS_INFO]({
    commit
  }: AugmentedActionContext): Promise<void>;
  [ActionTypes.LOAD_SCRIPT](
    { commit }: AugmentedActionContext,
    uid: string
  ): Promise<void>;
  [ActionTypes.UPDATE_SCRIPT](
    { state }: AugmentedActionContext,
    data: {
      script?: Script;
      frames?: Frame[];
      frameIdsToDel?: string[];
      actionIdsToDel?: string[];
    }
  ): Promise<void>;
};

export const actions: Actions = {
  [ActionTypes.LOAD_SCRIPTS_INFO]({ commit }) {
    return new Promise((resolve, reject) => {
      getScriptsInfo()
        .then(data => {
          commit(MutationTypes.SET_SCRIPTS_INFO, data.scripts);
          resolve();
        })
        .catch(reject);
    });
  },
  [ActionTypes.LOAD_SCRIPT]({ commit }, uid) {
    return new Promise((resolve, reject) => {
      getScript(uid)
        .then(data => {
          const script: Script = data.script;

          const frameByUid: Record<string, Frame> = {};
          script.frames.forEach(frame => {
            frameByUid[frame.uid] = frame;
          });

          const traversableScript: TraversableScript = {
            ...script,
            frameByUid,
            path: [],
            branchNumByUid: {}
          };
          commit(MutationTypes.SET_SCRIPT, traversableScript);
          commit(MutationTypes.SELECT_FRAME, traversableScript.firstFrame.uid);
          resolve();
        })
        .catch(err => {
          reject(err);
        });
    });
  },
  [ActionTypes.UPDATE_SCRIPT](
    { state, commit },
    { script, frames, frameIdsToDel, actionIdsToDel }
  ) {
    return new Promise((resolve, reject) => {
      if (!state.script.uid) return;

      return updateScript(
        { uid: state.script.uid, ...script, frames } as Script,
        {
          frameIdsToDel,
          actionIdsToDel
        }
      )
        .then(res => {
          commit(MutationTypes.UPDATE_SCRIPT, {
            script,
            frames,
            uids: res.uids
          });
          resolve();
        })
        .catch(err => {
          reject(err);
        });
    });
  }
};
