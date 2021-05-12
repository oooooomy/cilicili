package entity

type User struct {
	Id          string `json:"id"`
	Nickname    string `json:"nickname"`
	Address     string `json:"address"`
	Gender      string `json:"gender"`
	AvatarUrl   string `json:"avatarUrl"`
	HasInfo     bool   `json:"hasInfo"`
	FansCount   int    `json:"fansCount"`
	FollowCount int    `json:"followCount"`
}

type Admin struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Follow struct {
	Id         string `json:"id"`
	FromUserId string `json:"fromUserId" gorm:"form_user_id"` //谁关注的 user_id
	ToUserId   string `json:"toUserId" gorm:"to_user_id"`     //关注的谁 user_id
}
