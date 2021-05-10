const App = getApp()

Page({

    data: {
        base: App.globalData.baseUrl,
        videoList: [],
        barrageInputValue: '',
    },

    onLoad: function (options) {
        if (options.type === 'follow') {
            wx.request({
                url: this.data.base + '/upload-service/video/findAllByFollow?userId=' + App.globalData.userInfo.id,
                method: 'GET',
                success: (res) => {
                    this.setData({
                        videoList: res.data.data
                    })
                }
            })
        } else {
            this.setData({
                videoList: JSON.parse(options.videos)
            })
        }
    },

    onClickLeft() {
        wx.navigateBack({delta: 1})
    },

    barrageInput(e) {
        this.setData({
            barrageInputValue: e.detail.value
        })
    },

    clickSubmit(){
        
    },

});