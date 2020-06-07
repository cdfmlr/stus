// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package model

type Course struct {
	Cid    string  `gorm:"PRIMARY_KEY";json:"cid"` // 课程id
	Cname  string  `json:"cname"`                  // 课程名称
	Ctype  string  `json:"ctype"`                  // 课程类型：比如"必修"，"公选"什么的
	Cpoint float32 `json:"cpoint"`                 // 学分
	Cweek  string  `json:"cweek"`                  // 开课周次
	Ctime  string  `json:"ctime"`                  // 上课时间
	Caddr  string  `json:"caddr"`                  // 上课地点（教室）
}
