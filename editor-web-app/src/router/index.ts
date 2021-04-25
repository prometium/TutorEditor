import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import Editor from "../views/Editor.vue";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "Editor",
    component: Editor
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
