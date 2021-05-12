const App = getApp()

Component({

    data: {
        base: App.globalData.baseUrl,
        current: 0,
        active: 0,
        posterList: [{}, {}, {}, {}],
        roomList: []
    },

    // 在组件实例进入页面节点树时执行
    attached: function () {
        this.loadRooms()
    },

    // 在组件实例被从页面节点树移除时执行
    detached: function () {
        console.log("Live page detached")
    },

    methods: {

        loadRooms(){
            wx.request({
              url: this.data.base + '/room-service/findByLiving/true',
              method: 'GET',
              success: (res)=>{
                  const r = res.data
                  if(r.code === 200){
                    this.setData({
                        roomList: r.data
                    })
                  }
              }
            })
        },

        clickRoom(e) {
            const roomId = e.currentTarget.dataset.roomid
            wx.navigateTo({
                url: '/pages/room/room?id=' + roomId
            })
        },

        swiperChange(e) {
            this.setData({current: e.detail.current})
        },
    },

})