// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package model

type Course struct {
	Cid    string  `gorm:"PRIMARY_KEY"` // 课程id
	Cname  string  // 课程名称
	Ctype  string  // 课程类型：比如"必修"，"公选"什么的
	Cpoint float32 // 学分
	Cweek  string  // 开课周次
	Ctime  string  // 上课时间
	Caddr  string  // 上课地点（教室）
}
