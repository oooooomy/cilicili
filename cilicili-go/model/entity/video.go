package entity

type Video struct {
	User          User
	Music         Music
	Id            string
	UserId        string
	Title         string
	Url           string
	OriginalSound bool //是否为原声 否则播放mp3
	MusicId       string
	Description   string
	Location      string //定位
	LikeCount     int64  //点赞数
	CommentCount  int64  //评论数
	ShareCount    int64  //分享数
	CreateAt      int64
}

type Music struct {
	Id        string
	Name      string //名字
	Author    string //作者
	count     int64  //使用该音乐人数
	PosterUrl string //音乐海报封面图片地址
	Url       string //mp3地址
	CreateAt  int64
}

type VideoLike struct {
	id      string
	videoId string
	userId  string
}
