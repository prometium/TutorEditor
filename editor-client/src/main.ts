import { createApp } from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
import store from "./store";

import "normalize.css";
import "./styles/theme.scss";

createApp(App)
  .use(store)
  .use(router)
  .mount("#app");
