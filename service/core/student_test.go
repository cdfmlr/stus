// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package core

import (
	"fmt"
	"stus/data"
	"stus/data/model"
	"testing"
)

var d = data.NewDatabase("mysql", "user:password@/stus?charset=utf8&parseTime=True&loc=Local")
var db, _ = d.Open()

func TestSerCoreStudent(t *testing.T) {
	s := SerCoreStudent{DB: db}
	student := &model.Student{
		Sid:    "201810000998",
		Sname:  "李四",
		Sdept:  "数理",
		Smajor: "信息与计算科学",
		Sage:   21,
		Ssex:   true,
		Sgrade: "2018",
		Sclass: "1809",
	}
	s.Create(student)

	ss := s.Read("sclass=1809")
	t.Log(fmt.Sprintf("%#v", ss))

	ss[0].Ssex = false
	t.Log(ss[0])

	s.Update([]model.Student{ss[0]})
	t.Log(fmt.Sprintf("%#v", s.Read("sclass=1809")))

	s.Delete("sid=201810000998")
	t.Log(fmt.Sprintf("%#v", s.Read("sclass=1809")))

}
