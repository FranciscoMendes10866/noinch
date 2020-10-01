<template>
  <div>
    <a-table :columns="columns" :data-source="managers" bordered>
      <template slot="delete" slot-scope="text, record">
        <a-popconfirm
          v-if="managers.length"
          title="Sure to delete?"
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
    dataIndex: 'website',
    key: 'website',
  },
  {
    title: 'Password',
    dataIndex: 'wpass',
    key: 'wpass',
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

const managers = [
  {
    id: 1,
    website: 'www.facebook.com',
    wpass: 'dadsdas8767A',
  },
  {
    id: 2,
    website: 'www.youtube.com',
    wpass: 'dsasda76UY',
  },
]

interface Record {
  id: number
  website: string
  wepass: string
}

export default Vue.extend({
  data: () => ({
    columns,
    managers,
  }),
  methods: {
    Delete(record: Record) {
      // eslint-disable-next-line no-console
      console.log(record)
    },
  },
})
</script>
