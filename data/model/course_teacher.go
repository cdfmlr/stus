// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package model

type CourseTeacherRelation struct {
	Cid string `gorm:"primary_key" json:"cid"` // 课程
	Tid string `gorm:"primary_key" json:"tid"` // 教师

	Course  Course  `gorm:"foreignkey:Cid; association_foreignkey:Cid"`
	Teacher Teacher `gorm:"foreignkey:Tid; association_foreignkey:Tid"`
}
