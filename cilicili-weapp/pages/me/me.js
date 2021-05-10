const App = getApp()

Component({

    data: {
        base: App.globalData.baseUrl,
        list: [{}, {}, {}, {}, {}, {}, {}, {}, {}],
        userInfo: {},
        active: 0,
        uploadList: [],
        likeList: [],
    },

    // 在组件实例进入页面节点树时执行
    attached: function () {
        console.log("Me page detached")
        this.setData({
            userInfo: App.globalData.userInfo
        })
        this.getUserInfo()
        this.getMyVideos()
    },

    // 在组件实例被从页面节点树移除时执行
    detached: function () {
    },

    methods: {

        onClickStartLive(){
            wx.navigateTo({
              url: '/pages/start-live/start-live',
            })
        },

        getMyVideos() {
            wx.request({
                url: this.data.base + "/upload-service/video/findUserLikeAndUpload/" + App.globalData.userInfo.id,
                method: 'GET',
                success: (res) => {
                    if (res.data.code === 200) {
                        this.setData({
                            uploadList: res.data.data.uploadList ? res.data.data.uploadList : '',
                            likeList: res.data.data.likeList ? res.data.data.likeList : '',
                        })
                    }
                }
            })
        },

        getUserInfo() {
            let user = App.globalData.userInfo
            if (!user.hasInfo) {
                wx.getUserProfile({
                    desc: '用于完善会员资料', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
                    success: (res) => {
                        const info = res.userInfo
                        user.avatarUrl = info.avatarUrl
                        user.hasInfo = true
                        user.nickname = info.nickName
                        user.address = info.country + ' - ' + info.province + " - " + info.city
                        user.gender = info.gender === 0 ? '女' : '男'
                        console.log(user)
                        wx.request({
                            url: App.globalData.baseUrl + '/account-service/user/update',
                            method: "PUT",
                            data: user,
                            success: (res) => {
                                if (res.data.code === 200) {
                                    App.globalData.userInfo = res.data.data
                                    this.setData({
                                        userInfo: res.data.data
                                    })
                                }
                            }
                        })
                    },
                    fail: (res) => {
                        wx.navigateTo({
                            url: '/pages/index/index'
                        })
                    },
                })
            }
        },

        onClickUpload() {
            wx.navigateTo({url: '/pages/release/release'})
        },

        onLike() {
            wx.navigateTo({
                url: '/pages/videos/videos?videos=' + JSON.stringify(this.data.likeList)
            })
        },

        onUpload() {
            wx.navigateTo({
                url: '/pages/videos/videos?videos=' + JSON.stringify(this.data.uploadList)
            })
        },

    },

})