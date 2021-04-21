<template>
  <v-container>
    <v-data-iterator
        :items="this.torrents"
        :items-per-page="torrentPerPage"
        :page.sync="page"
        :search="search"
        hide-default-footer
    >
      <!--      FIXME: adding search function in version 1.2-->
      <div>
        <v-toolbar
            dark
            color="blue darken-3"
            class="mb-1"
        >
          <router-link to="/pubTorrent">
            <v-icon
                large
                dark
                class="ma-2"
            >
              mdi-plus-circle
            </v-icon>
          </router-link>

          <v-text-field
              v-model="search"
              clearable
              flat
              solo-inverted
              hide-details
              perpend-inner-icon="mdi-magnify"
              label="Search"
          ></v-text-field>


        </v-toolbar>
      </div>

      <v-row>
        <v-col
            v-for="torrent in torrents"
            :key="torrent.name"
            cols="12"
            sm="6"
            md="4"
            lg="3"
        >
          <TorrentUnit
              v-bind:name="torrent.name"
              v-bind:picture="torrent.picture"
              v-bind:torrent-link="torrent.torrentLink"
              v-bind:description="torrent.description"
          />
        </v-col>
      </v-row>

      <div>
        <v-row
            class="mt-2"
            align="center"
            justify="center"
        >
          <v-btn
              fab
              dark
              color="blue darken-3"
              class="mr-1"
              @click="formerPage"
          >
            <v-icon>mdi-chevron-left</v-icon>
          </v-btn>
          <v-btn
              fab
              dark
              color="blue darken-3"
              class="ml-1"
              @click="nextPage"
          >
            <v-icon>mdi-chevron-right</v-icon>
          </v-btn>
        </v-row>
        <v-row
            class="mt-2"
            align="center"
            justify="center">
          <span
              class="mr-4
        grey--text"
          >
        Page {{ page }} of {{ numberOfPages }}
        </span>
        </v-row>
      </div>
    </v-data-iterator>
  </v-container>
</template>

<script>
import TorrentUnit from "@/components/TorrentUnit";
import axios from "axios";
import JwtToken from "@/shared-state/token";

const getTorrentBaseUrl = "http://localhost:5678/test/torrents"
const staticResourceBaseUrl = "http://localhost:5678/"

export default {
  name: "TorrentTable",

  components: {
    TorrentUnit,
  },

  data: () => ({
    get_response: {},
    page: 1,
    search: '',
    torrentPerPage: 8,
    count: 8,
    keys: [],

    torrents: [],
  }),

  computed: {
    numberOfPages() {
      return Math.ceil(this.count / this.torrentPerPage)
    },
    filteredKeys() {
      return this.keys.filter(key => key !== 'Name')
    },
  },

  methods: {
    nextPage() {
      if (this.page + 1 <= this.numberOfPages) this.page += 1
      // FIXME: change data
      this.updateTorrents("", "", 2)
    },
    formerPage() {
      if (this.page - 1 >= 1) this.page -= 1
      // FIXME: change data
      this.updateTorrents("", "", "")
    },
    updateTorrents(insertStartTime, name, number) {
      const url = getTorrentBaseUrl +
          "?insertStartTime=" + insertStartTime +
          "&name=" + name +
          "&number=" + number

      let header = {
        headers: {
          "Access-Control-Allow-Origin": "*",
          "Authorization": "Bearer " + JwtToken.getToken(),
        }
      }
      console.log(header)
      axios
          .get(url, header)
          .then(response => {
            // FIXME: for test
            console.log(response)
            this.torrents = []
            this.keys = []
            this.count = response.data["data"]["number"]
            const torrent_infos = response.data["data"]["torrent_infos"]
            for (let i = 0; i < torrent_infos.length; i++) {
              this.torrents.push({
                name: torrent_infos[i]["name"],
                description: torrent_infos[i]["description"],
                picture: staticResourceBaseUrl + torrent_infos[i]["picture_path"][0],
                torrentLink: torrent_infos[i]["magnet"],
              })
              this.keys.push(torrent_infos[i]["name"])
            }
          })
          .catch(error => (console.log(error)))
    },
  },
  created() {
    this.updateTorrents("", "", "")
  },
}

</script>

<style scoped>

</style>