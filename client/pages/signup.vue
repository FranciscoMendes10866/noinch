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
          mb="6"
          color="csystem.btx"
          font-size="2xl"
          font-weight="semibold"
          >Create a new account.</c-text
        >
        <c-input
          v-model="localUsername"
          text-align="center"
          focus-border-color="csystem.strk"
          type="text"
          variant="flushed"
          placeholder="Type your username."
          size="lg"
          @input="setRegisterUsername"
        />
        <c-input
          v-model="localEmail"
          text-align="center"
          mt="4"
          focus-border-color="csystem.strk"
          type="email"
          variant="flushed"
          placeholder="Type your email."
          size="lg"
          @input="setRegisterEmail"
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
          @input="setRegisterPassword"
        />
        <c-box text-align="center" mt="6">
          <c-button bg="csystem.btn" size="lg" box-shadow="sm" @click="SignUp">
            <c-text color="csystem.btx">Sign Up</c-text>
          </c-button>
        </c-box>
        <c-text mt="6">
          Already have an account? Then
          <c-link as="nuxt-link" to="/" color="csystem.btn">Sign In</c-link>
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
  name: 'SignUp',
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
    localUsername: '',
  }),
  methods: {
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
