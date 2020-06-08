// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package service

import (
	"github.com/jinzhu/gorm"
	"stus/data/model"
	"sync"
	"time"
)

type AccessTokenHolder struct {
	mux       sync.Mutex
	accesses  map[string]AccessToken // token值 -> AccessToken对象 		// TODO: accesses 直接放在 map 里不太好，关机就没了，可以用 redis ，或者更直接用 mysql 来存
	DB        *gorm.DB
	expiresIn int64
}

func NewAccessTokenHolder(DB *gorm.DB, expiresIn int64) *AccessTokenHolder {
	return &AccessTokenHolder{
		accesses:  map[string]AccessToken{},
		DB:        DB,
		expiresIn: expiresIn,
	}
}

// Get 返回给定 token 对应的用户 user 以及成功标示 true，
// 如果 token 不存在于当前 holder，返回一个空 Passwd 以及失败标示 false
func (h *AccessTokenHolder) Get(token string) (user *model.Passwd, ok bool) {
	h.mux.Lock()
	defer h.mux.Unlock()

	a, ok := h.accesses[token]
	if ok && a.IsAvailable(h.DB) {
		a.LastTouch = time.Now()
		return a.User, true
	} else {
		delete(h.accesses, token)
		return a.User, false
	}
}

// New 为给定 user 分配一个新的 AccessToken，返回新生成的 token
func (h *AccessTokenHolder) New(user *model.Passwd) (token string) {
	h.mux.Lock()
	defer h.mux.Unlock()

	a := *NewAccessToken(user, h.expiresIn)
	h.accesses[a.Token] = a
	return a.Token
}

// CleanUnavailable 清理无效的 token 映射
func (h *AccessTokenHolder) CleanUnavailable() {
	h.mux.Lock()
	defer h.mux.Unlock()

	for k, v := range h.accesses {
		if !v.IsAvailable(h.DB) {
			delete(h.accesses, k)
		}
	}
}
