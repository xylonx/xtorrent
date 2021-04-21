<template>
  <v-footer
      dark
      padless
  >
    <v-card
        flat
        tile
        class="indigo lighten-1 white--text text-center"
    >
      <v-card-text>
        <v-btn
            v-for="icon in icons"
            :key="icon"
            class="mx-4 white--text"
            icon
        >
          <v-icon size="24px">
            {{ icon }}
          </v-icon>
        </v-btn>
      </v-card-text>

      <v-card-text class="white--text pt-0">
        The free software movement is basically a movement for freedom. It's based on values that are not purely
        material and practical. It's based on the idea that freedom is a benefit in itself. And that being allowed to be
        part of a community is a benefit in itself, having neighbors who can help you, who are free to help you—they are
        not told that they are pirates if they help you—is a benefit in itself, and that that's even more important than
        how powerful and reliable your software is.
      </v-card-text>

      <v-divider></v-divider>

      <v-card-text class="white--text">
        Powered By <strong>Vuetify</strong> |
        Made By <strong>Xylonx</strong>
      </v-card-text>
    </v-card>
  </v-footer>
</template>

<script>
export default {
  name: "Footer",
  data: () => ({
    icons: [
      "mdi-facebook",
      "mdi-twitter",
      "mdi-linkedin",
      "mdi-instagram",
    ]
  })
}
</script>

<style scoped>

</style>