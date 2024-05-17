package main

import (
	"log"
	"time"
	snmp "github.com/gosnmp/gosnmp"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"libs"
)

//mqtt
var mqtt_client mqtt.Client
var mqtt_host = "tcp://YOUR_MQTT_HOST_IP"
var mqtt_user = "YOUR_MQTT_USER_NAME"
var mqtt_pass = "YOUR_MQTT_PASSWORD"

//snmp
var snmp_conn_public = snmp.Default
var snmp_conn_partpub = snmp.Default
var snmp_host = "YOUR_HOST"

func main() {
	//mqtt
    mqtt_opts := mqtt.NewClientOptions()
    mqtt_opts.AddBroker(mqtt_host)
    mqtt_opts.SetUsername(mqtt_user)
    mqtt_opts.SetPassword(mqtt_pass)
	mqtt_opts.SetPingTimeout(10 * time.Second)
	mqtt_opts.SetKeepAlive(10 * time.Second)
	mqtt_opts.SetAutoReconnect(true)
	mqtt_opts.SetMaxReconnectInterval(10 * time.Second)
    mqtt_opts.SetDefaultPublishHandler(libs.MessagePubHandler)
    mqtt_opts.OnConnect = libs.ConnectHandler
    mqtt_opts.OnConnectionLost = libs.ConnectLostHandler
    mqtt_client = mqtt.NewClient(mqtt_opts)
	defer mqtt_client.Disconnect(250)
	//connect
	if token := mqtt_client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//snmp
	snmp_conn_public.Target = snmp_host
	snmp_conn_public.Community = "public"
	snmp_conn_partpub.Target = snmp_host
	snmp_conn_partpub.Community = "partpub"
	err_public := snmp_conn_public.Connect()
	if err_public != nil {
		log.Printf("Connect() err: %v\n", err_public)
	}
	err_partpub := snmp_conn_partpub.Connect()
	if err_partpub != nil {
		log.Printf("Connect() err: %v\n", err_partpub)
	}
	defer snmp_conn_public.Conn.Close()
	defer snmp_conn_partpub.Conn.Close()

	oids_public := []string{
		"YOUT_Oids_Public_List"
	}
	oids_partpub := []string{
		"YOUT_oids_partpub_List"
	}
	for {
		result_public, err2_public := snmp_conn_public.Get(oids_public)
		if err2_public != nil {
			log.Printf("Get() err: %v\n", err2_public)
			time.Sleep(10 * time.Second)
			continue
		}
		res_public := libs.FuelCellDecode(result_public.Variables, 2, 0)
		//log.Println(res_public)

		result_partpub, err2_partpub := snmp_conn_partpub.Get(oids_partpub)
		if err2_partpub != nil {
			log.Printf("Get() err: %v\n", err2_partpub)
			time.Sleep(10 * time.Second)
			continue
		}
		res_partpub := libs.FuelCellDecode(result_partpub.Variables, 2, 0)

		res_public.Stack1Fault = res_partpub.Stack1Fault
		res_public.Stack2Fault = res_partpub.Stack2Fault
		//log.Println(res_partpub)
		go libs.Read_FuelCell(res_public, mqtt_client, 2)
		time.Sleep(10 * time.Second)
	}
}