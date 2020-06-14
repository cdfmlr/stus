// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package core

import (
	"github.com/jinzhu/gorm"
	"stus/data/model"
)

type SerCoreStudentCourseRelation struct {
	DB *gorm.DB
}

func (s SerCoreStudentCourseRelation) Create(studentCourseRelation *model.StudentCourseRelation) {
	s.DB.Create(studentCourseRelation)
}

func (s SerCoreStudentCourseRelation) Read(query string) []model.StudentCourseRelation {
	studentCourseRelations := make([]model.StudentCourseRelation, 0)
	s.DB.Where(query).Find(&studentCourseRelations)
	return studentCourseRelations
}

func (s SerCoreStudentCourseRelation) Update(studentCourseRelations []model.StudentCourseRelation) {
	for _, studentCourseRelation := range studentCourseRelations {
		s.DB.Save(&studentCourseRelation)
	}
}

func (s SerCoreStudentCourseRelation) Delete(query string) {
	s.DB.Where(query).Delete(model.StudentCourseRelation{})
}
