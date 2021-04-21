<template>
  <v-layout
      wrap
      align-center
      justify-center
      fill-height>
    <v-container
        v-show="!alert"
    >
      <validation-observer
          ref="observer"
          v-slot="{ invalid }"
      >
        <form
            @submit.prevent="submit"
        >
          <v-row>
            <v-spacer></v-spacer>
            <v-col cols="7">
              <validation-provider
                  v-slot="{errors}"
                  name="magnet"
                  rules="required"
              >
                <v-text-field
                    v-model="magnet"
                    :error-messages="errors"
                    label="magnet"
                    required
                    filled
                    rounded
                ></v-text-field>
              </validation-provider>
            </v-col>
            <v-spacer></v-spacer>
          </v-row>

          <v-row>
            <v-spacer></v-spacer>
            <v-col cols="7">
              <validation-provider
                  v-slot="{errors}"
                  name="name"
                  rules="required"
              >
                <v-text-field
                    v-model="name"
                    :error-messages="errors"
                    label="name"
                    required
                    filled
                    rounded
                ></v-text-field>
              </validation-provider>
            </v-col>
            <v-spacer></v-spacer>
          </v-row>

          <v-row>
            <v-spacer></v-spacer>
            <v-col cols="7">
              <validation-provider
                  v-slot="{errors}"
                  name="description"
                  rules="required"
              >
                <v-text-field
                    v-model="description"
                    :error-messages="errors"
                    label="description"
                    required
                    filled
                    rounded
                ></v-text-field>
              </validation-provider>
            </v-col>
            <v-spacer></v-spacer>
          </v-row>

          <v-row>
            <v-spacer></v-spacer>
            <v-col cols="7">
              <validation-provider
                  v-slot="{errors}"
                  name="picture"
                  rules="required"
              >
                <v-file-input
                    v-model="picture"
                    accept="image/*"
                    label="picture"
                    :error-messages="errors"
                    requied
                    filled
                    rounded
                ></v-file-input>
              </validation-provider>
            </v-col>
            <v-spacer></v-spacer>
          </v-row>

          <!--          Button-->
          <v-row>
            <v-spacer></v-spacer>
            <v-btn
                class="mr-10 pa-2"
                type="submit"
                :disabled="invalid"
            >submit
            </v-btn>
            <v-btn
                class="mr-4 pa-2"
                @click="clear"
            >clear
            </v-btn>

            <v-spacer></v-spacer>
          </v-row>
        </form>
      </validation-observer>
    </v-container>

    <v-container>
      <v-alert
          v-model="alert"
          border="left"
          dismissible
          elevation="15"
          type="error"
          prominent
      >
        {{ alertMeg }}
      </v-alert>
    </v-container>

  </v-layout>
</template>

<script>
import {required} from 'vee-validate/dist/rules'
import {extend, ValidationObserver, ValidationProvider, setInteractionMode} from 'vee-validate'
import axios from "axios";
import JwtToken from "@/shared-state/token";

setInteractionMode("eager")

extend('required', {
  ...required,
  message: '{_field_} can not be empty'
})

export default {
  name: "PubTorrent",
  components: {
    ValidationProvider,
    ValidationObserver,
  },

  data: () => ({
    valid: true,
    magnet: "",
    name: "",
    description: "",
    picture: "",
    alert: false,
    alertMeg: "",
  }),

  methods: {
    submit() {
      let formData = new FormData()
      formData.append("magnet", this.magnet)
      formData.append("name", this.name)
      formData.append("description", this.description)
      formData.append("picture", this.picture)

      let header = {
        headers: {
          "Access-Control-Allow-Origin": "*",
          "Authorization": "Bearer " + JwtToken.getToken(),
        }
      }

      axios
          .post("http://localhost:5678/torrent", formData, header)
          .then(() => {
            this.$router.back()
          })
          .catch(error => {
            this.alert = true
            this.alertMeg = error.response.data["message"]
          })
    },

    clear() {
      this.valid  = false
      this.magnet = ""
      this.name = ""
      this.description = ""
      this.picture = ""
    },
  },

}
</script>

<style scoped>

</style>