// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package model

type Passwd struct {
	Utype    string `gorm:"PRIMARY_KEY" json:"utype"` // 用户类型，student 或 teacher 或 admin
	Uid      string `gorm:"PRIMARY_KEY" json:"uid"`   // 用户id，根据 Utype 的不同，取 Sid 或 Tid 或 Aid
	Password string `json:"password"`                 // 密码，储存摘要加密的值
}
