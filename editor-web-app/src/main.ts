import { createApp } from "vue";
import { createPinia } from "pinia";
import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";
import { aliases, mdi } from "vuetify/iconsets/mdi";

import App from "./App.vue";
import router from "./router";

import "vuetify/styles";
import "@mdi/font/css/materialdesignicons.css";
import "./assets/main.css";

const pinia = createPinia();
const vuetify = createVuetify({
  components,
  directives,
  icons: {
    defaultSet: "mdi",
    aliases,
    sets: {
      mdi,
    },
  },
});
const app = createApp(App);

app.use(router);
app.use(vuetify);
app.use(pinia);

app.mount("#app");
