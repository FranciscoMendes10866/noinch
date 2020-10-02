<template>
  <div class="wrapper">
    <a-card class="w">
      <a-tabs default-active-key="1">
        <a-tab-pane key="1" tab="Table">
          <c-table />
        </a-tab-pane>
        <a-tab-pane key="2" tab="Form" force-render>
          <c-form />
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
    CForm: () => import('@/components/PanelForm.vue'),
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
  max-width: 700px;
}
</style>
