<template>
  <div>
    <c-flex direction="row" justify="space-between" mt="6">
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
    <div v-for="item in dbData" :key="item.id">
      <c-flex direction="row" justify="space-between" align="center" py="2">
        <c-box>
          <a href="" target="blank"
            ><c-text color="csystem.pp">{{ item.Website }}</c-text></a
          >
        </c-box>
        <c-box ml="6">
          <c-text color="csystem.pp">{{ item.WPass }}</c-text>
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
    <c-button bg="csystem.btn" color="csystem.btx" mt="6" @click="LogOut"
      >Logout</c-button
    >
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapActions } from 'vuex'
import { CFlex, CBox, CText, CDivider, CButton } from '@chakra-ui/vue'

export default Vue.extend({
  name: 'Table',
  components: {
    CFlex,
    CBox,
    CText,
    CDivider,
    CButton,
  },
  data: () => ({
    dbData: [],
  }),
  mounted() {
    this.Fetch()
  },
  methods: {
    ...mapActions({
      LogOut: 'LogOut',
    }),
    Fetch() {
      this.$axios
        .get('http://localhost:4444/api/v1/manager', {
          headers: {
            Authorization: `Bearer ${this.$store.state.Token}`,
          },
        })
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        .then((res: any) => {
          // eslint-disable-next-line no-console
          console.log(res)
          this.dbData = res.data
        })
        // eslint-disable-next-line no-console
        .catch((error: Error) => console.log(error))
    },
  },
})
</script>
