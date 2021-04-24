App({

    globalData: {
        userInfo: null,
        BottomTabBarHeight: 0,
        windowHeight: 0,
        windowWidth: 0,
        baseUrl: 'http://172.20.10.5:8762',
    },

    onLaunch() {

        //获取手机信息
        wx.getSystemInfo({
            success: (res) => {
                const model = res.model
                const x = model.search('iPhone X') === 0 || model.search('iPhone 1') === 0
                this.globalData.BottomTabBarHeight = x ? 90 : 56
                this.globalData.windowWidth = res.windowWidth
                this.globalData.windowHeight = res.windowHeight
            }
        })

    },

})
