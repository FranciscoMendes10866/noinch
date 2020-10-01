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
        <!-- Header -->
        <c-divider orientation="horizontal" />
        <c-flex direction="row" justify="space-between">
          <c-box>
            <c-text color="csystem.hdl">Website</c-text>
          </c-box>
          <c-box ml="6">
            <c-text color="csystem.hdl">Password</c-text>
          </c-box>
          <c-box ml="6">
            <c-text color="csystem.hdl">🗑️</c-text>
          </c-box>
          <c-box ml="6">
            <c-text color="csystem.hdl">✏️</c-text>
          </c-box>
        </c-flex>
        <c-divider orientation="horizontal" />
        <!-- End of the Header -->
        <!--  -->
        <!-- Content -->
        <div>
          <c-flex direction="row" justify="space-between" align="center" py="2">
            <c-box>
              <a href="" target="blank"
                ><c-text color="csystem.pp">www.google.pt</c-text></a
              >
            </c-box>
            <c-box ml="6">
              <c-text color="csystem.pp">sadasdsaIY79</c-text>
            </c-box>
            <c-box ml="6">
              <c-button bg="csystem.thr" color="csystem.btx" size="sm"
                >Delete</c-button
              >
            </c-box>
            <c-box ml="6">
              <c-button bg="csystem.pp" color="csystem.btx" size="sm"
                >Edit</c-button
              >
            </c-box>
          </c-flex>
          <c-divider orientation="horizontal" />
        </div>
        <!-- End of the Content -->
        <c-button bg="csystem.btn" color="csystem.btx" mt="6">Logout</c-button>
      </c-box>
    </c-flex>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { CFlex, CBox, CText, CDivider, CButton } from '@chakra-ui/vue'

export default Vue.extend({
  name: 'Panel',
  components: {
    CFlex,
    CBox,
    CText,
    CDivider,
    CButton,
  },
})
</script>
