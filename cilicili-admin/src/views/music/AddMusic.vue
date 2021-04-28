<template>
  <div class="main">
    <a-form id="components-form-demo-validate-other" v-bind="formItemLayout">
      <a-form-item label="Plain Text">
        <span class="ant-form-text"> 添加视频音乐MP3 </span>
      </a-form-item>
      <a-form-item label="Music Name" has-feedback>
        <a-input v-model="musicForm.name" />
      </a-form-item>

      <a-form-item label="Music Author">
        <a-input v-model="musicForm.author" />
      </a-form-item>

      <a-form-item label="Select File" extra="请选择MP3格式的音频文件">
        <a-upload
          accept=".mp3"
          name="file"
          @change="uploadMusicChange"
          :multiple="true"
          :action="BASE_URL + '/file/mp3'"
          list-type="picture"
        >
          <a-button>
            <a-icon type="upload" />
            Click to upload
          </a-button>
        </a-upload>
      </a-form-item>

      <a-form-item label="封面海报">
        <div class="dropbox">
          <a-upload-dragger
            @change="uploadPosterChange"
            accept="image/jpeg"
            :multiple="true"
            name="file"
            :action="BASE_URL + '/file/jpg'"
          >
            <p class="ant-upload-drag-icon">
              <a-icon type="inbox" />
            </p>
            <p class="ant-upload-text">
              Click or drag file to this area to upload
            </p>
            <p class="ant-upload-hint">Support for a single or bulk upload.</p>
          </a-upload-dragger>
        </div>
      </a-form-item>

      <a-form-item :wrapper-col="{ span: 12, offset: 6 }">
        <a-button @click="submit" type="primary" html-type="submit">
          Submit
        </a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script>
import { BASE_GATEWAY_URL } from "../../utils/request";
import { SaveMusic } from "../../api/music";

export default {
  data: () => ({
    BASE_URL: BASE_GATEWAY_URL + "/upload-service",
    musicForm: {
      poster: "",
      name: "",
      author: "",
      url: "",
    },
    formItemLayout: {
      labelCol: { span: 6 },
      wrapperCol: { span: 14 },
    },
  }),

  beforeCreate() {
    this.form = this.$form.createForm(this, { name: "validate_other" });
  },

  methods: {
    nameInputChange(e) {
      this.musicForm.name += e.data.toString();
    },

    authorInputChange(e) {
      this.musicForm.author += e.data.toString();
    },

    submit() {
      console.log(this.musicForm);
      if (
        this.musicForm.poster &&
        this.musicForm.name &&
        this.musicForm.author &&
        this.musicForm.url
      ) {
        SaveMusic(this.musicForm).then((res) => {
          if (res.status) {
            this.$message.success("音乐保存成功");
            this.musicForm = {};
            this.$router.push("/music");
          } else {
            this.$message.error(res.msg);
          }
        });
      } else {
        this.$message.error("请完善表单");
      }
    },

    uploadMusicChange(e) {
      if (e.file.response) {
        const res = e.file.response;
        console.log("uploadMusicChange ", res);
        if (res.status) {
          this.musicForm.url = res.data;
          this.$message.success("MP3上传成功");
        } else {
          this.$message.error(res.msg);
        }
      }
    },

    uploadPosterChange(e) {
      if (e.file.response) {
        const res = e.file.response;
        console.log("uploadPosterChange ", res);
        if (res.status) {
          this.musicForm.poster = res.data;
          this.$message.success("封面海报上传成功");
        } else {
          this.$message.error(res.msg);
        }
      }
    },

    normFile(e) {
      console.log("Upload event:", e);
      if (Array.isArray(e)) {
        return e;
      }
      console.log(e);
      return e && e.fileList;
    },
  },
};
</script>
<style>
#components-form-demo-validate-other .dropbox {
  height: 180px;
  line-height: 1.5;
}

.main {
  background: #ffffff;
  padding: 20px;
  text-align: left;
}
</style>