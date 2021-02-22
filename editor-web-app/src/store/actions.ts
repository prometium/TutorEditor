import { ActionTree, ActionContext } from "vuex";
import { State } from "./state";
import { ActionTypes } from "./action-types";
import { Mutations } from "./mutations";
import { MutationTypes } from "./mutation-types";
import { Script, TraversableScript, Frame } from "@/common/types";
import { API_ROOT, getScriptsInfo, getScript } from "@/common/requests";

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
};

export const actions: ActionTree<State, State> & Actions = {
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
            frameByUid[frame.uid].pictureLink = API_ROOT + frame.pictureLink;
          });

          const traversableScript: TraversableScript = {
            ...script,
            frameByUid,
            path: [],
            branchNumByUid: {}
          };
          commit(MutationTypes.SET_SCRIPT, traversableScript);
          commit(MutationTypes.SET_FRAME, traversableScript.firstFrame.uid);
          commit(MutationTypes.CONFIGURE_PATH);
          resolve();
        })
        .catch(err => {
          reject(err);
        });
    });
  }
};
