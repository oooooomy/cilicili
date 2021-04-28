const App = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        content: '',
        videoFrom: {},
        userInfo: App.globalData.userInfo,
    },

    onLoad: function (options) {
        this.setData({
            videoFrom: JSON.parse(options.videoForm)
        })
        console.log('页面初始表单：', this.data.videoFrom)
    },

    onClickLeft() {
        wx.navigateBack({delta: 1})
    },

    bindFormSubmit: function (e) {
        let form = this.data.videoFrom
        form.title = e.detail.value.textarea
        wx.showModal({
            title: '确认你的标题',
            content: form.title ? form.title : 'Null',
            success(res) {
                if (res.confirm) {
                    wx.showLoading({
                        title: "正在保存中"
                    })
                    wx.request({
                        url: App.globalData.baseUrl + '/upload-service/video',
                        method: 'POST',
                        data: form,
                        success: (res) => {
                            wx.hideLoading({})
                            wx.showToast({
                                title: '作品发布成功',
                                icon: 'success',
                                duration: 2000
                            })
                            setTimeout(() => {
                                wx.navigateBack({
                                    delta: 2,
                                })
                            }, 500)
                        },
                    })
                }
            }
        })

    }

})