package entity

type User struct {
	Id          string
	Username    string
	Password    string
	FollowCount int64
	LikeCount   int64
	AvatarUrl   string
	CreateAt    int64
}
