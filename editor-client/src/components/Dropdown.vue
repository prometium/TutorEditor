<template>
  <div class="dropdown">
    <button @click.stop="toggle" tabindex="0" class="dropdown__button">
      {{ text }}
    </button>
    <div v-show="open" ref="dropdown-menu" class="dropdown__menu">
      <slot></slot>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Dropdown",
  props: {
    text: String
  },
  data() {
    return {
      open: false
    };
  },
  methods: {
    toggle() {
      this.open = !this.open;
    },
    documentClick(event: Event) {
      const el = this.$refs["dropdown-menu"] as HTMLDivElement;
      const target = event.target as HTMLElement;
      if (el !== target) {
        this.open = false;
      }
    }
  },
  mounted() {
    document.addEventListener("click", this.documentClick);
  },
  unmounted() {
    document.removeEventListener("click", this.documentClick);
  }
});
</script>

<style scoped lang="scss">
.dropdown {
  position: relative;
  display: inline-block;

  &__button {
    padding: 4px 8px;
    border: 0;
    border-radius: var(--shape-border-radius);
    outline: 0;
    font-size: 0.875rem;
    background-color: transparent;
    cursor: pointer;
  }

  &__menu {
    position: absolute;
  }
}
</style>
