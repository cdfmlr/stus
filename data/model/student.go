// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package model

type Student struct {
	Sid    string `gorm:"PRIMARY_KEY";json:"sid"` // 学生id（学号）
	Sname  string `json:"sname"`                  // 姓名
	Sdept  string `json:"sdept"`                  // 所在系
	Smajor string `json:"smajor"`                 // 专业
	Sage   int    `json:"sage"`                   // 年龄
	Ssex   bool   `json:"ssex"`                   // 性别
	Sgrade string `json:"sgrade"`                 // 年级
	Sclass string `json:"sclass"`                 // 班级
}
