import { GetterTree } from "vuex";
import { State } from "./state";
import { Frame, PathItem } from "@/common/types";

type Getters = {
  path(state: State): Array<PathItem>;
  frame(state: State): Frame;
};

export const getters: GetterTree<State, State> & Getters = {
  path(state) {
    if (state.script == null) return [];
    return state.script.path;
  },
  frame(state) {
    return state.script.frameByUid[state.frameUid] || {}
  }
};
