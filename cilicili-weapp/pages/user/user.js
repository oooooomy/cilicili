const App = getApp()

Page({

    data: {
        base: App.globalData.baseUrl,
        likeList: [],
        uploadList: [],
        userId: '',
        active: 0,
        isFollow: false,
        user: {},
    },

    onLoad: function (options) {
        wx.showLoading({
          title: 'Loading',
        })
        this.setData({
            userId: options.id
        })
        console.log(options.id)
        this.isFollow()
        wx.request({
            url: this.data.base + "/account-service/user/" + options.id,
            method: 'GET',
            success: (res) => {
                const r = res.data
                if (r.code === 200) {
                    this.setData({user: r.data})
                }
            }
        })
        this.getMyVideos()
        setTimeout(()=>{
            wx.hideLoading({})
        }, 400)
    },

    //查看是否关注了
    isFollow(){
        wx.request({
          url: this.data.base + '/account-service/follow/from/' + App.globalData.userInfo.id+'/to/' + this.data.userId,
          method: 'GET',
          success: (res)=>{
              const r = res.data
              this.setData({
                  isFollow: r.data
              })
          }
        })
    },

    onClickDeleteFollow(){
        const url = this.data.base + '/account-service/follow/from/' + App.globalData.userInfo.id+'/to/' + this.data.userId
        let _this = this
        wx.showModal({
            title: '用户提示',
            content: '确定要取消关注吗?',
            success (res) {
              if (res.confirm) {
                wx.request({
                    url: url,
                    method: 'DELETE',
                    success: (res)=>{
                        const r = res.data
                        if(r.code === 200){
                            _this.setData({
                                isFollow: false,
                            })
                            wx.showToast({
                                title: '取消关注成功',
                                icon: 'success',
                                duration: 2000
                            })                              
                        }
                    }
                  })
              }
            }
        })
    },

    onClickFollow(){
        wx.vibrateShort({type: 'heavy'})
        let form = {
            fromUserId: App.globalData.userInfo.id,
            toUserId: this.data.userId,
        }
        wx.request({
            url: this.data.base + '/account-service/follow/create',
            method: 'POST',
            data: form,
            success: (res) => {
                if (res.data.code === 200) {
                    this.setData({
                        isFollow: true,
                    })
                    wx.showToast({
                        title: 'Thank you',
                        icon: 'success',
                        duration: 2000
                    })
                }else {
                    wx.showToast({
                        title: res.data.msg,
                        icon: 'error',
                        duration: 2000
                    })
                }
            }
        })
    },

    getMyVideos() {
        wx.request({
            url: this.data.base + "/upload-service/video/findUserLikeAndUpload/" + this.data.userId,
            method: 'GET',
            success: (res) => {
                if (res.data.code === 200) {
                    this.setData({
                        uploadList: res.data.data.uploadList ? res.data.data.uploadList : '',
                        likeList: res.data.data.likeList ? res.data.data.likeList : '',
                    })
                }
            }
        })
    },

    onLike() {
        wx.navigateTo({
            url: '/pages/videos/videos?videos=' + JSON.stringify(this.data.likeList)
        })
    },

    onUpload() {
        wx.navigateTo({
            url: '/pages/videos/videos?videos=' + JSON.stringify(this.data.uploadList)
        })
    },

    onClickLeft() {
        wx.navigateBack({delta: 1})
    },

});