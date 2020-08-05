package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Message struct {
	Id      int64
	Content string
	AddTime int64
}

type MessageUser struct {
	Id        int
	MessageId int64
	UserId    int
	AddTime   int64
	Status    int
}

func init() {
	orm.RegisterModel(new(Message), new(MessageUser))
}

func SendMessageDo(content string) (messageId int64, err error) {
	message := Message{
		Content: content,
		AddTime: time.Now().Unix(),
	}
	messageId, err = orm.NewOrm().Insert(&message)
	return
}

func SendMessageUser(
	messageId int64,
	userId int,
) (err error) {
	messageUser := MessageUser{
		MessageId: messageId,
		UserId:    userId,
		AddTime:   time.Now().Unix(),
		Status:    1,
	}

	_, err = orm.NewOrm().Insert(&messageUser)
	return
}
