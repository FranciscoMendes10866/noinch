<template>
  <div>
    <c-flex w="100vw" h="100vh" bg="csystem.bg" justify="center" align="center">
      <c-box
        maxw="lg"
        bg="csystem.mn"
        px="16"
        py="12"
        rounded="lg"
        box-shadow="lg"
      >
        <c-text
          text-align="center"
          mb="8"
          color="csystem.btx"
          font-size="2xl"
          font-weight="semibold"
          >Welcome back.</c-text
        >
        <c-input
          v-model="localEmail"
          text-align="center"
          focus-border-color="csystem.strk"
          type="email"
          variant="flushed"
          placeholder="Type your email."
          size="lg"
          @input="setLoginEmail"
        />
        <c-input
          v-model="localPassword"
          text-align="center"
          focus-border-color="csystem.strk"
          mt="4"
          type="password"
          variant="flushed"
          placeholder="Type your password."
          size="lg"
          @input="setLoginPassword"
        />
        <c-box text-align="center" mt="8">
          <c-button bg="csystem.btn" size="lg" box-shadow="sm" @click="SignIn">
            <c-text color="csystem.btx">Sign In</c-text>
          </c-button>
        </c-box>
        <c-text mt="8">
          Don't have an account yet? Then
          <c-link as="nuxt-link" to="/signup" color="csystem.btn"
            >Sign Up</c-link
          >
          .
        </c-text>
      </c-box>
    </c-flex>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapMutations, mapActions } from 'vuex'
import { CFlex, CBox, CInput, CButton, CText, CLink } from '@chakra-ui/vue'

export default Vue.extend({
  name: 'SignIn',
  components: {
    CFlex,
    CBox,
    CInput,
    CButton,
    CText,
    CLink,
  },
  data: () => ({
    localEmail: '',
    localPassword: '',
  }),
  methods: {
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
