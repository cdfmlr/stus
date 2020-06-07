// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package data

import (
	"fmt"
	"stus/data/model"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	database := NewDatabase("mysql", "user:password@/stus?charset=utf8&parseTime=True&loc=Local")
	db, err := database.Open()
	if err != nil {
		t.Error(err)
	}
	t.Log(fmt.Sprintf("%#v", db))
	p := &model.Passwd{
		Utype:    "admin",
		Uid:      "0",
		Password: "123",
	}

	pp := &model.Passwd{}
	db.Where("uid = ?", "0").First(&pp)
	if pp.Password == "" {
		db.Create(p)
	} else {
		t.Log(fmt.Sprintf("exist: %#v", pp))
	}

	s := &model.Student{
		Sid:    "201810000999",
		Sname:  "张三",
		Sdept:  "数理",
		Smajor: "信息与计算科学",
		Sage:   20,
		Ssex:   true,
		Sgrade: "2018",
		Sclass: "1809",
	}
	db.Create(s)
	ss := &model.Student{}
	db.First(ss)
	t.Log(ss)

	c := &model.Course{
		Cid:    "w233333",
		Cname:  "白学",
		Ctype:  "公选",
		Cpoint: 1.5,
		Cweek:  "8",
		Ctime:  "105, 305",
		Caddr:  "教四十八楼九阶梯",
	}
	db.Create(c)
	cc := &model.Course{}
	db.First(cc)
	t.Log(cc)

	sc := &model.StudentCourseRelation{
		Sid:    "201810000999",
		Cid:    "w233333",
		Result: 0,
	}
	db.Create(sc)

}
