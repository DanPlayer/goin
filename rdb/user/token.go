package user

import (
	"errors"
	"goin/rdb"
	"strconv"
)

/**
用户TOKEN缓存
*/
const tokenKey = "login:users:"

type Token struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

func (rs Token) Set() error {
	if rs.ID <= 0 || len(rs.Token) <= 0 {
		return errors.New("token缓存的必要属性没有设置")
	}
	rdb.Set(tokenKey+strconv.Itoa(rs.ID), rs.Token, CacheTime)
	return nil
}
