const App = getApp()

Page({

    data: {
        base: App.globalData.baseUrl,
        likeList: [],
        uploadList: [],
        userId: '',
        active: 0,
        user: {},
    },

    onLoad: function (options) {
        this.setData({
            userId: options.id
        })
        console.log(options.id)
        wx.request({
            url: this.data.base + "/account-service/getInfo/" + options.id,
            method: 'GET',
            success: (res) => {
                const r = res.data
                if (r.status) {
                    this.setData({user: r.data})
                }
            }
        })
        this.getMyVideos()
    },

    onClickFollow(){
        wx.vibrateShort({type: 'heavy'})
        let form = {
            fromUserId: App.globalData.userInfo.id,
            toUserId: this.data.userId,
        }
        wx.request({
            url: this.data.base + '/account-service/follow',
            method: 'POST',
            data: form,
            success: (res) => {
                if (res.data.status) {
                    wx.showToast({
                        title: 'Thank you',
                        icon: 'success',
                        duration: 2000
                    })
                }else {
                    wx.showToast({
                        title: res.data.msg,
                        icon: 'error',
                        duration: 2000
                    })
                }
            }
        })
    },

    getMyVideos() {
        wx.request({
            url: this.data.base + "/upload-service/video/findUserLikeAndUpload/" + this.data.userId,
            method: 'GET',
            success: (res) => {
                if (res.data.status) {
                    this.setData({
                        uploadList: res.data.data.uploadList,
                        likeList: res.data.data.likeList,
                    })
                }
            }
        })
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

    onClickLeft() {
        wx.navigateBack({delta: 1})
    },

});