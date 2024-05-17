package libs

import (
	"log"
	"strings"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	//log.Printf("[MQTT] Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	if(strings.Contains(msg.Topic(), "ems/meter")) {
		log.Printf("%s\n", msg.Topic())
	}
	if(strings.Contains(msg.Topic(), "ems/battery")) {
		log.Printf("%s\n", msg.Topic())
	}
	if(strings.Contains(msg.Topic(), "ems/inverter")) {
		log.Printf("%s\n", msg.Topic())
	}
	if(strings.Contains(msg.Topic(), "ems/main")) {
		log.Printf("%s\n", msg.Topic())
	}
}
var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("[MQTT] Connected")
}
var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("[MQTT] Connect lost: %v", err)
}