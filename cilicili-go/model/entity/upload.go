package entity

type Video struct {
	Id           string `json:"id"`
	UserId       string `json:"userId"`
	UserNickname string `json:"userNickname"`
	UserAvatar   string `json:"userAvatar"`
	MusicId      string `json:"musicId"`
	Title        string `json:"title"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
	ShareCount   int    `json:"shareCount"`
	Poster       string `json:"poster"`
	Url          string `json:"url"`
	UseMusic     bool   `json:"useMusic"`
	MusicName    string `json:"musicName"`
	MusicAuthor  string `json:"musicAuthor"`
	MusicPoster  string `json:"musicPoster"`
	MusicUrl     string `json:"musicUrl"`
	Status       bool   `json:"status"`
}

type Music struct {
	Id     string `json:"id"`
	Name   string `json:"name"`   //名字
	Author string `json:"author"` //作者
	Poster string `json:"poster"` //音乐海报封面图片地址
	Url    string `json:"url"`    //mp3地址
}

type VideoLike struct {
	Id      string `json:"id"`
	VideoId string `json:"videoId"`
	UserId  string `json:"userId"`
}

type Comment struct {
	Id           string `json:"id"`
	VideoId      string `json:"videoId"`
	UserId       string `json:"userId"`
	UserNickname string `json:"userNickname" gorm:"user_nickname"`
	UserAvatar   string `json:"userAvatar"`
	Content      string `json:"content"`
	CreateAt     string `json:"createAt"`
}
