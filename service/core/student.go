// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package core

import (
	"github.com/jinzhu/gorm"
	"stus/data/model"
)

type SerCoreStudent struct {
	DB *gorm.DB
}

func (s SerCoreStudent) Create(student *model.Student) {
	s.DB.Create(student)
}

func (s SerCoreStudent) Read(query string) []model.Student {
	students := make([]model.Student, 0)
	s.DB.Where(query).Find(&students)
	return students
}

func (s SerCoreStudent) Update(students []model.Student) {
	for _, student := range students {
		if student.Sid != "" {
			s.DB.Save(&student)
		}
	}
}

func (s SerCoreStudent) Delete(query string) {
	s.DB.Where(query).Delete(model.Student{})
}
