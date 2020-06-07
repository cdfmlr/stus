// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package model

type Teacher struct {
	Tid   string `gorm:"PRIMARY_KEY"` // 教师id
	Tname string // 姓名
	Tdept string // 所在系
	Tsex  bool   // 性别
	Tpro  string // 职称：讲师、教授这些，professional title
}
