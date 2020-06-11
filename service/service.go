// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
	"stus/data/model"
	"stus/service/core"
	"stus/service/student"
	"stus/service/teacher"
	"stus/util/logging"
	"stus/util/response"
)

type Service struct {
	DB        *gorm.DB
	StaticDir string

	tokenHolder *AccessTokenHolder

	coreApiServer    *core.CoreApi
	studentApiServer *student.StudentApi
	teacherApiServer *teacher.TeacherApi
	fileServer       http.Handler
}

func NewService(DB *gorm.DB, staticDir string) *Service {
	s := &Service{DB: DB, StaticDir: staticDir}

	s.tokenHolder = NewAccessTokenHolder(DB, 86400*15)

	s.fileServer = http.FileServer(http.Dir(s.StaticDir))
	s.coreApiServer = core.NewCoreApi(DB)
	s.studentApiServer = student.NewStudentApi(DB)
	s.teacherApiServer = teacher.NewTeacherApi(DB)
	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logging.Info("HTTP Serve: ", r.Method, r.URL.Path)

	// CORS
	if r.Method == "OPTIONS" {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "token, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Write([]byte{})
		return
	}

	if strings.HasPrefix(r.URL.Path, "/api") {
		if strings.HasPrefix(r.URL.Path, "/api/login") {
			s.login(w, r)
			return
		}

		// 除了 login 的操作都要检查 token
		user, err := s.checkAccess(r)
		if err != nil {
			response.ResponseJson(&w, map[string]string{"error": err.Error()})
			return
		} else {
			logging.Info("Request from user:", *user)
		}
		//logging.Debug(fmt.Sprintf("%#v", r))

		switch {
		case strings.HasPrefix(r.URL.Path, "/api/core"):
			http.StripPrefix("/api/core", s.coreApiServer).ServeHTTP(w, r)
			return
		case strings.HasPrefix(r.URL.Path, "/api/student"):
			http.StripPrefix("/api/student", s.studentApiServer).ServeHTTP(w, r)
			return
		case strings.HasPrefix(r.URL.Path, "/api/teacher"):
			http.StripPrefix("/api/teacher", s.teacherApiServer).ServeHTTP(w, r)
			return
		default:
			http.NotFound(w, r)
			return
		}
	} else { // !/api
		s.fileServer.ServeHTTP(w, r)
	}
}

func (s *Service) checkAccess(r *http.Request) (*model.Passwd, error) {
	token := r.Header.Get("token")
	if token == "" {
		return &model.Passwd{}, fmt.Errorf("missing token")
	}
	user, ok := s.tokenHolder.Get(token)
	if !ok {
		return &model.Passwd{}, fmt.Errorf("not login")
	}

	return user, nil
}
