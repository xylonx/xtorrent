<template>
  <v-hover v-slot="{ hover }">
    <v-card
        :elevation="hover ? 12 : 2"
        :class="{ 'on-hover': hover }"
        class="mx-auto"
        max-width="360"
    >
      <v-img
          class="white--text align-end"
          height="160px"
          v-bind:src="picture"
      >
        <v-card-title>{{ name }}</v-card-title>
      </v-img>

      <v-card-text class="text--primary">
        <div>{{ description }}</div>
      </v-card-text>

      <v-card-actions>
        <v-btn
            color="orange"
            text
            v-bind:href="torrentLink"
        >
          TorrentLink
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-hover>
</template>

<script>
export default {
  name: "TorrentUnit",
  props: {
    name: String,
    description: String,
    picture: String,
    torrentLink: String,
  },
}
</script>

<style scoped>

</style>