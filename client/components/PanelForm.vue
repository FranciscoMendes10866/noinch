<template>
  <div>
    <a-card style="max-width: 400px">
      <a-input
        v-model="localWebsite"
        placeholder="Enter website"
        type="text"
        class="mt"
      />
      <a-input
        v-model="localWebsitePassword"
        placeholder="Enter password"
        type="text"
        class="mt"
      />
      <a-button type="primary" class="mt" @click="Create">Add.</a-button>
    </a-card>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  data: () => ({
    localWebsite: '',
    localWebsitePassword: '',
  }),
  methods: {
    async Create() {
      const componentState = {
        website: this.localWebsite,
        wpass: this.localWebsitePassword,
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
        this.localWebsite = ''
        this.localWebsitePassword = ''
      }
    },
  },
})
</script>

<style scoped>
.mt {
  margin-top: 20px;
}
</style>
