package models

import (
	orm "goin/db"
)

type User struct {
	ID                int     `json:"id" gorm:"column:id;primary_key"`                    //用户ID
	UserType          int     `json:"user_type"`                                          //用户类型
	Sex               int     `json:"sex"`                                                //性别
	Birthday          int     `json:"birthday"`                                           //生日
	LastLoginTime     int     `json:"last_login_time"`                                    //最后登录时间
	Score             int     `json:"score"`                                              //用户积分
	Coin              int     `json:"coin"`                                               //金币
	Balance           float32 `json:"balance"`                                            //余额
	CreateTime        int64   `json:"create_time" gorm:"column:create_time;default:null"` //故事分类
	UserStatus        int     `json:"user_status"`                                        //用户状态;0:禁用,1:正常,2:未验证'
	UserLogin         string  `json:"user_login"`                                         //用户名
	UserPass          string  `json:"user_pass"`                                          //密码
	UserNickname      string  `json:"user_nickname"`                                      //昵称
	UserEmail         string  `json:"user_email"`                                         //用户登录邮箱
	UserUrl           string  `json:"user_url"`                                           //用户个人网址
	Avatar            string  `json:"avatar"`                                             //用户头像
	Signature         string  `json:"signature"`                                          //个性签名
	LastLoginIp       string  `json:"last_login_ip"`                                      //最后登录ip
	UserActivationKey string  `json:"user_activation_key"`                                //激活码
	Mobile            string  `json:"mobile"`                                             //手机号
	More              string  `json:"more"`
}

func (user User) Error() string {
	panic("implement me")
}

func (User) TableName() string {
	return "hello"
}

func (user *User) Info() (info User, err error) {
	if err = orm.Eloquent.Where(&user).First(&info).Error; err != nil {
		return
	}
	return
}
