// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package teacher

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"stus/data/model"
	"stus/util/response"
)

type TeacherApi struct {
	DB *gorm.DB
}

func NewTeacherApi(DB *gorm.DB) *TeacherApi {
	return &TeacherApi{DB: DB}
}

func (t TeacherApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/courses":
		t.courses(w, r)
	case "/data_of_course":
		t.dataOfCourse(w, r)
	}
}

func (t TeacherApi) courses(w http.ResponseWriter, r *http.Request) {
	tid := r.FormValue("tid")
	if tid == "" {
		response.ResponseJson(&w, map[string]string{"error": "unexpected empty tid"})
		return
	}

	var courseTeachers []model.CourseTeacherRelation
	t.DB.Where("tid = ?", tid).Preload("Course").Find(&courseTeachers)

	result := make([]CourseResp, 0)
	for _, ct := range courseTeachers {
		studentNum := 0
		t.DB.Model(&model.StudentCourseRelation{}).
			Where("cid = ?", ct.Cid).Count(&studentNum)
		result = append(result, CourseResp{
			Course:     ct.Course,
			StudentNum: studentNum,
		})
	}

	response.ResponseJson(&w, result)
}

type CourseResp struct {
	model.Course
	StudentNum int `json:"student_num"`
}

func (t TeacherApi) dataOfCourse(w http.ResponseWriter, r *http.Request) {
	cid := r.FormValue("cid")
	if cid == "" {
		response.ResponseJson(&w, map[string]string{"error": "unexpected empty cid"})
		return
	}

	var studentCourses []model.StudentCourseRelation
	t.DB.Where("cid = ?", cid).Preload("Student").Find(&studentCourses)

	resp := DataOfCourseResp{}

	t.DB.Raw("SELECT COUNT(result) AS count, AVG(result) AS average, MAX(result) AS best, MIN(result) AS worst, COUNT(CASE WHEN result<60 THEN 1 END) AS not_pass_count FROM student_course_relations WHERE cid = ?", cid).Scan(&resp)

	var students []StudentWithResult
	for _, sc := range studentCourses {
		students = append(students, StudentWithResult{
			Student: sc.Student,
			Result:  sc.Result,
		})
	}
	resp.Students = students

	response.ResponseJson(&w, resp)
}

type DataOfCourseResp struct {
	Students     []StudentWithResult `json:"students" gorm:"column:students"`
	Count        int                 `json:"count" gorm:"column:count"`
	Average      float32             `json:"average" gorm:"column:average"`
	Best         float32             `json:"best" gorm:"column:best"`
	Worst        float32             `json:"worst" gorm:"column:worst"`
	NotPassCount float32             `json:"not_pass_count" gorm:"column:not_pass_count"`
}

type StudentWithResult struct {
	model.Student
	Result float32 `json:"result"`
}
