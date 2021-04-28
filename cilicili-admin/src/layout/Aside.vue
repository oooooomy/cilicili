<template>
  <a-layout-sider
      :style="{ overflow: 'auto', height: '100vh', position: 'fixed', left: 0 }"
      width="220">
    <div class="logo">
      <h5>cilicili</h5>
    </div>
    <a-menu theme="dark" mode="inline">
      <a-sub-menu v-for="(item, index) in menus" :key="index">
        <span slot="title">
          <a-icon :type="item.icon"/>
          <span>{{ item.title }}</span>
        </span>
        <a-menu-item v-for="menu in item.children" :key="menu.title">
          <router-link :to="menu.path">
            {{ menu.title }}
          </router-link>
        </a-menu-item>
      </a-sub-menu>
    </a-menu>
  </a-layout-sider>
</template>

<script>

export default {

  data() {
    return {
      menus: [
        {
          title: '系统监控',
          icon: 'pie-chart',
          children: [{title: '服务实例', path: '/'}]
        },
        {
          title: '音乐管理',
          icon: 'sound',
          children: [{title: '音乐列表', path: '/music'},{title: '添加音乐', path: '/add'}]
        },
        {
          title: '视频审核',
          icon: 'play-square',
          children: [
            {title: '等待审核', path: '/wait'},
            {title: '复核视频', path: '/videos'}
          ]
        },
        {
          title: '直播管理',
          icon: 'hourglass',
          children: [
            {title: '正在直播', path: '/live'},
            {title: '资质审核', path: '/audit'},
          ]
        },
        {
          title: '系统设置',
          icon: 'setting',
          children: [
            {title: '我的账号', path: '/account'},
            {title: '管理员账号', path: '/worker'},
          ]
        },
      ]
    }
  },

}

</script>

<style scoped>

#components-layout-demo-fixed-sider .logo {
  padding: 10px 10px;
  height: 50px;
  font-size: 15px;
  margin: 12px;
  color: #ffffff;
  letter-spacing: 2px;
}

.ant-menu {
  letter-spacing: 1px;
}

.logo {
  display: flex;
}

.logo img {
  width: 35px;
  height: 35px;
  margin-right: 5px;
}

.logo h5 {
  color: #1890ff;
  font-weight: 900;
  letter-spacing: 2px;
  padding-left: 10px;
  font-size: 35px;
  transform: translateY(-5px);
}

</style>