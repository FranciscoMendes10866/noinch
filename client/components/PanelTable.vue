<template>
  <div>
    <a-table :columns="columns" :data-source="managers" bordered>
      <template slot="delete" slot-scope="text, record">
        <a-popconfirm
          v-if="managers.length"
          title="Are you sure?"
          @confirm="() => Delete(record)"
        >
          <a-button type="danger" shape="circle" icon="delete" />
        </a-popconfirm>
      </template>
      <template v-slot:edit>
        <a-button type="primary" shape="circle" icon="edit" />
      </template>
    </a-table>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

const columns = [
  {
    title: 'Website',
    dataIndex: 'Website',
    key: 'Website',
  },
  {
    title: 'Password',
    dataIndex: 'WPass',
    key: 'WPass',
  },
  {
    title: 'Delete',
    key: 'delete',
    scopedSlots: { customRender: 'delete' },
  },
  {
    title: 'Edit',
    key: 'edit',
    scopedSlots: { customRender: 'edit' },
  },
]

interface Record {
  ID: number
  Website: string
  WPass: string
}

export default Vue.extend({
  data: () => ({
    columns,
    managers: [],
  }),
  mounted() {
    this.Fetch()
  },
  methods: {
    Fetch() {
      this.$axios
        .get('http://localhost:4444/api/v1/manager', {
          headers: {
            Authorization: `Bearer ${this.$store.state.Token}`,
          },
        })
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        .then((res: any) => {
          this.managers = res.data
        })
        // eslint-disable-next-line no-console
        .catch((error: Error) => console.log(error))
    },
    Delete(record: Record) {
      this.$axios
        .delete(`http://localhost:4444/api/v1/manager/${record.ID}`, {
          headers: {
            Authorization: `Bearer ${this.$store.state.Token}`,
          },
        })
        .then(() => this.Fetch())
        // eslint-disable-next-line no-console
        .catch((error: Error) => console.log(error))
    },
  },
})
</script>
