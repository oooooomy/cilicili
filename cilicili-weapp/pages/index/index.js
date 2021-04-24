const App = getApp()

Page({

    data: {
        Height: App.globalData.windowHeight -  App.globalData.BottomTabBarHeight,
        loading: false,
        active: 0,
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
