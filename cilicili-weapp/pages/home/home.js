Component({

    properties: {
        videoList: {type: Array}
    },

    /**
     * 页面的初始数据
     */
    data: {
        
    },

    methods: {
        onClickLeft() {
            wx.navigateTo({
                url: '/pages/videos/videos?type=follow'
            })
        },
    },

})