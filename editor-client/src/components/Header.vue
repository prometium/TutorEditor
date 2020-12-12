<template>
  <header class="header">
    <img alt="Vue logo" src="../assets/logo.png" class="logo" />
    <span class="title">{{ script.name }}</span>
    <div class="menu">
      <Dropdown text="Файл">
        <span class="menu-item">Файл</span>
        <span class="menu-item">Главная</span>
      </Dropdown>
      <span class="menu-item">Файл</span>
      <span class="menu-item">Главная</span>
    </div>
    <div class="user">User</div>
  </header>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mapState } from "vuex";
import Dropdown from "./Dropdown.vue";

export default defineComponent({
  name: "Header",
  components: {
    Dropdown
  },
  computed: {
    ...mapState(["script"])
  }
});
</script>

<style scoped lang="scss">
.header {
  display: grid;
  grid-template-areas:
    "logo title user"
    "logo menu user";
  grid-template-columns: 48px 1fr 20%;
  padding: 4px 8px 2px;
  background-color: var(--palette-background-paper);
}

.logo {
  grid-area: logo;
  align-self: center;
  width: 100%;
}

.title {
  grid-area: title;
  align-self: center;
  justify-self: start;
  padding-left: 8px;
}

.menu {
  display: flex;
  grid-area: menu;
  align-self: end;
  justify-self: start;
}

.menu-item {
  padding: 4px 8px;
  border-radius: var(--shape-border-radius);
  font-size: 0.875rem;

  &:hover,
  &:focus {
    background-color: var(--palette-action-hover);
  }

  &:active {
    background-color: var(--palette-action-selected);
  }
}

.user {
  grid-area: user;
  align-self: center;
  justify-self: end;
}
</style>
