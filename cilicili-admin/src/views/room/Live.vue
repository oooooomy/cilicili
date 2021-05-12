<template>
  <div>
    <a-spin :spinning="loading" size="large">
      <a-row :gutter="16">
        <a-col
          class="gutter-row"
          v-for="(item, index) in rooms"
          :key="index"
          :span="6"
        >
          <div class="room-card">
            <a>
              <img
                @click="openRoom"
                :src="base + '/upload-service/file/' + item.posterUrl"
                alt=""
              />
            </a>
            <div class="room-title">Title: {{ item.title }}</div>
            <div>
              <a-button @click="stopLiving" type="link">
                暂停直播画面
              </a-button>
            </div>
          </div>
        </a-col>
      </a-row>
    </a-spin>

    <a-modal v-model="roomVisible" title="视频预览" @ok="roomVisible = false">
      <video width="100%" height="240" controls>
        <source :src="base + '/upload-service/file/'" type="video/mp4" />
      </video>
    </a-modal>
  </div>
</template>

<script>
import { FindAllRoomsByLiving } from "../../api/room";
import { BASE_GATEWAY_URL } from "../../utils/request";
export default {
  data() {
    return {
      base: BASE_GATEWAY_URL,
      rooms: [],
      loading: false,
      roomVisible: false,
    };
  },

  mounted() {
    this.load();
  },

  methods: {
    load() {
      this.loading = true;
      FindAllRoomsByLiving("true").then((res) => {
        if (res.code === 200) {
          this.rooms = res.data;
          setTimeout(() => {
            this.loading = false;
          }, 600);
        }
      });
    },

    openRoom() {
      this.roomVisible = true;
    },

    stopLiving() {
      alert("stop");
    },
  },
};
</script>

<style scoped>
.room-card {
  width: 100%;
  height: 270px;
  border-radius: 15px;
  background: #ffffff;
  margin-bottom: 20px;
  letter-spacing: 1px;
}

.room-card img {
  width: 100%;
  height: 180px;
  border-top-left-radius: 15px;
  border-top-right-radius: 15px;
}

.room-title {
  font-size: 15px;
  font-weight: 500;
  color: rgba(0, 0, 0, 0.8);
  padding: 15px 15px 8px 15px;
}
</style>