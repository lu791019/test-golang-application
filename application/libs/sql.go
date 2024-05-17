package libs

import (
	"log"
	"fmt"
	"reflect"
	"database/sql"
)

func Generate_Insert_SQL(x interface{}, table_name string, timestamp string) string {
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
	//log.Println(sql)
	return sql
}

func Not_in_db(sql_conn *sql.DB, table_name string, id int32, chanel int32, timestamp string) bool {
	sql := "SELECT `id` FROM `%s_data` WHERE `%s_id` = %d AND `chanel` = %d AND `timestamp` LIKE '%s'"
	sql = fmt.Sprintf(sql, table_name, table_name, id, chanel, timestamp)
	var cnt int
    _ = sql_conn.QueryRow(sql).Scan(&cnt)
	if cnt > 0 {
		return false
	}
	return true
}

func Insert_Battery_db(sql_conn *sql.DB, timestamp string, battery_cells []BatteryCell) bool {
	for _,cell := range battery_cells {
		if Not_in_db(sql_conn, "battery", cell.ID, cell.Chanel, timestamp) == false {
			continue
		}
		sql := Generate_Insert_SQL(cell, "battery_data", timestamp)
		var err error
		_, err = sql_conn.Exec(sql)
		if err != nil {
			log.Println(err)
		}
    }
	return true
}

func Insert_Main_db(sql_conn *sql.DB, timestamp string, cell MainCell) bool {
	if Not_in_db(sql_conn, "mainsys", cell.ID, cell.Chanel, timestamp) == false {
		return false
	}
	sql := Generate_Insert_SQL(cell, "mainsys_data", timestamp)
	var err error
	_, err = sql_conn.Exec(sql)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Insert_MTR_db(sql_conn *sql.DB, timestamp string, cell MTRCell) bool {
	if Not_in_db(sql_conn, "meter", cell.ID, cell.Chanel, timestamp) == false {
		return false
	}
	sql := Generate_Insert_SQL(cell, "meter_data", timestamp)
	var err error
	_, err = sql_conn.Exec(sql)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Insert_Inverter_db(sql_conn *sql.DB, timestamp string, cell InverterCell) bool {
	if Not_in_db(sql_conn, "inverter", cell.ID, cell.Chanel, timestamp) == false {
		return false
	}
	sql := Generate_Insert_SQL(cell, "inverter_data", timestamp)
	var err error
	_, err = sql_conn.Exec(sql)
	if err != nil {
		log.Println(err)
	}
	return true
}

func Insert_FuelCell_db(sql_conn *sql.DB, timestamp string, cell FuelCell) bool {
	if Not_in_db(sql_conn, "fuelcell", cell.ID, cell.Chanel, timestamp) == false {
		return false
	}
	sql := Generate_Insert_SQL(cell, "fuelcell_data", timestamp)
//	log.Println(cell)
//	log.Println(sql)
	var err error
	_, err = sql_conn.Exec(sql)
	if err != nil {
		log.Println(err)
	}
	return true
}




