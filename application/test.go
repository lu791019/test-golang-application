package main

import (
	"fmt"
	"log"
//	"time"
//	"sync"
//	"strings"
	"reflect"
//	"database/sql"
//	"encoding/json"
//	_ "github.com/go-sql-driver/mysql"
//	mqtt "github.com/eclipse/paho.mqtt.golang"
	"libs"
)

func main() {
	
	var bat libs.BatteryCell
	bat_sql := Generate_SQL_Insert(bat, "battery_data", "")
	log.Println(bat_sql)
}

func Generate_SQL_Insert(x interface{}, table_name string, timestamp string) string {
	values := reflect.ValueOf(x)
	typesOf := values.Type()
	sql_p1 := "`id`,"
	sql_p2 := "NULL,"
	for i := 0; i < values.NumField(); i++ {
		sql_p1 += " `" + typesOf.Field(i).Tag.Get("sql") + "`,"
		sql_p2 += " " + typesOf.Field(i).Tag.Get("format") + ","
		sql_p2 = fmt.Sprintf(sql_p2, values.Field(i).Interface())
	}
	sql_p1 += " `timestamp`"
	sql_p2 += " %s"
	if timestamp == "" {
		sql_p2 = fmt.Sprintf(sql_p2, "NOW()")
	} else {
		sql_p2 = fmt.Sprintf(sql_p2, "'" + timestamp + "'")
	}
	sql := "INSERT INTO `%s` (%s) VALUES (%s);"
	sql = fmt.Sprintf(sql, table_name, sql_p1, sql_p2)
	return sql
}