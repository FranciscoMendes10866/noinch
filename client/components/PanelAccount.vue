<template>
  <div>
    <div>
      <span>
        <a-badge dot color="#87d068" :size="64"
          ><a-avatar :size="64" shape="square">{{
            userName
          }}</a-avatar></a-badge
        >
      </span>
    </div>
    <div>
      <a-button class="mt" type="primary" @click="LogOut">Logout.</a-button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapActions, mapGetters } from 'vuex'

export default Vue.extend({
  computed: {
    ...mapGetters({
      userName: 'userName',
    }),
  },
  methods: {
    ...mapActions({
      LogOut: 'LogOut',
    }),
  },
})
</script>

<style scoped>
.mt {
  margin-top: 10px;
}
</style>
