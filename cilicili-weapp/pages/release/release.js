const App = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        base: App.globalData.baseUrl,
        musicList: [],
        showMusic: false,
        time: 30 * 1000,
        timeData: {},
        cameraAuthorize: false,
        over: false,
        src: '',
        start: false,
        videoSrc: '',
        poster: '',
        tempThumbPath: '',
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

    //选择使用的音乐
    selectMusic(e) {
        console.log(e.currentTarget.dataset.music)
        this.setData({
            showMusic: false,
            selectMusic: e.currentTarget.dataset.music,
        })
    },

    loadMusicList() {
        wx.request({
            url: App.globalData.baseUrl + "/upload-service/music",
            method: "GET",
            success: (res) => {
                this.setData({
                    musicList: res.data.data
                })
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
            wx.showLoading({
                title: '发起摄像头',
            })
            this.cameraContext.startRecord({
                success: () => {
                    this.setData({start: true})
                    wx.hideLoading({})
                },
                timeoutCallback: (res) => {
                    this.setData({
                        videoSrc: res.tempVideoPath,
                        tempThumbPath: res.tempThumbPath,
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
                    tempThumbPath: res.tempThumbPath,
                    start: false,
                    over: true,
                })
                wx.hideLoading({})
            }
        })
    },

    //点击发布视频
    onClickSubmitVidoe() {
        let videoForm = {
            url: '',
            poster: '',
            userId: App.globalData.userInfo.id,
            userNickname: App.globalData.userInfo.nickname,
            userAvatar: App.globalData.userInfo.avatarUrl,
            musicName: '',
            musicAuthor: '',
            musicPoster: '',
            musicUrl: '',
            musicId: '',
            useMusic: false,
            title: '',
            likeCount: 0,
            commentCount: 0,
            shareCount: 0,
        }
        wx.showLoading({
            title: '正在上传中',
        })
        wx.uploadFile({
            timeout: 1000 * 60 * 60 * 10,
            filePath: this.data.tempThumbPath,
            name: 'file',
            url: App.globalData.baseUrl + '/upload-service/file',
            success: (res) => {
                const parse = JSON.parse(res.data);
                videoForm.poster = parse.data
            }
        })
        wx.uploadFile({
            timeout: 1000 * 60 * 60 * 10,
            filePath: this.data.videoSrc,
            name: 'file',
            url: App.globalData.baseUrl + '/upload-service/file',
            success: (res) => {
                const data = JSON.parse(res.data)
                if (data.code === 200) {
                    wx.hideLoading({})
                    videoForm.url = data.data
                    if (this.data.selectMusic) {
                        videoForm.musicId = this.data.selectMusic.id
                        videoForm.musicName = this.data.selectMusic.name
                        videoForm.musicAuthor = this.data.selectMusic.author
                        videoForm.musicPoster = this.data.selectMusic.poster
                        videoForm.musicUrl = this.data.selectMusic.url
                        videoForm.useMusic = true
                    }
                    wx.navigateTo({
                        url: '/pages/post/post?videoForm=' + JSON.stringify(videoForm),
                        success: (res) => {
                            this.setData({
                                selectMusic: {}
                            })
                        }
                    })
                } else {
                    wx.showToast({
                        title: '上传失败',
                        icon: 'error',
                        duration: 2000
                    })
                    console.log("Request error: ", data.msg)
                }
            },
            fail: (err) => {
                wx.showToast({
                    title: '上传失败',
                    icon: 'error',
                    duration: 2000
                })
                console.log('文件上传失败： ', err)
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
    onClickSelectMusic() {
        this.loadMusicList()
        this.setData({
            showMusic: true
        })
    },

    onCloseMusic() {
        this.setData({
            showMusic: false
        })
    },


})