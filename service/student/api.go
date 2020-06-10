// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package student

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"stus/data/model"
	"stus/util/response"
)

type StudentApi struct {
	DB *gorm.DB
}

func NewStudentApi(DB *gorm.DB) *StudentApi {
	return &StudentApi{DB: DB}
}

func (s StudentApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/exam_result":
		s.examResult(w, r)
	}
}

func (s StudentApi) examResult(w http.ResponseWriter, r *http.Request) {
	sid := r.FormValue("sid")
	if sid == "" {
		response.ResponseJson(&w, map[string]string{"error": "sid missing"})
		return
	}

	var studentCourses []model.StudentCourseRelation
	s.DB.Where("sid = ?", sid).Preload("Course").Find(&studentCourses)

	result := make([]ExamResult, 0)
	for _, sc := range studentCourses {
		result = append(result, ExamResult{
			Cid:    sc.Cid,
			Cname:  sc.Course.Cname,
			Result: sc.Result,
		})
	}

	response.ResponseJson(&w, result)
}

type ExamResult struct {
	Cid    string  `json:"cid"`
	Cname  string  `json:"cname"`
	Result float32 `json:"result"`
}
