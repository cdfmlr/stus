// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package service

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
	"stus/service/core"
	"stus/util/logging"
)

type Service struct {
	DB        *gorm.DB
	StaticDir string

	coreApiServer *core.CoreApi
	fileServer    http.Handler
}

func NewService(DB *gorm.DB, staticDir string) *Service {
	s := &Service{DB: DB, StaticDir: staticDir}
	s.fileServer = http.FileServer(http.Dir(s.StaticDir))
	s.coreApiServer = core.NewCoreApi(DB)
	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logging.Info("HTTP Serve: ", r.Method, r.URL.Path)

	switch {
	case strings.HasPrefix(r.URL.Path, "/api/core"):
		http.StripPrefix("/api/core", s.coreApiServer).ServeHTTP(w, r)
	default:
		s.fileServer.ServeHTTP(w, r)
	}
}
