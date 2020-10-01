<template>
  <div class="wrapper">
    <a-card class="w">
      <a-tabs default-active-key="1">
        <a-tab-pane key="1" tab="Table">
          <c-table />
        </a-tab-pane>
        <a-tab-pane key="2" tab="Form" force-render>
          Content of Tab Pane 2
        </a-tab-pane>
      </a-tabs>
    </a-card>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  components: {
    CTable: () => import('@/components/PanelTable.vue'),
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
.w {
  min-width: 700px;
}
@media only screen and (max-width: 600px) {
  .w {
    min-width: 400px;
  }
}
</style>
