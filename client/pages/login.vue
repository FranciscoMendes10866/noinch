<template>
  <div class="wrapper">
    <a-card title="Welcome back." style="max-width: 400px">
      <a-input
        v-model="localEmail"
        placeholder="Enter email"
        type="email"
        class="mt"
        @input="handleLoginEmail"
      />
      <a-input
        v-model="localPassword"
        placeholder="Enter password"
        type="password"
        class="mt"
        @input="handleLoginPassword"
      />
      <a-button type="primary" class="mt" @click="SignIn">Login.</a-button>
      <a-button class="mt" type="link" @click="ToRegister"
        >Create account.</a-button
      >
    </a-card>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapMutations, mapActions } from 'vuex'

export default Vue.extend({
  data: () => ({
    localEmail: '',
    localPassword: '',
  }),
  methods: {
    ToRegister() {
      this.$router.push('/')
    },
    ...mapMutations({
      setLoginEmail: 'SET_LOGIN_EMAIL',
      setLoginPassword: 'SET_LOGIN_PASSWORD',
    }),
    ...mapActions({
      SignIn: 'SignIn',
    }),
    handleLoginEmail() {
      this.setLoginEmail(this.localEmail)
    },
    handleLoginPassword() {
      this.setLoginPassword(this.localPassword)
    },
  },
})
</script>

<style scoped>
.wrapper {
  height: 100vh;
  width: 100vw;
  display: flex;
  justify-content: center;
  align-items: center;
}
.mt {
  margin-top: 20px;
}
</style>
