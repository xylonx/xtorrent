<template>
  <v-container
      wrap
      align-center
      justify-center
      fill-height
  >
    <v-layout
        wrap
        align-center
        justify-center
        fill-height
        v-show="!alert"
    >
      <v-container>
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
                    name="email"
                    rules="required|email"
                >
                  <v-text-field
                      v-model="email"
                      :error-messages="errors"
                      label="Email"
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
                    name="password"
                    rules="required|min:8"
                >
                  <v-text-field
                      v-model="password"
                      :error-messages="errors"
                      label="password"
                      :append-icon="showPassword?'mdi-eye':'mdi-eye-off'"
                      :type="showPassword? 'text' : 'password'"
                      hint="At lease 8 characters"
                      @click:append="showPassword = !showPassword"
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
              <v-btn
                  class="mr-10 pa-2"
                  type="submit"
                  :disabled="invalid"
              >submit
              </v-btn>
              <v-btn
                  class="mr-10 pa-2"
                  @click="clear"
              >clear
              </v-btn>
              <router-link to="/register">
                <v-btn
                    class="mr-10 pa-2">
                  Register
                </v-btn>
              </router-link>

              <v-spacer></v-spacer>
            </v-row>

          </form>
        </validation-observer>
      </v-container>
    </v-layout>


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


  </v-container>
</template>

<script>
import {required, email, min} from 'vee-validate/dist/rules'
import {extend, ValidationObserver, ValidationProvider, setInteractionMode} from 'vee-validate'
import axios from "axios";
import JwtToken from "@/shared-state/token";

const LoginURL = "http://localhost:5678/login"

setInteractionMode("eager")

extend('required', {
  ...required,
  message: '{_field_} can not be empty'
})

extend('min', {
  ...min,
  message: '{_field_} may not be less than {length} characters'
})

extend('email', {
  ...email,
  message: "Email must be valid",
})

export default {

  name: "Login",

  components: {
    ValidationProvider,
    ValidationObserver,
  },

  data: () => ({
    email: "",
    password: "",
    showPassword: false,
    alert: false,
    alertMeg: "",
  }),

  methods: {
    submit() {
      const body = {
        "email_address": this.email,
        "password": this.password,
      }
      axios
          .post(LoginURL, body)
          .then(response => {
            // Store token
            JwtToken.setToken(response.data["token"])
            this.$router.push("/torrents")
          })
          .catch(error => {
            this.alert = true
            this.alertMeg = error.response.data["message"]
            this.clear()
          })
    },

    clear() {
      this.email = ""
      this.password = ""
    },
  },
}
</script>

<style scoped>

</style>