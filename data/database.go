// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"stus/data/model"
	"time"
)

type Database struct {
	dialect string
	source  string
}

// NewDatabase try to connect a database with given dialect and source, returns a Database.
// It do auto migrate and check foreign key restrict where connected successfully.
// `dialect` could be one of: "mssql", "mysql", "postgres" and "sqlite".
func NewDatabase(dialect string, source string) *Database {
	// 连接尝试
	db, err := gorm.Open(dialect, source)
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(
		&model.Student{},
		&model.Teacher{},
		&model.Course{},
		&model.CourseTeacherRelation{},
		&model.StudentCourseRelation{},
		&model.Passwd{},
	)

	// 外键约束
	db.Model(&model.CourseTeacherRelation{}).AddForeignKey("cid", "courses(cid)", "RESTRICT", "RESTRICT")
	db.Model(&model.CourseTeacherRelation{}).AddForeignKey("tid", "teachers(tid)", "RESTRICT", "RESTRICT")
	db.Model(&model.StudentCourseRelation{}).AddForeignKey("sid", "students(sid)", "RESTRICT", "RESTRICT")
	db.Model(&model.StudentCourseRelation{}).AddForeignKey("cid", "courses(cid)", "RESTRICT", "RESTRICT")

	return &Database{
		dialect: dialect,
		source:  source,
	}
}

// Open initialize a new db connection
func (d *Database) Open() (*gorm.DB, error) {
	db, err := gorm.Open(d.dialect, d.source)
	if err != nil {
		return &gorm.DB{}, err
	}

	// 连接池
	db.DB().SetMaxIdleConns(10)                 // 设置连接池中的最大闲置连接数
	db.DB().SetMaxOpenConns(100)                // 设置数据库的最大连接数量
	db.DB().SetConnMaxLifetime(5 * time.Minute) // 设置连接的最大可复用时间

	return db, nil
}
