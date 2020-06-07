// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package core

import (
	"github.com/jinzhu/gorm"
	"stus/data/model"
)

type SerCoreTeacher struct {
	DB *gorm.DB
}

func (s SerCoreTeacher) Create(teacher *model.Teacher) {
	s.DB.Create(teacher)
}

func (s SerCoreTeacher) Read(query string) []model.Teacher {
	teachers := make([]model.Teacher, 0)
	s.DB.Where(query).Find(&teachers)
	return teachers
}

func (s SerCoreTeacher) Update(teachers []model.Teacher) {
	for _, teacher := range teachers {
		s.DB.Save(&teacher)
	}
}

func (s SerCoreTeacher) Delete(query string) {
	s.DB.Where(query).Delete(model.Teacher{})
}
