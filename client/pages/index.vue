<template>
  <div class="wrapper">
    <a-card title="Create a new account." style="max-width: 400px">
      <a-input
        v-model="localUsername"
        placeholder="Enter username"
        type="text"
        @input="handleRegisterUsername"
      />
      <a-input
        v-model="localEmail"
        placeholder="Enter email"
        type="email"
        class="mt"
        @input="handleRegisterEmail"
      />
      <a-input
        v-model="localPassword"
        placeholder="Enter password"
        type="password"
        class="mt"
        @input="handleRegisterPassword"
      />
      <a-button type="primary" class="mt" @click="SignUp"
        >Create account.</a-button
      >
      <a-button class="mt" type="link" @click="ToLogin">Login.</a-button>
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
    localUsername: '',
  }),
  methods: {
    ToLogin() {
      this.$router.push('/login')
    },
    ...mapMutations({
      setRegisterEmail: 'SET_REGISTER_EMAIL',
      setRegisterPassword: 'SET_REGISTER_PASSWORD',
      setRegisterUsername: 'SET_REGISTER_USERNAME',
    }),
    ...mapActions({
      SignUp: 'SignUp',
    }),
    handleRegisterEmail() {
      this.setRegisterEmail(this.localEmail)
    },
    handleRegisterPassword() {
      this.setRegisterPassword(this.localPassword)
    },
    handleRegisterUsername() {
      this.setRegisterUsername(this.localUsername)
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
