// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"stus/data"
	"stus/service"
)

func main() {
	var port int
	var staticDir string
	var dbDialect string
	var dbSource string

	flag.IntVar(&port, "port", 9001, "`port` for service")
	flag.StringVar(&staticDir, "static", "./static", "static (web ui) `dist` path")
	flag.StringVar(&dbDialect, "db_dialect", "mysql", "`dialect` of database. Expect one of mssql, mysql, postgres and sqlite")
	flag.StringVar(&dbSource, "db_source", "c:000123@/stus?charset=utf8&parseTime=True&loc=Local", "database `source`")

	flag.Parse()

	database := data.NewDatabase(dbDialect, dbSource)

	db, err := database.Open()
	if err != nil {
		panic(err)
	}
	server := service.NewService(db, staticDir)

	fmt.Println("Listen and serve at http://0.0.0.0:" + strconv.Itoa(port))
	err = http.ListenAndServe(":"+strconv.Itoa(port), server)
	if err != nil {
		panic(err)
	}
}
