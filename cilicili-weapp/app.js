App({

    globalData: {
        userInfo: {
            id: '',
            nickname: '未授权',
            address: '',
            gender: '',
            avatarUrl: '',
            hasInfo: '',
            fansCount: 0,
            followCount: 0,
            likeCount: 0,
        },
        BottomTabBarHeight: 0,
        windowHeight: 0,
        windowWidth: 0,
        baseUrl: 'http://172.20.10.5:8762',
    },

    onLaunch() {

        wx.login({
            success: (res) => {
                console.log("wx login success")
                wx.request({
                    url: this.globalData.baseUrl + '/account-service/login?code=' + res.code,
                    method: "POST",
                    success: (res) => {
                        if (res.data.code === 200) {
                            this.globalData.userInfo = res.data.data
                            console.log('get account success: ', res.data.data)
                        }
                    }
                })
            }
        })

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
