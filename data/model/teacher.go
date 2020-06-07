// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package model

type Teacher struct {
	Tid   string `gorm:"PRIMARY_KEY";json:"tid"` // 教师id
	Tname string `json:"tname"`                  // 姓名
	Tdept string `json:"tdept"`                  // 所在系
	Tsex  bool   `json:"tsex"`                   // 性别
	Tpro  string `json:"tpro"`                   // 职称：讲师、教授这些，professional title
}
