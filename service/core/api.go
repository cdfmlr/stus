// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package core

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"stus/data/model"
	"stus/util/response"
)

type CoreApi struct {
	DB *gorm.DB
}

func NewCoreApi(db *gorm.DB) *CoreApi {
	return &CoreApi{DB: db}
}

func (a CoreApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/student":
		a.ServeStudent(w, r)
	case "/course":
		a.ServeCourse(w, r)
	case "/teacher":
		a.ServeTeacher(w, r)
	case "/ct":
		a.ServeCourseTeacherRelation(w, r)
	case "/sc":
		a.ServeStudentCourseRelation(w, r)
	}
}

func (a CoreApi) ServeStudent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		st := &model.Student{}
		if err := json.Unmarshal([]byte(r.FormValue("students")), st); err != nil {
			response.ResponseJson(&w, map[string]string{"error": err.Error()})
			return
		}
		if st.Sid != "" {
			SerCoreStudent{a.DB}.Create(st)
		}
		responseSuccess(w)
		return
	case "GET":
		query := r.FormValue("query")
		if query == "" {
			response.ResponseJson(&w, map[string]string{"error": "unexpected empty query"})
		}
		response.ResponseJson(&w, SerCoreStudent{a.DB}.Read(query))
		return
	case "PUT":
		var st []model.Student
		if err := json.Unmarshal([]byte(r.FormValue("students")), &st); err != nil {
			response.ResponseJson(&w, map[string]string{"error": err.Error()})
			return
		}
		SerCoreStudent{a.DB}.Update(st)
		responseSuccess(w)
		return
	case "DELETE":
		query := r.FormValue("query")
		if query == "" {
			response.ResponseJson(&w, map[string]string{"error": "unexpected empty query"})
		}
		SerCoreStudent{a.DB}.Delete(query)
		responseSuccess(w)
		return
	}
}

func (a CoreApi) ServeCourse(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		st := &model.Course{}
		if err := json.Unmarshal([]byte(r.FormValue("record")), st); err != nil {
			response.ResponseJson(&w, map[string]string{"error": err.Error()})
			return
		}
		if st.Cid != "" {
			SerCoreCourse{a.DB}.Create(st)
		}
		responseSuccess(w)
		return
	case "GET":
		query := r.FormValue("query")
		if query == "" {
			response.ResponseJson(&w, map[string]string{"error": "unexpected empty query"})
		}
		response.ResponseJson(&w, SerCoreCourse{a.DB}.Read(query))
		return
	case "PUT":
		var st []model.Course
		if err := json.Unmarshal([]byte(r.FormValue("record")), &st); err != nil {
			response.ResponseJson(&w, map[string]string{"error": err.Error()})
			return
		}
		SerCoreCourse{a.DB}.Update(st)
		responseSuccess(w)
		return
	case "DELETE":
		query := r.FormValue("query")
		if query == "" {
			response.ResponseJson(&w, map[string]string{"error": "unexpected empty query"})
		}
		SerCoreCourse{a.DB}.Delete(query)
		responseSuccess(w)
		return
	}
}

func (a CoreApi) ServeTeacher(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		st := &model.Teacher{}
		if err := json.Unmarshal([]byte(r.FormValue("record")), st); err != nil {
			response.ResponseJson(&w, map[string]string{"error": err.Error()})
			return
		}
		if st.Tid != "" {
			SerCoreTeacher{a.DB}.Create(st)
		}
		responseSuccess(w)
		return
	case "GET":
		query := r.FormValue("query")
		if query == "" {
			response.ResponseJson(&w, map[string]string{"error": "unexpected empty query"})
		}
		response.ResponseJson(&w, SerCoreTeacher{a.DB}.Read(query))
		return
	case "PUT":
		var st []model.Teacher
		if err := json.Unmarshal([]byte(r.FormValue("record")), &st); err != nil {
			response.ResponseJson(&w, map[string]string{"error": err.Error()})
			return
		}
		SerCoreTeacher{a.DB}.Update(st)
		responseSuccess(w)
		return
	case "DELETE":
		query := r.FormValue("query")
		if query == "" {
			response.ResponseJson(&w, map[string]string{"error": "unexpected empty query"})
		}
		SerCoreTeacher{a.DB}.Delete(query)
		responseSuccess(w)
		return
	}
}

func (a CoreApi) ServeCourseTeacherRelation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		st := &model.CourseTeacherRelation{}
		if err := json.Unmarshal([]byte(r.FormValue("record")), st); err != nil {
			response.ResponseJson(&w, map[string]string{"error": err.Error()})
			return
		}
		if st.Cid != "" && st.Tid != "" {
			SerCoreCourseTeacherRelation{a.DB}.Create(st)
		}
		responseSuccess(w)
		return
	case "GET":
		query := r.FormValue("query")
		if query == "" {
			response.ResponseJson(&w, map[string]string{"error": "unexpected empty query"})
		}
		response.ResponseJson(&w, SerCoreCourseTeacherRelation{a.DB}.Read(query))
		return
	case "PUT":
		var st []model.CourseTeacherRelation
		if err := json.Unmarshal([]byte(r.FormValue("record")), &st); err != nil {
			response.ResponseJson(&w, map[string]string{"error": err.Error()})
			return
		}
		SerCoreCourseTeacherRelation{a.DB}.Update(st)
		responseSuccess(w)
		return
	case "DELETE":
		query := r.FormValue("query")
		if query == "" {
			response.ResponseJson(&w, map[string]string{"error": "unexpected empty query"})
		}
		SerCoreCourseTeacherRelation{a.DB}.Delete(query)
		responseSuccess(w)
		return
	}
}

func (a CoreApi) ServeStudentCourseRelation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		st := &model.StudentCourseRelation{}
		if err := json.Unmarshal([]byte(r.FormValue("record")), st); err != nil {
			response.ResponseJson(&w, map[string]string{"error": err.Error()})
			return
		}
		if st.Sid != "" && st.Cid != "" {
			SerCoreStudentCourseRelation{a.DB}.Create(st)
		}
		responseSuccess(w)
		return
	case "GET":
		query := r.FormValue("query")
		if query == "" {
			response.ResponseJson(&w, map[string]string{"error": "unexpected empty query"})
		}
		response.ResponseJson(&w, SerCoreStudentCourseRelation{a.DB}.Read(query))
		return
	case "PUT":
		var st []model.StudentCourseRelation
		if err := json.Unmarshal([]byte(r.FormValue("record")), &st); err != nil {
			response.ResponseJson(&w, map[string]string{"error": err.Error()})
			return
		}
		SerCoreStudentCourseRelation{a.DB}.Update(st)
		responseSuccess(w)
		return
	case "DELETE":
		query := r.FormValue("query")
		if query == "" {
			response.ResponseJson(&w, map[string]string{"error": "unexpected empty query"})
		}
		SerCoreStudentCourseRelation{a.DB}.Delete(query)
		responseSuccess(w)
		return
	}
}

func responseSuccess(w http.ResponseWriter) {
	response.ResponseJson(&w, map[string]string{"success": "success"})
}
