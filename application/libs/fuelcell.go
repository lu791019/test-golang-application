package libs

import (
//	"log"
	"github.com/gosnmp/gosnmp"
)

type FuelCell struct {
	ID int32								`json:"ID" sql:"fuelcell_id" format:"%d"`
	Chanel int32							`json:"Chanel" sql:"chanel" format:"%d"`
	SystemState int32						`json:"SystemState" sql:"SystemState" format:"%d" snmp:".1.3.36466.1.3.1.1.1.1.0"`
	DoorOpen int32							`json:"DoorOpen" sql:"DoorOpen" format:"%d" snmp:".1.3.36466.1.3.1.1.1.4.0"`
	SystemFault int32						`json:"SystemFault" sql:"SystemFault" format:"%d" snmp:".1.3.36466.1.3.1.1.1.2.0"`
	SystemWarning int32						`json:"SystemWarning" sql:"SystemWarning" format:"%d" snmp:".1.3.36466.1.3.1.1.1.3.0"`
	Voltage float32							`json:"Voltage" sql:"Voltage" format:"%f" snmp:".1.3.36466.1.1.1.2.1.10.0"`
	Current float32							`json:"Current" sql:"Current" format:"%f" snmp:".1.3.36466.1.1.1.2.1.11.0"`
	GrossPower float32						`json:"GrossPower" sql:"GrossPower" format:"%f" snmp:".1.3.36466.1.1.1.2.1.12.0"`
	KilowattHours float32					`json:"KilowattHours" sql:"KilowattHours" format:"%f" snmp:".1.3.36466.1.1.1.2.1.9.0"`
	FuelLevel float32						`json:"FuelLevel" sql:"FuelLevel" format:"%f" snmp:".1.3.36466.1.3.1.2.1.4.0"`
	RunTime float32							`json:"RunTime" sql:"RunTime" format:"%f" snmp:".1.3.36466.1.1.1.2.1.8.0"`
	FilterMaint float32						`json:"FilterMaint" sql:"FilterMaint" format:"%f" snmp:".1.3.36466.1.3.1.4.1.7.0"`

	LowFuelLevelSetpoint float32			`json:"LowFuelLevelSetpoint" sql:"LowFuelLevelSetpoint" format:"%f" snmp:".1.3.36466.1.3.1.5.1.4.0"`
	StartVoltageFirstStack float32			`json:"StartVoltageFirstStack" sql:"StartVoltageFirstStack" format:"%f" snmp:".1.3.36466.1.3.1.5.1.5.0"`
	StartVoltageSecondStack float32			`json:"StartVoltageSecondStack" sql:"StartVoltageSecondStack" format:"%f" snmp:".1.3.36466.1.3.1.5.1.6.0"`
	FloatVoltage float32					`json:"FloatVoltage" sql:"FloatVoltage" format:"%f" snmp:".1.3.36466.1.3.1.5.1.7.0"`
	
	Stack1Fault float32						`json:"Stack1Fault" sql:"Stack1Fault" format:"%f" snmp:".1.3.36466.1.3.1.1.2.3.0"`
	Stack2Fault float32						`json:"Stack2Fault" sql:"Stack2Fault" format:"%f" snmp:".1.3.36466.1.3.1.1.2.4.0"`
}

func FuelCellDecode(input_val []gosnmp.SnmpPDU, input_id int32, input_chanel int32) FuelCell {
	var tmp FuelCell
	tmp.ID = input_id
	tmp.Chanel = input_chanel
	for _, variable := range input_val {
		//log.Println(variable.Type, variable.Name, variable.Value)
		if variable.Name == ".1.3.36466.1.3.1.1.1.1.0" { tmp.SystemState = int32(gosnmp.ToBigInt(variable.Value).Int64()) }
		if variable.Name == ".1.3.36466.1.3.1.1.1.4.0" { tmp.DoorOpen = int32(gosnmp.ToBigInt(variable.Value).Int64()) }
		if variable.Name == ".1.3.36466.1.3.1.1.1.2.0" { tmp.SystemFault = int32(gosnmp.ToBigInt(variable.Value).Int64()) }
		if variable.Name == ".1.3.36466.1.3.1.1.1.3.0" { tmp.SystemWarning = int32(gosnmp.ToBigInt(variable.Value).Int64()) }
		if variable.Name == ".1.3.36466.1.1.1.2.1.10.0" { 
			tmp_a := gosnmp.ToBigInt(variable.Value).Int64()
			tmp_b := float32(tmp_a)
			tmp.Voltage = float32(tmp_b/float32(100)) 
		}
		if variable.Name == ".1.3.36466.1.1.1.2.1.11.0" { 
			tmp_a := gosnmp.ToBigInt(variable.Value).Int64()
			tmp_b := float32(tmp_a)
			tmp.Current = float32(tmp_b/float32(100)) 
		}
		if variable.Name == ".1.3.36466.1.1.1.2.1.12.0" { 
			tmp_a := gosnmp.ToBigInt(variable.Value).Int64()
			tmp_b := float32(tmp_a)
			tmp.GrossPower = float32(tmp_b/float32(100)) 
		}
		if variable.Name == ".1.3.36466.1.1.1.2.1.9.0" {
			tmp_a := gosnmp.ToBigInt(variable.Value).Int64()
			tmp_b := float32(tmp_a)
			tmp.KilowattHours = float32(tmp_b/float32(100)) 
		}
		if variable.Name == ".1.3.36466.1.3.1.2.1.4.0" {
			tmp_a := gosnmp.ToBigInt(variable.Value).Int64()
			tmp_b := float32(tmp_a)
			tmp.FuelLevel = float32(tmp_b/float32(100))
		}
		if variable.Name == ".1.3.36466.1.1.1.2.1.8.0" {
			tmp_a := gosnmp.ToBigInt(variable.Value).Int64()
			tmp_b := float32(tmp_a)
			tmp.RunTime = float32(tmp_b/float32(100))
		}
		if variable.Name == ".1.3.36466.1.3.1.4.1.7.0" {
			tmp_a := gosnmp.ToBigInt(variable.Value).Int64()
			tmp_b := float32(tmp_a)
			tmp.FilterMaint = float32(tmp_b/float32(100))
		}
		if variable.Name == ".1.3.36466.1.3.1.5.1.4.0" {
			tmp_a := gosnmp.ToBigInt(variable.Value).Int64()
			tmp_b := float32(tmp_a)
			tmp.LowFuelLevelSetpoint = float32(tmp_b/float32(100))
		}
		if variable.Name == ".1.3.36466.1.3.1.5.1.5.0" {
			tmp_a := gosnmp.ToBigInt(variable.Value).Int64()
			tmp_b := float32(tmp_a)
			tmp.StartVoltageFirstStack = float32(tmp_b/float32(100))
		}
		if variable.Name == ".1.3.36466.1.3.1.5.1.6.0" {
			tmp_a := gosnmp.ToBigInt(variable.Value).Int64()
			tmp_b := float32(tmp_a)
			tmp.StartVoltageSecondStack = float32(tmp_b/float32(100))
		}
		if variable.Name == ".1.3.36466.1.3.1.5.1.7.0" {
			tmp_a := gosnmp.ToBigInt(variable.Value).Int64()
			tmp_b := float32(tmp_a)
			tmp.FloatVoltage = float32(tmp_b/float32(100))
		}
		if variable.Name == ".1.3.36466.1.3.1.1.2.3.0" { tmp.Stack1Fault = float32(gosnmp.ToBigInt(variable.Value).Int64()) }
		if variable.Name == ".1.3.36466.1.3.1.1.2.4.0" { tmp.Stack2Fault = float32(gosnmp.ToBigInt(variable.Value).Int64()) }
	}
	return tmp
}