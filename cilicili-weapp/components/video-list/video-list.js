Component({
    /**
     * 组件的属性列表
     */
    properties: {
        videoList: {
            type: Array
        }
    },

    /**
     * 组件的初始数据
     */
    data: {
        isStop: false,
        showComment: false,
    },

    // 在组件实例进入页面节点树时执行
    attached: function () {
        this.videoContext = wx.createVideoContext('0', this)
        this.videoContext.play()
    },

    // 在组件实例被从页面节点树移除时执行
    detached: function () {
        console.log("组件实例被从页面节点树移除")
    },

    /**
     * 组件的方法列表
     */
    methods: {
        swiperChange(e) {
            if (e.detail.source === 'touch') {
                this.videoContext.pause()
                this.videoContext = wx.createVideoContext(e.detail.current.toString(), this)
                this.videoContext.play()
                this.setData({isStop: false})
            }
        },

        openUserDetails(e) {
            wx.navigateTo({
                url: '/pages/user/user?id=' + e.currentTarget.dataset.id
            })
        },

        openComment() {
            this.setData({
                showComment: true
            })
        },

        onCloseComment() {
            this.setData({
                showComment: false
            })
        },

    }

})
