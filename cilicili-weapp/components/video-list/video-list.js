const App = getApp()

Component({
    /**
     * 组件的属性列表
     */
    properties: {
        videoList: {
            type: Array
        }
    },

    /**
     * 组件的初始数据
     */
    data: {
        videoList: [],
        current: 0,
        showShare: false,
        options: [
            {name: '微信', icon: 'wechat', openType: 'share'},
            {name: '微博', icon: 'weibo'},
            {name: '复制链接', icon: 'link'},
            {name: '分享海报', icon: 'poster'},
            {name: '二维码', icon: 'qrcode'},
        ],
        commentList: [],
        selectCommentVideoId: '',
        base: App.globalData.baseUrl,
        barrageInputValue: '',
        isReady: false,
        showComment: false,
    },

    // 在组件实例进入页面节点树时执行
    attached: function () {
        this.videoContext = wx.createVideoContext('0', this)
        this.videoContext.play()
    },

    // 在组件实例被从页面节点树移除时执行
    detached: function () {
        console.log("组件实例被从页面节点树移除")
    },

    /**
     * 组件的方法列表
     */
    methods: {

        openShare() {
            this.setData({
                showShare: true
            })
        },

        onCloseShare() {
            this.setData({
                showShare: false
            })
        },

        swiperChange(e) {
            if (e.detail.source === 'touch') {
                this.videoContext.pause()
                this.videoContext = wx.createVideoContext(e.detail.current.toString(), this)
                this.videoContext.play()
                this.setData({
                    current: e.detail.current
                })
            }
        },

        openUserDetails(e) {
            wx.navigateTo({
                url: '/pages/user/user?id=' + e.currentTarget.dataset.id
            })
        },

        onClickLike(e) {
            let index = e.currentTarget.dataset.videoid
            wx.vibrateShort({type: 'heavy'})
            let form = {
                userId: App.globalData.userInfo.id,
                videoId: e.currentTarget.dataset.videoid,
            }
            wx.request({
                url: this.data.base + '/upload-service/video/like',
                method: 'POST',
                data: form,
                success: (res) => {
                    if (res.data.code === 200) {
                        let count = "videoList[" + this.data.current + "].likeCount"
                        this.setData({
                            [count]: this.data.videoList[this.data.current].likeCount+1
                        })
                        wx.showToast({
                            title: 'Thank you',
                            icon: 'success',
                            duration: 2000
                        })
                    } else {
                        wx.showToast({
                            title: res.data.msg,
                            icon: 'error',
                            duration: 2000
                        })
                    }
                }
            })
        },

        openComment(e) {
            this.setData({
                showComment: true,
                selectCommentVideoId: e.currentTarget.dataset.videoid,
            })
            wx.request({
                url: this.data.base + '/upload-service/comment/video/' + this.data.selectCommentVideoId,
                method: 'GET',
                success: (res) => {
                    if (res.data.code === 200) {
                        this.setData({
                            commentList: res.data.data
                        })
                        console.log(this.data.commentList)
                    }
                }
            })
        },

        onCloseComment() {
            this.setData({
                showComment: false
            })
        },

        barrageInput(e) {
            this.setData({
                barrageInputValue: e.detail.value
            })
        },

        clickSubmit() {
            if (!this.data.barrageInputValue) {
                wx.showToast({
                    title: '请输入内容',
                    duration: 2000,
                    icon: 'error'
                })
                return
            }
            console.log("send comment = ", this.data.barrageInputValue)
            let form = {
                videoId: this.data.selectCommentVideoId,
                userId: App.globalData.userInfo.id,
                userNickname: App.globalData.userInfo.nickname,
                userAvatar: App.globalData.userInfo.avatarUrl,
                content: this.data.barrageInputValue,
            }
            wx.request({
                url: this.data.base + '/upload-service/comment',
                method: 'POST',
                data: form,
                success: (res) => {
                    if (res.data.code === 200) {
                        let count = "videoList[" + this.data.current + "].commentCount"
                        this.setData({
                            barrageInputValue: '',
                            [count]: this.data.videoList[this.data.current].commentCount+1
                        })
                        wx.showToast({
                            title: "发送成功",
                            icon: 'success',
                            duration: 2000
                        })
                    }
                },
            })
        },

    }

})
