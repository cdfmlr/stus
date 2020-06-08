// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package service

import (
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/gorm"
	"math/rand"
	"strings"
	"stus/data/model"
	"time"
)

type AccessToken struct {
	Token     string        // token值
	User      *model.Passwd // 对应用户
	LastTouch time.Time     // 最近访问时间
	expiresIn int64         // 过期时间，秒, accessToken.LastTouch.Unix() 与 time.Now.Unix() 的差值
}

// NewAccessToken 生成一个新的 AccessToken
func NewAccessToken(user *model.Passwd, expiresIn int64) *AccessToken {
	return &AccessToken{
		Token:     encodeToken(user),
		User:      user,
		LastTouch: time.Now(),
		expiresIn: expiresIn,
	}
}

// IsAvailable 返回 AccessToken 是否有效
// 一个「有效」的 AccessToken 满足：
//		1. 不是空的 AccessToken{}
//		2. LastTouch 距离现在时间不超过 expiresIn 秒
//		3. User 在数据库中存在
func (a *AccessToken) IsAvailable(DB *gorm.DB) bool {
	if a.Token == "" || a.expiresIn == 0 || a.User.Uid == "" {
		return false
	}

	userInDatabase := &model.Passwd{}
	DB.Where("utype = ? AND uid = ? AND password = ?",
		a.User.Utype, a.User.Uid, a.User.Password).
		First(userInDatabase)

	return a.Token != "" && // Token = "" 肯定是空的
		a.LastTouch.Unix()-time.Now().Unix() <= a.expiresIn && // LastTouch 距离现在时间不超过 expiresIn 秒
		userInDatabase.Uid != "" // User 在数据库中存在
}

// encodeToken 计算出一个 token
func encodeToken(user *model.Passwd) (token string) {
	// User 相关信息部分
	sl := []string{user.Utype, user.Uid, user.Password}
	data := []byte(strings.Join(sl, ""))
	tSite := fmt.Sprintf("%x", md5.Sum(data))

	// 随机部分
	rand.Seed(time.Now().UnixNano())
	sl = []string{fmt.Sprint(time.Now()), fmt.Sprint(rand.Float64())}
	data = []byte(strings.Join(sl, ""))
	tRand := fmt.Sprintf("%x", md5.Sum(data))

	return fmt.Sprintf("%s%s", tRand, tSite)
}
