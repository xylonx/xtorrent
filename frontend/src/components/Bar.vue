<template>
  <v-app-bar
      app
      color="primary"
      dark
  >
    <div class="d-flex align-center">
      <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="https://cdn.vuetifyjs.com/images/logos/vuetify-logo-dark.png"
          transition="scale-transition"
          width="40"
      />

      <v-img
          alt="Vuetify Name"
          class="shrink mt-1 hidden-sm-and-down"
          contain
          min-width="100"
          src="https://cdn.vuetifyjs.com/images/logos/vuetify-name-dark.png"
          width="100"
      />
    </div>

    <v-spacer></v-spacer>

    <v-btn
        target="_blank"
        v-if="hasToken"
        @click="pubTorrent"
        text>
      <span class="mr-2">
        Pub Torrent
      </span>
    </v-btn>

    <v-btn
        href="https://github.com/xylonx/xtorrent"
        target="_blank"
        text
    >
      <span class="mr-2">Github</span>
      <v-icon>mdi-github</v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script>
import JwtToken from "@/shared-state/token";

export default {
  name: "Bar",
  data: () => ({
    hasToken: false,
  }),
  methods: {
    pubTorrent() {
      this.$router.push("/pubTorrent")
    }
  },
  created() {
    this.hasToken = JwtToken.hasToken()
  },
}
</script>

<style scoped>
a {
  text-decoration: none;
}

.router-link-active {
  text-decoration: none;
}
</style>