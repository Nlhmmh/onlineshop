<template>
  <div id="nav">
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link> |
    <router-link to="/test">Test</router-link> |
    <a href="#" @click="login">Login</a>
  </div>
  <router-view />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mapMutations } from "vuex";
import { RouteLocationNormalized } from "vue-router";

export default defineComponent({
  methods: {
    ...mapMutations(["setUser"]),
    login() {
      console.log("called login method");
      const user = {
        name: "Fake user",
      };

      // mapMutations will effectively create a method in our component with the same name as our mutation
      this.setUser(user);
    },
  },
  watch: {
    $route: {
      handler: function (to: RouteLocationNormalized): void {
        document.title = to.meta.title || "Online Shopping";
      },
      immediate: true,
    },
  },
});
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
