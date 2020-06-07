// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package service

import (
	"net/http"
	"stus/data"
	"testing"
)

func TestService(t *testing.T) {
	var d = data.NewDatabase("mysql", "user:password@/stus?charset=utf8&parseTime=True&loc=Local")
	var db, _ = d.Open()

	s := NewService(db, "/Users/user/Projects/stus/")
	err := http.ListenAndServe(":9001", s)
	if err != nil {
		t.Error(err)
	}
}
