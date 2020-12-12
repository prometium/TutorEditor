<template>
  <div class="frame">
    <!-- <button @click="loadScriptsInfo">Загрузить скрипты</button>
    {{ scriptsInfo }}
    <input v-model="scriptUid" />
    <button @click="loadScript(scriptUid)">Загрузить скрипт</button>
    {{ scriptUid }}
    {{ frame }} -->
    <img
      :src="frame.pictureLink + '#' + new Date().getTime()"
      :alt="frame.uid"
      class="frame__img"
    />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState, mapMutations, mapActions, mapGetters } from "vuex";
import { MutationTypes } from "@/store/mutation-types";
import { ActionTypes } from "@/store/action-types";

export default Vue.extend({
  name: "Frame",
  data() {
    return {
      scriptUid: ""
    };
  },
  computed: {
    ...mapState(["scriptsInfo", "script", "frame"]),
    ...mapGetters(["path"])
  },
  methods: {
    ...mapMutations({
      setFrame: MutationTypes.SET_FRAME
    }),
    ...mapActions({
      loadScriptsInfo: ActionTypes.LOAD_SCRIPTS_INFO,
      loadScript: ActionTypes.LOAD_SCRIPT
    })
  }
});
</script>

<style scoped lang="scss">
.frame {
  display: flex;
  align-items: center;
  justify-content: center;
  overflow-y: auto;
}

.frame__img {
  max-height: 100%;
  max-width: 100%;
}
</style>
