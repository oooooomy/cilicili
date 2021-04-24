<template>
  <div class="login-box">
    <div>
      <div class="box-header">
        <img class="logo" src="../assets/logo.svg" alt=""/>
        <div class="box-header-t">cilicili - Admin</div>
      </div>
      <div class="title">cilicili - 管理员登录 请妥善保存您的账户信息</div>
      <a-tabs @change="tabClick" default-active-key="1" :tabBarStyle="{ textAlign: 'center' }">
        <a-tab-pane key="1" tab="账号密码登陆">
          <a-input
              v-model="form.username"
              size="large"
              style="margin-top: 10px"
              class="input"
              placeholder="邮箱">
            <a-icon slot="prefix" type="mail"/>
          </a-input>
          <a-input-password
              v-model="form.password"
              size="large"
              class="input"
              placeholder="密码">
            <a-icon slot="prefix" type="lock"/>
          </a-input-password>
        </a-tab-pane>
        <a-tab-pane key="2" tab="邮箱登陆" force-render>
          <a-input
              v-model="emailForm.email"
              size="large"
              style="margin-top: 10px"
              class="input"
              placeholder="邮箱">
            <a-icon slot="prefix" type="mail"/>
          </a-input>
          <div style="display: flex">
            <a-input
                v-model="emailForm.code"
                size="large"
                class="input"
                placeholder="验证码">
              <a-icon slot="prefix" type="safety-certificate"/>
            </a-input>
            <a-button class="code-btn">
              <a-icon type="message"/>
              获取验证码
            </a-button>
          </div>
        </a-tab-pane>
      </a-tabs>
      <div style="margin-bottom: 20px">
        <a-checkbox>自动登录</a-checkbox>
        <a-button type="link" style="margin-left: 158px">
          忘记密码 ?
        </a-button>
      </div>
      <a-button :loading="submitLoading" class="submit-btn" type="primary" @click="submitLogin">
        确认登陆
      </a-button>
      <div class="des">cilicili - Copyright 2021 高元明</div>
    </div>
  </div>
</template>

<script>
export default {

  data() {
    return {
      submitType: '1',
      submitLoading: false,
      form: {username: '', password: '',},
      emailForm: {email: '', code: ''},
      remember: false
    }
  },

  methods: {

    submitLogin() {
      this.submitLoading = true
      let result = this.submitType === '1' ? this.checkUsernameAndPwd() : this.checkEmailAndCode()
      if (result) {
        const user = {
          username: '1729677089@qq.com',
          id: "1",
          avatar: 'https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png'
        }
        const token = "ja9s-a3a2-j38i-991j-9912-kkd8"
        //保存登录数据
        this.$store.commit('user/saveToken', token)
        this.$store.commit('user/saveLoginUser', user)
        setTimeout(() => {
          this.$router.push("/")
          this.submitLoading = false
        }, 700)
      } else {
        setTimeout(() => {
          this.submitLoading = false
        }, 700)
      }
    },

    tabClick(key) {
      this.submitType = key
    },

    checkUsernameAndPwd() {
      if (this.form.username === 'admin' && this.form.password === 'admin') {
        return true
      } else {
        this.$message.error('用户名或密码错误');
        return false
      }
    },

    checkEmailAndCode() {
      const emailRegex = new RegExp('^\\w{3,}(\\.\\w+)*@[A-z0-9]+(\\.[A-z]{2,5}){1,2}$')
      const codeRegex = new RegExp('^[1-9]\\d{0,6}$')
      if (!emailRegex.test(this.emailForm.email) || !codeRegex.test(this.emailForm.code)) {
        this.$message.error('请输入正确格式的邮箱或验证码');
        return false
      } else {
        return true
      }
    },

  }

}
</script>

<style scoped>

body {
  background: #000000 !important;
}

>>> .ant-tabs-bar {
  border-bottom: none !important;
}

>>> .ant-btn-primary {
  border-color: #5a84fd;
}

.login-box {
  width: 350px;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}

.box-header {
  display: flex;
}

.box-header-t {
  font-weight: 900;
  font-size: 30px;
}

.logo {
  width: 44px;
  height: 44px;
  font-weight: 900;
  margin-right: 20px;
  margin-left: 25px;
}

>>> .ant-tabs-nav {
  width: 350px;
}

>>> .ant-tabs-ink-bar {
  left: 52px;
}

>>> .ant-input-affix-wrapper .ant-input {
  font-size: 12px !important;
}

.title {
  color: #91949c;
  padding-top: 15px;
  padding-bottom: 35px;
  font-size: 13px;
  text-align: center;
}

.input {
  margin-bottom: 25px;
  font-size: 10px;
}

.code-btn {
  height: 40px;
  margin-left: 30px;
}

.submit-btn {
  letter-spacing: 2px;
  background: #5a84fd;
  width: 100%;
  height: 45px;
}

.des {
  padding-top: 25px;
  font-size: 13px;
  text-align: center;
  color: #91949c;
  letter-spacing: 2px;
}
</style>