package main

import (
	"fmt"
	"log"
	"time"
	"sync"
	"strings"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"libs"
)

//mysql
var mysql_client *sql.DB
var mysql_host = "ncu-ems-db.csv.tw"
var mysql_port = 3306
var mysql_user = "ncu_ems"
var mysql_pass = "O*/*YmlBe25ono5("
var mysql_dbse = "ncu_ems"

//mqtt
var mqtt_client mqtt.Client
var mqtt_host = "tcp://ncu-ems-mqtt.csv.tw:1883"
var mqtt_user = "worker"
var mqtt_pass = "V5u9URTy"

//cache
type DBCache struct {
	Topic string
	UpdateTime int64
}
var dcs []DBCache

func main() {
	//mqtt
    mqtt_opts := mqtt.NewClientOptions()
    mqtt_opts.AddBroker(mqtt_host)
    //mqtt_opts.SetClientID("go_mqtt_client_v2")
    mqtt_opts.SetUsername(mqtt_user)
    mqtt_opts.SetPassword(mqtt_pass)
	mqtt_opts.SetPingTimeout(10 * time.Second)
	mqtt_opts.SetKeepAlive(10 * time.Second)
	mqtt_opts.SetAutoReconnect(true)
	mqtt_opts.SetMaxReconnectInterval(10 * time.Second)
    mqtt_opts.SetDefaultPublishHandler(MessageHandler)
    mqtt_opts.OnConnect = ConnectHandler
    mqtt_opts.OnConnectionLost = libs.ConnectLostHandler
    mqtt_client = mqtt.NewClient(mqtt_opts)
	defer mqtt_client.Disconnect(250)
	
	//connect
	if token := mqtt_client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	
	if true {
		//mysql
		var err error
		mysql_conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysql_user, mysql_pass, mysql_host, mysql_port, mysql_dbse)
		mysql_client, err = sql.Open("mysql", mysql_conn)
		if err != nil {
			log.Println("connection to mysql failed:", err)
		}
		defer mysql_client.Close()
	}
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

func UpDCS(DCS []DBCache, Topic string) ([]DBCache, bool) {
	now := time.Now()
	//每60秒更新一次資料庫
	tmp_dc := DBCache { Topic: Topic, UpdateTime: now.Unix() + 60 }
	for x := range DCS {
		if DCS[x].Topic == Topic {
			if DCS[x].UpdateTime <= now.Unix() {
				DCS[x].UpdateTime = now.Unix() + 60
				return DCS, true		//有更新
			} else {
				return DCS, false		//無更新
			}
		}
	}
	return append(DCS, tmp_dc), true	//有更新
}

var MessageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02d %02d:%02d:00",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute())
	upflag := false
	dcs, upflag = UpDCS(dcs, msg.Topic())
	if(upflag == true) {
		log.Printf("[UPLOAD] %s\n", msg.Topic())
		if(true && strings.Contains(msg.Topic(), "ems/meter")) {
			var meter_cell libs.MTRCell
			json.Unmarshal(msg.Payload(), &meter_cell)
			go libs.Insert_MTR_db(mysql_client, formatted, meter_cell)
		}
		if(true && strings.Contains(msg.Topic(), "ems/battery")) {
			var battery_cells []libs.BatteryCell
			json.Unmarshal(msg.Payload(), &battery_cells)
			go libs.Insert_Battery_db(mysql_client, formatted, battery_cells)
		}
		if(true && strings.Contains(msg.Topic(), "ems/inverter")) {
			var inverter_cells libs.InverterCell
			json.Unmarshal(msg.Payload(), &inverter_cells)
			go libs.Insert_Inverter_db(mysql_client, formatted, inverter_cells)
		}
		if(true && strings.Contains(msg.Topic(), "ems/fuelcell")) {
			var fuelcell_cells libs.FuelCell
			json.Unmarshal(msg.Payload(), &fuelcell_cells)
			go libs.Insert_FuelCell_db(mysql_client, formatted, fuelcell_cells)
		}
		if(true && strings.Contains(msg.Topic(), "ems/main")) {
			var main_cell libs.MainCell
			json.Unmarshal(msg.Payload(), &main_cell)
			go libs.Insert_Main_db(mysql_client, formatted, main_cell)
		}
	} else {
		//log.Printf("[SKIP] %s\n", msg.Topic())
	}
}

var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	//subscribe
	token := client.Subscribe("ems/#", 1, nil)
	token.Wait()
	for token.Error() != nil {
		log.Println("[MQTT] subscribe failed.")
		time.Sleep(1 * time.Second)
		token.Wait()
	}
	log.Println("[MQTT] subscribe topic ems/#.")
}