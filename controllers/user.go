package controllers

import (
	"fyoukuApi/models"
	"fyoukuApi/util"
	"regexp"
	"strconv"
	"strings"
)

type UserController struct {
	CommonController
}

//用户注册
// @router /register/save [post]
func (this *UserController) UserRegister() {
	mobile := this.GetString("mobile")
	password := this.GetString("password")

	if (mobile == "" || password == "") {
		this.Response(Failed(4001, "手机号或密码不能为空"))
	}

	regexResult, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)

	if !regexResult {
		this.Response(Failed(4002, "手机号格式不正确"))
	}

	// 判断用户是否已经注册
	isExist := models.GetUserIsExist(mobile)
	if isExist {
		this.Response(Failed(4005, "此手机号已经注册"))
	}

	err := models.UserSave(mobile, util.Md5(password))

	if err != nil {
		this.Response(Failed(4005, "用户注册失败"))
	}

	this.Data["json"] = Ok(0)
	this.ServeJSON()
}

// @router /user/login [post]
func (this *UserController) Login()  {
	mobile := this.GetString("mobile")
	password := this.GetString("password")

	if (mobile == "" || password == "") {
		this.Response(Failed(4001, "手机号或密码不能为空"))
	}

	regexResult, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)

	if !regexResult {
		this.Response(Failed(4002, "手机号格式不正确"))
	}

	user := models.GetUser(mobile, util.Md5(password))
	if user.Id == 0 {
		this.Response(Failed(-1, "手机号或密码不正确"))
	}
	this.Response(ResponseData(0, user, nil))
}

// 批量发送消息
// @router /send/message [post]
func (this *VideoController) SendMessage()  {
	uids := this.GetString("uids", "")
	content := this.GetString("content", "")
	if uids == "" || content == "" {
		this.Response(Failed(-1, "params error"))
	}

	messageId, err := models.SendMessageDo(content)
	if err != nil {
		this.Response(Failed(-1, "发送消息失败"))
	}

	userIds := strings.Split(uids, ",")
	for _, userId := range userIds {
		uid, _ := strconv.Atoi(userId)
		models.SendMessageUser(messageId, uid)
	}

	this.Response(Ok(0))

}