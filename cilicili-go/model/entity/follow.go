package entity

type Follow struct {
	FromUser User
	ToUser   User
	Id       string
	FromId   string //谁关注的 user_id
	ToId     string //关注的谁 user_id
}
