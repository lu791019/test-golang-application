package main

import (
	"fmt"
	"log"
	"time"
	"regexp"
	"strings"
	"strconv"
	"encoding/json"
	"github.com/simonvetter/modbus"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"libs"
)

//mqtt
var mqtt_client mqtt.Client
var mqtt_host = "tcp://YOUR_MQTT_HOST_IP"
var mqtt_user = "YOUR_MQTT_USER_NAME"
var mqtt_pass = "YOUR_MQTT_PASSWORD"

//modbus tcp
var modbus_client *modbus.ModbusClient
var modbus_host = "tcp://YOUR_HOST_IP"

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
    mqtt_opts.SetDefaultPublishHandler(MessagePubHandler)
    mqtt_opts.OnConnect = ConnectHandler
    mqtt_opts.OnConnectionLost = libs.ConnectLostHandler
    mqtt_client = mqtt.NewClient(mqtt_opts)
	defer mqtt_client.Disconnect(1500)
	//connect
	if token := mqtt_client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//modbus TCP
	var err error
	modbus_client, err = modbus.NewClient(&modbus.ClientConfiguration{
        URL: modbus_host,
        Timeout: 60 * time.Second,
    })
	if err != nil {
        print("connect failed.")
    }
	for modbus_client.Open() != nil {
		log.Println("[MODBUS TCP] failed to connect")
		time.Sleep(1 * time.Second)
	}
	defer modbus_client.Close()
	for {
//		for modbus_client.Open() != nil {
//			//modbus_client.Connect()
//			time.Sleep(1 * time.Second)
//			log.Println("[MODBUS TCP] failed to connect")
//			continue
//		}
		if true {	//Main
			go libs.Read_main(modbus_client, mqtt_client, 1, 0)				//ID=1, Chanel=0
		}
		if true {	//Inverter
			go libs.Read_inverter(modbus_client, mqtt_client, 1, 0, 118)	//ID=1, Chanel=0
			go libs.Read_inverter(modbus_client, mqtt_client, 1, 1, 180)	//ID=1, Chanel=1
			go libs.Read_inverter(modbus_client, mqtt_client, 1, 2, 242)	//ID=1, Chanel=2
		}
		if true {	//Meter
			go libs.Read_meter(modbus_client, mqtt_client, 1, 0, 2148)		//ID=1, Chanel=0
		}
		if true {	//Battery
			go libs.Read_battery(modbus_client, mqtt_client, 1, 18)			//ID=1
		}
		time.Sleep(1 * time.Second)
	}
}

var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("[MQTT] Connected")
	log.Println("[MQTT] subscribe topic (ems/control/1/modbus).")
	client.Subscribe("ems/control/1/modbus", 1, nil)
}

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	//log.Printf("[MQTT] Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	if(strings.Contains(msg.Topic(), "ems/control")) {
		msg_str := fmt.Sprintf("%s", msg.Payload())
		match, _ := regexp.MatchString(`([a-z]+)\:(\d+)\:(\d+)`, msg_str)
		if match {
			res := strings.Split(msg_str, ":")
			typ := fmt.Sprintf("%s", res[0])
			adr, _ := strconv.Atoi(res[1])
			val, _ := strconv.Atoi(res[2])
			//log.Println(res)
			//log.Println(typ, adr, val)
			if(typ == "w") {
				//log.Println("[MODBUS TCP] Write [%d] -> [%d]", val, adr)
				libs.Write_reg(modbus_client, adr, val)
			}
			if(typ == "r") {
				udata := libs.Read_reg(modbus_client, adr, val)
				//log.Printf("[MODBUS TCP] Read adr[%d] len(%d) -> %v\n", adr, val, udata)
				var pkg libs.MQTTControlRead
				pkg.Topic = msg.Topic() + "/callback"
				pkg.Address = adr
				pkg.Length = val
				if udata != nil {
					pkg.Value = udata
					jdata, _ := json.Marshal(pkg)
					client.Publish(pkg.Topic, 0, false, string(jdata))
				} else {
					jdata, _ := json.Marshal(pkg)
					client.Publish(pkg.Topic, 0, false, string(jdata))
				}
			}
		}
	}
}