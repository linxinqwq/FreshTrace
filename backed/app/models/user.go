package models

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Disabled bool   `json:"disabled"` // 是否禁用
	Identity string `json:"identity"` // 身份鉴权
	Question string `json:"question"` // 密保问题
	Answer   string `json:"answer"`   // 答案答案
}
