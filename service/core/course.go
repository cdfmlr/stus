// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package core

import (
	"github.com/jinzhu/gorm"
	"stus/data/model"
)

type SerCoreCourse struct {
	DB *gorm.DB
}

func (s SerCoreCourse) Create(course *model.Course) {
	s.DB.Create(course)
}

func (s SerCoreCourse) Read(query string) []model.Course {
	courses := make([]model.Course, 0)
	s.DB.Where(query).Find(&courses)
	return courses
}

func (s SerCoreCourse) Update(courses []model.Course) {
	for _, course := range courses {
		s.DB.Save(&course)
	}
}

func (s SerCoreCourse) Delete(query string) {
	s.DB.Where(query).Delete(model.Course{})
}
