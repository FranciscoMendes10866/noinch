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
          >Create a new account.</c-text
        >
        <c-input
          text-align="center"
          focus-border-color="csystem.strk"
          type="email"
          variant="flushed"
          placeholder="Type your email."
          size="lg"
        />
        <c-input
          text-align="center"
          focus-border-color="csystem.strk"
          mt="4"
          type="password"
          variant="flushed"
          placeholder="Type your password."
          size="lg"
        />
        <c-box text-align="center" mt="8">
          <c-button bg="csystem.btn" size="lg" box-shadow="sm">
            <c-text color="csystem.btx">Sign Up</c-text>
          </c-button>
        </c-box>
        <c-text mt="8">
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
})
</script>
