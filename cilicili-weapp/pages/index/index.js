const App = getApp()

Page({

    data: {
        Height: App.globalData.windowHeight - App.globalData.BottomTabBarHeight,
        loading: false,
        userInfo: App.globalData.userInfo,
        active: 0,
        videoList: [],
    },

    onLoad: function (options) {
        this.loadIndexData()
    },

    /**
     * 加载首页需要的所有信息
     */
    loadIndexData() {
        this.setData({
            loading: true
        })
        wx.request({
            url: App.globalData.baseUrl + '/upload-service/video/findByStatus?status=true',
            method: 'GET',
            success: (res)=>{
                this.setData({
                    videoList: res.data.data
                })
                console.log('首页视频列表: ' , this.data.videoList)
            },
        })
        setTimeout(() => {
            this.setData({
                loading: false
            })
        }, 1500)
    },

    onChange(event) {
        this.setData({active: event.detail});
    },

})
