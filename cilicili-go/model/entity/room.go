package entity

type Room struct {
	Id        string `json:"id"`
	Avatar    string `json:"avatar"`
	Username  string `json:"username"`
	Token     string `json:"token"`
	Count     int    `json:"count"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	Email     string `json:"email"`
	PosterUrl string `json:"posterUrl" gorm:"poster_url"`
	Url       string `json:"url"`
	IsLiving  bool   `json:"isLiving" gorm:"is_living"`
}
