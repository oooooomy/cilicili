const App = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        showMusic: false,
        time: 30 * 1000,
        timeData: {},
        cameraAuthorize: false,
        over: false,
        src: '',
        start: false,
        videoSrc: '',
        cameraData: {
            flash: '',
            devicePosition: true,
        },
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        //查看是否打开摄像头权限
        wx.getSetting({
            success: (res) => {
                if (!res.authSetting['scope.camera']) {
                    wx.showModal({
                        title: '提示',
                        content: '是否打开设置页面，用来开启你的摄像头权限',
                        success(res) {
                            if (res.confirm) {
                                wx.openSetting({})
                            }
                        }
                    })
                } else {
                    this.setData({cameraAuthorize: true})
                    this.cameraContext = wx.createCameraContext();
                }
            }
        })
    },

    onChange(e) {
        this.setData({
            timeData: e.detail,
        });
    },

    //摄像头打开异常
    openCameraError() {
        console.log('用户未打开摄像头')
    },

    //打开震动反馈
    openVibrate() {
        wx.vibrateShort({
            type: 'heavy'
        })
    },

    onClickLeft() {
        wx.navigateBack({delta: 1})
    },

    //翻转摄像头
    onClickReversal() {
        this.openVibrate()
        this.setData({
            cameraData: {devicePosition: !this.data.cameraData.devicePosition,},
        })
    },

    //点击开始录制
    onClickStart() {
        if (this.data.cameraAuthorize) {
            this.openVibrate()
            this.cameraContext.startRecord({
                success: () => {
                    this.setData({start: true})
                    setTimeout(() => {
                        wx.showLoading({title: '加载中', mask: true})
                    }, 30 * 1000)
                },
                timeoutCallback: (res) => {
                    this.setData({
                        videoSrc: res.tempVideoPath,
                        start: false,
                        over: true,
                    })
                    wx.hideLoading({})
                }
            })
        }
    },

    //结束录像
    stopRecord() {
        this.openVibrate()
        wx.showLoading({title: '加载中', mask: true})
        this.cameraContext.stopRecord({
            success: (res) => {
                this.setData({
                    videoSrc: res.tempVideoPath,
                    start: false,
                    over: true,
                })
                wx.hideLoading({})
            }
        })
    },

    //点击发布视频
    onClickSubmitVidoe(){
        wx.showLoading({
          title: '正在上传中',
        })
        wx.uploadFile({
          filePath: this.data.videoSrc,
          name: 'file',
          url: App.globalData.baseUrl +  '/video/upload/mp4',
          success: (res)=>{
              console.log(res.data.data)
              wx.hideLoading({})
          },

          fail: (err) => {
              console.log(err)
          }
        })
    },

    //重新拍摄
    reShoot() {
        this.setData({over: false, videoSrc: ''})
    },

    openSetting() {
        wx.openSetting()
    },

    //点击选择音乐
    onClickSelectMusic(){
        this.setData({
            showMusic: true
        })
    },

    onCloseMusic(){
        this.setData({
            showMusic: false
        })
    },


})