<template>
  <div class="wrapper">
    <a-card title="Welcome back." style="max-width: 400px">
      <a-input placeholder="Enter email" type="email" class="mt" />
      <a-input placeholder="Enter password" type="password" class="mt" />
      <a-button type="primary" class="mt">Login.</a-button>
      <a-button class="mt" type="link" @click="ToRegister"
        >Create account.</a-button
      >
    </a-card>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  methods: {
    ToRegister() {
      this.$router.push('/')
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
