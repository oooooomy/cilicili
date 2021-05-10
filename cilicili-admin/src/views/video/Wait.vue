<template>
  <div>
    <a-table
      :loading="loading"
      :columns="columns"
      :data-source="data"
      rowKey="id"
    >
      <span slot="customTitle"><a-icon type="smile-o" /> title </span>
      <span slot="userNickname" slot-scope="userNickname">
        <a-tag color="cyan">
          {{ userNickname }}
        </a-tag>
      </span>
      <span slot="poster" slot-scope="poster">
        <a :href="poster" />
      </span>
      <span slot="url" slot-scope="url">
        <a :href="url" />
      </span>
      <span slot="action" slot-scope="text, record">
        <a @click="onclickPoster(record)">封面</a>
        <a-divider type="vertical" />
        <a @click="onclickVideo(record)">视频</a>
        <a-divider type="vertical" />
        <a-popconfirm
          title="你确定已经预览过，并且通过审核?"
          ok-text="Yes"
          cancel-text="No"
          @confirm="confirm(record)"
        >
          <a href="#">审核通过</a>
        </a-popconfirm>
      </span>
    </a-table>
    <a-modal
      v-model="posterVisible"
      title="封面预览"
      @ok="posterVisible = false"
    >
      <img
        class="poster"
        :src="base + '/upload-service/file/' + openVideo.poster"
        alt=""
      />
    </a-modal>
    <a-modal v-model="videoVisible" title="视频预览" @ok="videoVisible = false">
      <video width="300" height="240" controls>
        <source
          :src="base + '/upload-service/file/' + openVideo.url"
          type="video/mp4"
        />
      </video>
    </a-modal>
  </div>
</template>
<script>
import { FindVideoByStatus, SetVideoStatus } from "../../api/video";
import { BASE_GATEWAY_URL } from "../../utils/request";

const columns = [
  {
    dataIndex: "title",
    key: "title",
    slots: { title: "customTitle" },
    scopedSlots: { customRender: "title" },
  },
  {
    title: "userNickname",
    key: "userNickname",
    dataIndex: "userNickname",
    scopedSlots: { customRender: "userNickname" },
  },
  {
    title: "poster",
    dataIndex: "poster",
    key: "poster",
    width: "30%",
  },
  {
    title: "url",
    dataIndex: "url",
    key: "url",
    width: "30%",
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
      openVideo: {},
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
      FindVideoByStatus("false").then((res) => {
        if (res.code === 200) {
          setTimeout(() => {
            this.data = res.data;
            this.loading = false;
          }, 500);
        }
      });
    },

    confirm(record) {
      console.log(record);
      SetVideoStatus(record.id, "true").then((res) => {
        if (res.code === 200) {
          this.load();
          this.$message.success("视频审核成功");
        } else {
          this.$message.error(res.msg);
        }
      });
    },

    onclickPoster(record) {
      this.posterVisible = true;
      this.openVideo = record;
    },

    onclickVideo(record) {
      this.videoVisible = true;
      this.openVideo = record;
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