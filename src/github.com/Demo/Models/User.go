package Models

type User struct{
	Id int64 `xorm:"pk autoincr" json:"id"`
	Account string `xorm:"varchar(11)" json:"account"`
	PassWord string `xorm:"varchar(11)" json:"passWord"`
}
