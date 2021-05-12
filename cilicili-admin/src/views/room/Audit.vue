<template>
  <div>
    <a-table
      :loading="loading"
      :columns="columns"
      :data-source="data"
      rowKey="id"
    >
      <span slot="customTitle"><a-icon type="smile-o" /> title </span>
      <span slot="email" slot-scope="email">
        <a-tag color="cyan">
          {{ email }}
        </a-tag>
      </span>
      <span slot="posterUrl" slot-scope="posterUrl">
        <a :href="posterUrl" />
      </span>
      <span slot="token" slot-scope="token">
        <a :href="token" />
      </span>
      <span slot="status" slot-scope="status">
        <a-tag :color="status === 1 ? '#87d068' : '#f50'">
          {{ status === 1 ? "已通过" : "未通过" }}
        </a-tag>
      </span>
      <span slot="action" slot-scope="text, record">
        <a @click="onclickPoster(record)">封面</a>
        <a-divider type="vertical" />
        <a-popconfirm
          v-if="record.status === 0"
          title="我已确定要此直播间符合审核要求!"
          ok-text="Yes"
          cancel-text="No"
          @confirm="confirm(record)"
        >
          <a>通过</a>
        </a-popconfirm>
        <a-popconfirm
          v-if="record.status === 1"
          title="确定要撤销此直播间的审核?"
          ok-text="Yes"
          cancel-text="No"
          @confirm="confirm(record)"
        >
          <a>撤销</a>
        </a-popconfirm>
      </span>
    </a-table>
    <a-modal
      v-model="posterVisible"
      title="直播间封面海报预览"
      @ok="posterVisible = false"
    >
      <img
        class="poster"
        :src="base + '/upload-service/file/' + openRoom.posterUrl"
        alt=""
      />
    </a-modal>
  </div>
</template>
<script>
import { FindAllRooms, SaveRoom } from "../../api/room";
import { BASE_GATEWAY_URL } from "../../utils/request";

const columns = [
  {
    dataIndex: "title",
    key: "title",
    slots: { title: "customTitle" },
    scopedSlots: { customRender: "title" },
  },
  {
    title: "email",
    key: "email",
    dataIndex: "email",
    scopedSlots: { customRender: "email" },
  },
  {
    title: "posterUrl",
    dataIndex: "posterUrl",
    key: "posterUrl",
    width: "20%",
  },
  {
    title: "status",
    key: "status",
    dataIndex: "status",
    scopedSlots: { customRender: "status" },
  },
  {
    title: "token",
    dataIndex: "token",
    key: "token",
    width: "20%",
  },
  {
    title: "Action",
    key: "action",
    scopedSlots: { customRender: "action" },
  },
];

const data = [];

export default {
  data() {
    return {
      base: BASE_GATEWAY_URL,
      openRoom: {},
      posterVisible: false,
      videoVisible: false,
      data,
      columns,
      loading: false,
    };
  },

  mounted() {
    this.load();
  },

  methods: {
    load() {
      this.loading = true;
      FindAllRooms().then((res) => {
        if (res.code === 200) {
          setTimeout(() => {
            this.data = res.data;
            this.loading = false;
          }, 500);
        }
      });
    },

    confirm(record) {
      record.status = record.status === 0 ? 1 : 0;
      SaveRoom(record).then((res) => {
        if (res.code === 200) {
          this.load();
          this.$message.success("Request ok");
        } else {
          this.$message.error(res.msg);
        }
      });
    },

    onclickPoster(record) {
      this.posterVisible = true;
      this.openRoom = record;
      console.log(
        "onclickPoster : ",
        this.base + "/upload-service/file/" + this.openRoom.posterUrl
      );
    },

    onclickVideo(record) {
      this.videoVisible = true;
      this.openRoom = record;
    },
  },
};
</script>

<style scoped>
.poster {
  width: 200px;
  margin-left: 50%;
  transform: translateX(-100px);
}
</style>