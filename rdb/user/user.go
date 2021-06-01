package user

import (
	"encoding/json"
	"goin/models"
	"goin/rdb"
)

/**
 * 用户缓存
 */

const InfoKey = "login:token:"

type Info struct {
	ID                int     `json:"id"`                  //用户ID
	UserType          int     `json:"user_type"`           //用户类型
	Sex               int     `json:"sex"`                 //性别
	Birthday          int     `json:"birthday"`            //生日
	LastLoginTime     int     `json:"last_login_time"`     //最后登录时间
	Score             int     `json:"score"`               //用户积分
	Coin              int     `json:"coin"`                //金币
	Balance           float32 `json:"balance"`             //余额
	CreateTime        int64   `json:"create_time"`         //故事分类
	UserStatus        int     `json:"user_status"`         //用户状态;0:禁用,1:正常,2:未验证'
	UserLogin         string  `json:"user_login"`          //用户名
	UserPass          string  `json:"user_pass"`           //密码
	UserNickname      string  `json:"user_nickname"`       //昵称
	UserEmail         string  `json:"user_email"`          //用户登录邮箱
	UserUrl           string  `json:"user_url"`            //用户个人网址
	Avatar            string  `json:"avatar"`              //用户头像
	Signature         string  `json:"signature"`           //个性签名
	LastLoginIp       string  `json:"last_login_ip"`       //最后登录ip
	UserActivationKey string  `json:"user_activation_key"` //激活码
	Mobile            string  `json:"mobile"`              //手机号
	More              string  `json:"more"`
}

func (rs Info) Set(token string) error {
	bytes, e := json.Marshal(rs)
	if e != nil {
		return e
	}
	rdb.Set(InfoKey+token, string(bytes), CacheTime)
	return nil
}

func (rs Info) Get(token string) (re models.User, err error) {
	s := rdb.Get(InfoKey + token)
	if len(s) <= 0 {
		return models.User{}, nil
	}
	err = json.Unmarshal([]byte(s), &re)
	return
}

func (rs Info) Del(token string) error {
	rdb.Del(InfoKey + token)
	return nil
}

func (rs Info) Do(token string) error {
	rdb.Do(InfoKey+token, CacheTime)
	return nil
}
