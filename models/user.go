package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id       int
	Name     string
	Password string
	AddTime  int64
	Avatar   string
	Status   int
	Mobile   string
}

type UserInfo struct {
	Id       int `json:"id"`
	Name     string `json:"name"`
	AddTime  int64 `json:"addTime"`
	Avatar   string `json:"avatar"`
	Status   int  `json:"status"`
}

func init() {
	orm.RegisterModel(new(User))
}

func GetUserIsExist(mobile string) bool {
	newOrm := orm.NewOrm()
	user := User{Mobile: mobile}
	err := newOrm.Read(&user, "Mobile")
	if err != nil {
		return false;
	}
	return true
}

func UserSave(mobile, password string) error {
	newOrm := orm.NewOrm()
	user := User{
		Mobile:   mobile,
		Password: password,
		Name:     "",
		AddTime:  time.Now().Unix(),
		Status:   1,
		Avatar:   "",
	}
	_, err := newOrm.Insert(&user)

	return err
}

func GetUser(mobile, password string) (user User) {
	newOrm := orm.NewOrm()
	newOrm.QueryTable("user").
		Filter("mobile", mobile).
		Filter("password", password).
		One(&user)
	return
}

func GetUserInfo(userId int) (user User) {
	orm.NewOrm().
		QueryTable("user").
		Filter("id", userId).
		Limit(1).
		One(&user, "id", "name", "addTime", "avatar", "status")

	return
}
