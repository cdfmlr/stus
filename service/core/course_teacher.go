// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package core

import (
	"github.com/jinzhu/gorm"
	"stus/data/model"
)

type SerCoreCourseTeacherRelation struct {
	DB *gorm.DB
}

func (s SerCoreCourseTeacherRelation) Create(courseTeacherRelation *model.CourseTeacherRelation) {
	s.DB.Create(courseTeacherRelation)
}

func (s SerCoreCourseTeacherRelation) Read(query string) []model.CourseTeacherRelation {
	courseTeacherRelations := make([]model.CourseTeacherRelation, 0)
	s.DB.Where(query).Find(&courseTeacherRelations)
	return courseTeacherRelations
}

func (s SerCoreCourseTeacherRelation) Update(courseTeacherRelations []model.CourseTeacherRelation) {
	for _, courseTeacherRelation := range courseTeacherRelations {
		s.DB.Save(&courseTeacherRelation)
	}
}

func (s SerCoreCourseTeacherRelation) Delete(query string) {
	s.DB.Where(query).Delete(model.CourseTeacherRelation{})
}
