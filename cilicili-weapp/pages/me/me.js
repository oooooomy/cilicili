const App = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        list: [{}, {}, {}, {}, {}, {}, {}, {}, {}],
        active: 0,
        Height: App.globalData.windowHeight -  App.globalData.BottomTabBarHeight,
    },

    onChange(event) {
        wx.showToast({
            title: `切换到标签 ${event.detail.name}`,
            icon: 'none',
        });
    },

    onClickUpload(){
        wx.navigateTo({url: '/pages/release/release'})
    }

})