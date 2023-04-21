package entity

type User struct {
	UserId    string `json:"user_id" gorm:"column:userId;primary_key"`
	Pwd       string `json:"pwd" gorm:"column:pwd"`
	Nickname  string `json:"nickname" gorm:"column:nickname;default:默认用户名"`
	Signature string `json:"signature" gorm:"column:signature;default:该用户很懒，什么都没有留下..."`
}

func (u User) TableName() string {
	return "user_tb"
}
