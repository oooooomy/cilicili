Page({

    data: {
        list: [{}, {}, {}, {}, {}],
        userId: '',
        active: 0,
    },

    onLoad: function (options) {
        this.setData({
            userId: options.id
        })
    },

    onChange(event) {
        wx.showToast({
            title: `切换到标签 ${event.detail.name}`,
            icon: 'none',
        });
    },

    onClickLeft() {
        wx.navigateBack({delta: 1})
    },

});