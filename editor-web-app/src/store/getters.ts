import { GetterTree } from "vuex";
import { State } from "./state";
import { PathItem } from "@/common/types";

type Getters = {
  path(state: State): Array<PathItem>;
};

export const getters: GetterTree<State, State> & Getters = {
  path(state) {
    if (state.script == null) return [];
    return state.script.path;
  }
};
