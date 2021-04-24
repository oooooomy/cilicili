Page({

    data: {
        videoList: [{}, {}, {}, {}, {}]
    },

    onLoad: function (options) {

    },

    onClickLeft() {
        wx.navigateBack({delta: 1})
    },

});