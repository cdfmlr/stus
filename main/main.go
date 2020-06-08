// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"fmt"
	"net/http"
	"stus/data"
	"stus/service"
)

func main() {
	database := data.NewDatabase("mysql", "c:000123@/stus?charset=utf8&parseTime=True&loc=Local")
	staticDir := "/Users/c/Projects/stus/"

	db, err := database.Open()
	if err != nil {
		panic(err)
	}
	server := service.NewService(db, staticDir)

	err = http.ListenAndServe(":9001", server)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listen and serve at http://0.0.0.0:9001")
}
