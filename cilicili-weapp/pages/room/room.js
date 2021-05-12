const App = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        roomId: '',
        base: App.globalData.baseUrl,
        navBarTop: 90,
        active: 1,
        videoHeight: 0,
        barrageInputValue: '',
        room: {},
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        this.loadRoom(options.id)
        let windowWidth = wx.getSystemInfoSync().windowWidth;
        this.setData({
            roomId: options.id,
            videoHeight: windowWidth / this.data.room.ratio
        })
        this.videoContext = wx.createVideoContext('player')
        this.videoContext.play()
        this.onConnWebsocketServer()
    },

    loadRoom(id){
        wx.request({
            url: App.globalData.baseUrl + "/room-service/findById/" + id,
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

    onConnWebsocketServer(){
        console.log('连接websocket服务器')
        console.log(this.data.base + '/room-service/ws/conn?roomId=' + this.data.roomId + '&userId=' + App.globalData.userInfo.id)
        wx.connectSocket({
          url: 'http://localhost:9881/ws/conn?roomId=' + this.data.roomId + '&userId=' + App.globalData.userInfo.id,
          header:{
            'content-type': 'application/json'
          },
          protocols: ['protocol1'],
          success: (res)=>{
              console.log('connectSocket success : ', res)
          },
          fail: (err)=>{
            console.log('connectSocket error : ', err)
          }
        })
    },

    onClickLeft() {
        wx.navigateBack({delta: 1})
    },

    barrageInput(e) {
        this.setData({
            barrageInputValue: e.detail.value
        })
    },

    clickSubmit() {
        this.videoContext.sendDanmu({
            text: 'user: ' + this.data.barrageInputValue,
            color: this.randomColor()
        })
        //this.setData({barrageInputValue: ''})
    },

    randomColor() {
        let r = Math.floor(Math.random() * 256);
        let g = Math.floor(Math.random() * 256);
        let b = Math.floor(Math.random() * 256);
        return '#' + r.toString(16) + g.toString(16) + b.toString(16);
    },

    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady: function () {
    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow: function () {
        this.setData({
            barrageList: [{username: '卢本伟', content: '欢迎来到我的直播间'}]
        })
    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide: function () {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload: function () {

    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh: function () {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom: function () {

    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage: function () {

    }

})