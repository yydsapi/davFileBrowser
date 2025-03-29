// Copyright 2019 yydsapi Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package configs

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// _ "yyds/statik"
var Db *sql.DB // global variable to share it between main and the HTTP handler
var err error
var MySqlConnStr string

func Initdb() {
	Db, err = sql.Open("mysql", MySqlConnStr)
	if err != nil {
		panic(err)
	}
	// 设置数据库闲链接超时时间?[mysql] packets.go:122: closing bad idle connection: EOF
	Db.SetConnMaxLifetime(time.Second * 270)
	Db.SetMaxIdleConns(10000)
	var version string
	Db.QueryRow("SELECT VERSION()").Scan(&version)
	//fmt.println("---Connected to Mysql ok, version: " + version)
}
