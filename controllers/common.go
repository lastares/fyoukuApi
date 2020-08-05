package controllers

import "github.com/astaxie/beego"

type CommonController struct {
	beego.Controller
}

type Meta struct {
	Total int64 `json:"total"`
}

type ReturnSuccessData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

type ReturnSuccessDataNoMeta struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type Success struct {
	Code int `json:"code"`
}

type ReturnFailed struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 成功返回
func ResponseData(code int, data interface{}, meta interface{}) (json ReturnSuccessData) {
	json = ReturnSuccessData{
		Code: code,
		Data: data,
		Meta: meta,
	}
	return
}

func ResponseNoMeta(code int, data interface{}) (json ReturnSuccessDataNoMeta) {
	json = ReturnSuccessDataNoMeta{
		Code: code,
		Data: data,
	}
	return
}

func Ok(code int) Success {
	return Success{code}
}

func Failed(code int, msg string) ReturnFailed {
	return ReturnFailed{
		Code: code,
		Msg:  msg,
	}
}

func (this *CommonController) Response(response interface{}) {
	this.Data["json"] = response
	this.ServeJSON()
}
