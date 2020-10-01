<template>
  <div>
    <c-input
      v-model="localWebsite"
      text-align="center"
      focus-border-color="csystem.strk"
      mt="6"
      type="text"
      placeholder="Website domain."
      size="lg"
    />
    <c-input
      v-model="localPassword"
      text-align="center"
      focus-border-color="csystem.strk"
      mt="4"
      type="text"
      placeholder="Website password."
      size="lg"
    />
    <c-box text-align="center" mt="6">
      <c-button bg="csystem.btn" size="lg" box-shadow="sm" @click="Create">
        <c-text color="csystem.btx">Add</c-text>
      </c-button>
    </c-box>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { CBox, CInput, CButton, CText } from '@chakra-ui/vue'

export default Vue.extend({
  name: 'Form',
  components: {
    CInput,
    CButton,
    CText,
    CBox,
  },
  data: () => ({
    localWebsite: '',
    localPassword: '',
  }),
  methods: {
    async Create() {
      const componentState = {
        website: this.localWebsite,
        wpass: this.localPassword,
      }
      const send = await this.$axios.$post(
        'http://localhost:4444/api/v1/manager',
        componentState,
        {
          headers: {
            Authorization: `Bearer ${this.$store.state.Token}`,
          },
        }
      )
      if (!send) {
        // eslint-disable-next-line no-console
        console.log('Error.')
      } else {
        this.website = ''
        this.wpass = ''
      }
    },
  },
})
</script>
