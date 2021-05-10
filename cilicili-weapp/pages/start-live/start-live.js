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
      userInfo: App.globalData.userInfo
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
        }
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
    console.log(this.data.room)
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
        // tempFilePath可以作为img标签的src属性显示图片
        const tempFilePaths = res.tempFilePaths
        _this.setData({
          posterUrl: tempFilePaths
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

  onChangeSwitch({ detail }) {
    //震动反馈
    wx.vibrateShort({
      type: 'heavy'
    })
    // 需要手动对 checked 状态进行更新
    this.setData({ checked: detail });
  },

})