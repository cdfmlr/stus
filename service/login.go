// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package service

import (
	"net/http"
	"stus/data/model"
	"stus/util/response"
)

func (s Service) login(w http.ResponseWriter, r *http.Request) {
	//r.Header.Get("token")
	utype := r.FormValue("utype")
	uid := r.FormValue("uid")
	password := r.FormValue("password")

	if utype == "" || uid == "" || password == "" {
		response.ResponseJson(&w, map[string]string{"error": "missing values"})
		return
	}

	user := &model.Passwd{}
	s.DB.Where(map[string]interface{}{"utype": utype, "uid": uid, "password": password}).First(user)
	if user.Uid != "" {
		token := s.tokenHolder.New(user)
		response.ResponseJson(&w, map[string]string{"token": token})
	} else {
		response.ResponseJson(&w, map[string]string{"error": "用户不存在"})
	}
}
