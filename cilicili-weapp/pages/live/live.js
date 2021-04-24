// pages/home/home.js
Page({

    /**
     * 页面的初始数据
     */
    data: {
        current: 0,
        active: 0,
        posterList: [{}, {}, {}, {}],
        roomList: [
            {
                id: '1',
                title: '伞兵一号卢本伟准备就绪',
                name: 'white 卢本伟',
                url: 'https://gaoyuanming-photo.oss-cn-beijing.aliyuncs.com/cilicili/dy1%20%281%29.jpeg'
            },
            {
                id: '',
                title: '国服第一刀妹上线 冲大师',
                name: '即将拥有人鱼线的PDD',
                url: 'https://gaoyuanming-photo.oss-cn-beijing.aliyuncs.com/cilicili/dy1%20%282%29.jpeg'
            },
            {
                id: '3',
                title: 'SKT职业选手faker',
                name: 'STK T1 faker',
                url: 'https://gaoyuanming-photo.oss-cn-beijing.aliyuncs.com/cilicili/dy1%20%283%29.jpeg'
            },
            {
                id: '4',
                title: '家人们阿很好的撒娇回家',
                name: '吐字笑笑',
                url: 'https://gaoyuanming-photo.oss-cn-beijing.aliyuncs.com/cilicili/dy1%20%285%29.jpeg'
            },
            {
                id: '5',
                title: '家人们阿很好的撒娇回家',
                name: 'LPL官方赛事',
                url: 'https://gaoyuanming-photo.oss-cn-beijing.aliyuncs.com/cilicili/dy1.jpeg'
            },
            {
                id: '6',
                title: '与质询双排中',
                name: '呆妹',
                url: 'https://gaoyuanming-photo.oss-cn-beijing.aliyuncs.com/cilicili/dy1%20%284%29.jpeg'
            },
        ]
    },

    clickRoom(e) {
        wx.navigateTo({
            url: '/pages/room/room'
        })
    },

    swiperChange(e) {
        this.setData({current: e.detail.current})
    },


    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {

    },

    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady: function () {

    },

})