<template>
  <div>
    <c-flex
      w="100vw"
      h="100vh"
      bg="csystem.snd"
      justify="center"
      align="center"
    >
      <c-box
        text-align="center"
        maxw="lg"
        bg="csystem.bg"
        px="10"
        py="8"
        rounded="lg"
        box-shadow="lg"
      >
        <c-tabs variant="soft-rounded" variant-color="gray" align="center">
          <c-tab-list>
            <c-tab>Table</c-tab>
            <c-tab>Form</c-tab>
          </c-tab-list>

          <c-tab-panels>
            <c-tab-panel>
              <Table />
            </c-tab-panel>
            <c-tab-panel>
              <Form />
            </c-tab-panel>
          </c-tab-panels>
        </c-tabs>
      </c-box>
    </c-flex>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import {
  CTabs,
  CTabList,
  CTabPanels,
  CTab,
  CTabPanel,
  CFlex,
  CBox,
} from '@chakra-ui/vue'
import Table from '../components/panelTable.vue'
import Form from '../components/panelForm.vue'

export default Vue.extend({
  name: 'Panel',
  components: {
    CTabs,
    CTabList,
    CTabPanels,
    CTab,
    CTabPanel,
    CFlex,
    CBox,
    Table,
    Form,
  },
})
</script>
