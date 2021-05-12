const App = getApp()

Page({

  /**
   * 页面的初始数据
   */
  data: {
    base: App.globalData.baseUrl,
    value: '',
    checked: false,
    type: '',
    userInfo: {},
    showTitle: false,
    room: {
      id: '',
      username: '',
      avatar: '',
      isLiving: false,
      title: '',
      email: '',
      posterUrl: '',
    },
    tempFilePaths: '',
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    this.setData({
      userInfo: App.globalData.userInfo,
    })
    this.loadMyLiveForm()
  },

  loadMyLiveForm() {
    wx.request({
      url: App.globalData.baseUrl + "/room-service/findById/" + App.globalData.userInfo.id,
      method: "GET",
      success: (res) => {
        const r = res.data
        if (r.code === 200) {
          this.setData({
            room: r.data
          })
        }else{
          wx.showToast({
            title: r.msg,
          })
          const room = {
            id: App.globalData.userInfo.id,
            username: App.globalData.userInfo.nickname,
            avatar: App.globalData.userInfo.avatarUrl,
          }
          this.setData({
            room: room
          })
          wx.request({
            url: App.globalData.baseUrl + '/room-service/create',
            method: 'POST',
            data: room,
            success: (res) => {
              wx.showToast({ title: '初始化直播间' })
            }
          })
        }
      }
    })
  },

  //上播
  changeLivingUp() {
    let _this = this
    let upUrl = _this.data.base + '/room-service/up/' + this.data.room.id
    wx.showModal({
      title: '提示',
      content: '我已阅读用户直播协议，审核结果将以邮件形式发送给你',
      success(res) {
        if (res.confirm) {
          let room = _this.data.room
          room.isLiving = true
          wx.request({
            url: upUrl,
            method: 'GET',
            success: (res)=>{
              WSH.showToast({
                title: "开播成功"
              })
            }
          })
          _this.setData({
            room: room
          })
          _this.saveRoomEdit()
        }
      }
    })
  },

  //下播
  changeLivingDown() {
    let room = this.data.room
    room.isLiving = false
    this.setData({
      room: room
    })
    this.saveRoomEdit()
  },

  saveRoomEdit() {
    console.log(this.data.room)
    wx.request({
      url: App.globalData.baseUrl + '/room-service/create',
      method: 'POST',
      data: this.data.room,
      success: (res) => {
        wx.showToast({ title: '保存成功' })
      }
    })
  },

  onClickClose() {
    console.log("关闭直播")
  },

  onSubmit(e) {
    let room = this.data.room
    if (this.data.type === 'Title') {
      room.title = e.detail.value.Title
      this.setData({ room: room })
    }
    if (this.data.type === 'Email') {
      room.email = e.detail.value.Email
      this.setData({ room: room })
    }
    this.saveRoomEdit()
    this.onClose()
  },

  startLive(e) {
    let form = {
      id: App.globalData.userInfo.id,
      title: e.detail.value.title,
      email: e.detail.value.email,
      posterUrl: this.data.posterUrl[0],
    }
    console.log("form:", form)
    wx.uploadFile({
      timeout: 1000 * 60 * 60 * 10,
      filePath: form.posterUrl,
      name: 'file',
      url: App.globalData.baseUrl + '/upload-service/file',
      success: (res) => {
        const r = JSON.parse(res.data)
        if (r.code === 200) {
          form.posterUrl = r.data
          wx.request({
            url: App.globalData.baseUrl + '/room-service/create',
            method: 'POST',
            data: form,
            success: (res) => {
              console.log(res.data)
            }
          })
        }
      },
      fail: (err) => {
        console.log(err)
      }
    })
  },

  onClickLeft() {
    wx.navigateBack({ delta: 1 })
  },

  onClickPoster() {
    let _this = this
    wx.chooseImage({
      count: 1,
      sizeType: ['original', 'compressed'],
      sourceType: ['album', 'camera'],
      success(res) {
        // tempFilePath可以作为img标签的src属性显示图片
        const tempFilePaths = res.tempFilePaths
        _this.setData({
          tempFilePaths: tempFilePaths
        })
      }
    })
  },

  onClose() {
    this.setData({
      showTitle: false,
    })
  },

  openChoseImage() {
    let _this = this
    wx.chooseImage({
      count: 1,
      sizeType: ['original', 'compressed'],
      sourceType: ['album', 'camera'],
      success(res) {
        let room = _this.data.room
        wx.uploadFile({
          timeout: 1000 * 60 * 60 * 10,
          filePath: res.tempFilePaths[0],
          name: 'file',
          url: App.globalData.baseUrl + '/upload-service/file',
          success: (res2) => {
            const r = JSON.parse(res2.data)
            console.log(r)
            if (r.code === 200) {
              room.posterUrl = r.data
              _this.saveRoomEdit()
            }
          },
        })
      }
    })
  },

  openPopup(e) {
    const type = e.currentTarget.dataset.type
    const value = type === 'Title' ? this.data.room.title : this.data.room.email
    this.setData({
      value: value
    })
    this.setData({
      showTitle: true,
      type: type
    })
  },

})