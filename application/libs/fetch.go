package libs

import (
	"fmt"
	"log"
	"encoding/json"
	"github.com/simonvetter/modbus"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTControlRead struct {
	Topic string	`json:"Topic"`
	Address int		`json:"Address"`
	Length int		`json:"Length"`
	Value []uint16	`json:"Value"`
}

func Read_FuelCell(fc FuelCell, input_mqtt mqtt.Client, input_id int32) {
	jdata, err := json.Marshal(fc)
	if err != nil {
		//log.Printf("failed to read [fuelcell]: %v", err)
		return
	}
	//log.Println(string(jdata))
	topic := fmt.Sprintf("ems/fuelcell/%d", input_id)
	token := input_mqtt.Publish(topic, 0, false, string(jdata))
	token.Wait()
	//log.Println("[MODBUS TCP] fetch", topic)
}

func Read_battery(input_modbus *modbus.ModbusClient, input_mqtt mqtt.Client, input_id int32, input_count int32) {
	var bcs []BatteryCell
	for i:=0; i<=int(input_count/10); i++ {
		cell_begin := int32(i*10)
		if int(input_count/10) == i {
			cell_end := int32(i * 10 + int(input_count) % 10)
			Battery_Data(&bcs, input_modbus, cell_begin, cell_end, input_id)
		} else {
			cell_end := int32(i * 10 + 10)
			Battery_Data(&bcs, input_modbus, cell_begin, cell_end, input_id)
		}
	}
	jdata, err := json.Marshal(bcs)
	if err != nil {
		//log.Printf("failed to read [battery]: %v", err)
		return
	}
	topic := fmt.Sprintf("ems/battery/%d", input_id)
	token := input_mqtt.Publish(topic, 0, false, string(jdata))
	token.Wait()
	//log.Println("[MODBUS TCP] fetch", topic)
}
func Read_meter(input_modbus *modbus.ModbusClient, input_mqtt mqtt.Client, input_id int32, input_chanel int32, input_start int32) {
	udata, err := input_modbus.ReadRegisters(uint16(input_start), 74, modbus.HOLDING_REGISTER)
	if err != nil {
		//log.Printf("failed to read [meter]: %v", err)
		return
	}
	ddata := MTR_Cell_Decode(udata, input_id, input_chanel)	
	jdata, err := json.Marshal(ddata)
	if err != nil {
		//log.Printf("failed to read [meter]: %v", err)
		return
	}
	topic := fmt.Sprintf("ems/meter/%d", input_id)
	token := input_mqtt.Publish(topic, 0, false, string(jdata))
	token.Wait()
	//log.Println("[MODBUS TCP] fetch", topic)
}
func Read_inverter(input_modbus *modbus.ModbusClient, input_mqtt mqtt.Client, input_id int32, input_chanel int32, input_start int32) {
	udata, err := input_modbus.ReadRegisters(uint16(input_start), 43, modbus.HOLDING_REGISTER)	//第一組
	if err != nil {
		//log.Printf("failed to read [inverter]: %v", err)
		return
	}
	ddata := InverterDecode(udata, input_id, input_chanel)
	jdata, err := json.Marshal(ddata)
	if err != nil {
		//log.Printf("failed to read [inverter]: %v", err)
		return
	}
	topic := fmt.Sprintf("ems/inverter/%d/%d", input_id, input_chanel)
	token := input_mqtt.Publish(topic, 0, false, string(jdata))
	token.Wait()
	//log.Println("[MODBUS TCP] fetch", topic)
}
func Read_main(input_modbus *modbus.ModbusClient, input_mqtt mqtt.Client, input_id int32, input_chanel int32) {
	udata, err := input_modbus.ReadRegisters(1, 24, modbus.HOLDING_REGISTER)
	if err != nil {
		//log.Printf("failed to read [main]: %v", err)
		return
	}
	ddata := MainDecode(udata, input_id, input_chanel)
	jdata, err := json.Marshal(ddata)
	if err != nil {
		//log.Printf("failed to read [main]: %v", err)
		return
	}
	topic := fmt.Sprintf("ems/main/%d/%d", input_id, input_chanel)
	token := input_mqtt.Publish(topic, 0, false, string(jdata))
	token.Wait()
	//log.Println("[MODBUS TCP] fetch", topic)
}
func Read_reg(input_modbus *modbus.ModbusClient, input_adr int, input_len int) []uint16 {
	udata, err := input_modbus.ReadRegisters(uint16(input_adr), uint16(input_len), modbus.HOLDING_REGISTER)
	if err != nil {
		log.Printf("failed to read [%d] len(%d): %v\n", input_adr, input_len, err)
		input_modbus.Open();
		return nil
	}
	//log.Printf("[MODBUS TCP] Read [%d] len(%d) -> %v", input_adr, input_len, udata)
	return udata
}
func Write_reg(input_modbus *modbus.ModbusClient, input_adr int, input_val int) {
	err := input_modbus.WriteRegister(uint16(input_adr), uint16(input_val))
	if err != nil {
		log.Printf("failed to write [%d] -> [%d]: %v\n", input_val, input_adr, err)
		input_modbus.Open();
		return
	}
}
