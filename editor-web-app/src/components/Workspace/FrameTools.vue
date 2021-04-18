<template>
  <div class="frame-tools">
    <v-btn elevation="1" icon>
      <v-icon>mdi-arrow-up</v-icon>
    </v-btn>
    <v-btn elevation="1" icon>
      <v-icon>mdi-arrow-down</v-icon>
    </v-btn>
    <v-btn elevation="1" icon>
      <v-icon>mdi-plus</v-icon>
    </v-btn>
    <v-btn elevation="1" icon>
      <v-icon>mdi-delete</v-icon>
    </v-btn>
  </div>
</template>

<style scoped lang="scss">
.frame-tools {
  display: flex;
  justify-content: space-around;
  padding: 4px 0;
}
</style>
