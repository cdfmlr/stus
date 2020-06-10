// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package model

type StudentCourseRelation struct {
	Sid    string  `gorm:"PRIMARY_KEY" json:"sid"` // 学生id
	Cid    string  `gorm:"PRIMARY_KEY" json:"cid"` // 课程id
	Result float32 `json:"result"`                 // 成绩

	Student Student `gorm:"foreignkey:Sid; association_foreignkey:Sid"`
	Course  Course  `gorm:"foreignkey:Cid; association_foreignkey:Cid"`
}
