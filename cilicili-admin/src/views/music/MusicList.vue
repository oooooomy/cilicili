<template>
  <a-table :loading="loading" :columns="columns" :data-source="data" rowKey="id">
    <a slot="name" slot-scope="text">{{ text }}</a>
    <span slot="customTitle"><a-icon type="smile-o"/> name</span>
    <span slot="author" slot-scope="author">
      <a-tag color="cyan">
        {{ author }}
      </a-tag>
    </span>
    <span slot="poster" slot-scope="poster">
      <a :href="poster"/>
    </span>
    <span slot="url" slot-scope="url">
      <a :href="url"/>
    </span>
    <span slot="action" slot-scope="text, record">
      <a-popconfirm
          title="Are you sure delete this music?"
          ok-text="Yes"
          cancel-text="No"
          @confirm="confirm(record)"
      >
      <a href="#">Delete</a>
    </a-popconfirm>
    </span>
  </a-table>
</template>
<script>
import {DeleteMusic, FindAllMusic} from "../../api/music";

const columns = [
  {
    dataIndex: 'name',
    key: 'name',
    slots: {title: 'customTitle'},
    scopedSlots: {customRender: 'name'},
  },
  {
    title: 'author',
    key: 'author',
    dataIndex: 'author',
    scopedSlots: {customRender: 'author'},
  },
  {
    title: 'poster',
    dataIndex: 'poster',
    key: 'poster',
    width: '30%',
  },
  {
    title: 'url',
    dataIndex: 'url',
    key: 'url',
    width: '30%',
  },
  {
    title: 'Action',
    key: 'action',
    scopedSlots: {customRender: 'action'},
  },
];

const data = [];

export default {
  data() {
    return {
      data,
      columns,
      loading: false,
    };
  },

  mounted() {
    this.load()
  },

  methods: {

    load() {
      this.loading = true
      FindAllMusic().then((res) => {
        if (res.status) {
          setTimeout(() => {
            this.data = res.data
            this.loading = false
          }, 500)
        }
      })
    },

    confirm(record) {
      console.log(record);
      this.$message.success('Delete success');
      DeleteMusic(record.id).then((res) => {
        if (res.status) {
          this.load()
          this.$message.success("音乐删除成功")
        } else {
          this.$message.error(res.msg);
        }
      })
    },

  },

};
</script>
