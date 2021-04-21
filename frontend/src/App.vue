<template>
  <v-app>
    <Bar/>

    <vue-particles
        :clickEffect="true"
        clickMode="repulse"
    ></vue-particles>

    <v-main>
      <router-view>
        <Login/>
      </router-view>
    </v-main>

    <Footer/>
  </v-app>
</template>

<script>
import Footer from './components/Footer'
import Bar from './components/Bar'
import Login from "@/components/Login";

export default {
  name: 'App',
  components: {
    Login,
    Bar,
    Footer,
  },
  data: () => ({}),
};
</script>

<style>
#particles-js {
  width: 100%;
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  z-index: 0;
}
</style>