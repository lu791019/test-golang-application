package libs

import (
//	"log"
	"fmt"
//	"github.com/simonvetter/modbus"
)

type InverterCell struct {
	ID int32								`json:"ID" sql:"inverter_id" format:"%d"`
	Chanel int32							`json:"Chanel" sql:"chanel" format:"%d"`
	GridVoltageAC float32					`json:"GridVoltageAC" sql:"GridVoltageAC" format:"%f"`
	GridCurrentAC float32					`json:"GridCurrentAC" sql:"GridCurrentAC" format:"%f"`
	GridWattAC float32						`json:"GridWattAC" sql:"GridWattAC" format:"%f"`
	GridFrequency float32					`json:"GridFrequency" sql:"GridFrequency" format:"%f"`
	ProtectLoadVoltage float32				`json:"ProtectLoadVoltage" sql:"ProtectLoadVoltage" format:"%f"`
	ProtectLoadCurrent float32				`json:"ProtectLoadCurrent" sql:"ProtectLoadCurrent" format:"%f"`
	ProtectLoadWatt float32					`json:"ProtectLoadWatt" sql:"ProtectLoadWatt" format:"%f"`
	ProtectLoadFrequency float32			`json:"ProtectLoadFrequency" sql:"ProtectLoadFrequency" format:"%f"`
	InverterLoadVoltage float32				`json:"InverterLoadVoltage" sql:"InverterLoadVoltage" format:"%f"`
	InverterLoadCurrent float32				`json:"InverterLoadCurrent" sql:"InverterLoadCurrent" format:"%f"`
	InverterWatt float32					`json:"InverterWatt" sql:"InverterWatt" format:"%f"`
	InverterFrequency float32				`json:"InverterFrequency" sql:"InverterFrequency" format:"%f"`
	InverterInternalTemperature float32		`json:"InverterInternalTemperature" sql:"InverterInternalTemperature" format:"%f"`
	WorkingMode string						`json:"WorkingMode" sql:"WorkingMode" format:"'%s'"`
	BatteryVoltage float32					`json:"BatteryVoltage" sql:"BatteryVoltage" format:"%f"`
	BatteryTemperature float32				`json:"BatteryTemperature" sql:"BatteryTemperature" format:"%f"`
	BatteryCharingCurrent float32			`json:"BatteryCharingCurrent" sql:"BatteryCharingCurrent" format:"%f"`
	BatteryDischaringCurrent float32		`json:"BatteryDischaringCurrent" sql:"BatteryDischaringCurrent" format:"%f"`
	BatteryStatus string					`json:"BatteryStatus" sql:"BatteryStatus" format:"'%s'"`
	InputVoltageDC1st float32				`json:"InputVoltageDC1st" sql:"InputVoltageDC1st" format:"%f"`
	InputCurrentDC1st float32				`json:"InputCurrentDC1st" sql:"InputCurrentDC1st" format:"%f"`
	InputPower1st float32					`json:"InputPower1st" sql:"InputPower1st" format:"%f"`
	InputVoltageDC2nd float32				`json:"InputVoltageDC2nd" sql:"InputVoltageDC2nd" format:"%f"`
	InputCurrentDC2nd float32				`json:"InputCurrentDC2nd" sql:"InputCurrentDC2nd" format:"%f"`
	InputPower2nd float32					`json:"InputPower2nd" sql:"InputPower2nd" format:"%f"`
	ErrorCode1Low string					`json:"ErrorCode1Low" sql:"ErrorCode1Low" format:"'%s'"`
	ErrorCode1High string					`json:"ErrorCode1High" sql:"ErrorCode1High" format:"'%s'"`
	ErrorCode2Low string					`json:"ErrorCode2Low" sql:"ErrorCode2Low" format:"'%s'"`
	ErrorCode2High string					`json:"ErrorCode2High" sql:"ErrorCode2High" format:"'%s'"`
	LoadLoadTotal float32					`json:"LoadLoadTotal" sql:"LoadLoadTotal" format:"%f"`
	GridPowerTotal float32					`json:"GridPowerTotal" sql:"GridPowerTotal" format:"%f"`
	PVPowerTotal float32					`json:"PVPowerTotal" sql:"PVPowerTotal" format:"%f"`
}

func InverterDecode(input_val []uint16, input_id int32, input_chanel int32) InverterCell {
	var tmp InverterCell
	tmp.ID = input_id
	tmp.Chanel = input_chanel
	tmp.GridVoltageAC = int16_to_int(input_val[0], 0.1)
	tmp.GridCurrentAC = int16_to_int(input_val[1], 0.01)
	tmp.GridWattAC = int32_to_int(input_val[3], input_val[2], 0.1)
	tmp.GridFrequency = int16_to_int(input_val[4], 0.01)
	
	tmp.ProtectLoadVoltage = int16_to_int(input_val[6], 0.1)
	tmp.ProtectLoadCurrent = int16_to_int(input_val[7], 0.01)
	tmp.ProtectLoadWatt = int32_to_int(input_val[9], input_val[8], 0.1)
	tmp.ProtectLoadFrequency = int16_to_int(input_val[10], 0.01)
	
	
	tmp.InverterLoadVoltage = int16_to_int(input_val[12], 0.1)
	tmp.InverterLoadCurrent = int16_to_int(input_val[13], 0.01)
	tmp.InverterWatt = int32_to_int(input_val[15], input_val[14], 0.1)
	tmp.InverterFrequency = int16_to_int(input_val[16], 0.01)
	tmp.InverterInternalTemperature = int16_to_int(input_val[18], 0.1)
	
	tmp.WorkingMode = fmt.Sprintf("0x%02x", input_val[19])
	
	tmp.BatteryVoltage = int16_to_int(input_val[20], 0.01)
	tmp.BatteryTemperature = int16_to_int(input_val[21], 0.1)
	tmp.BatteryCharingCurrent = int16_to_int(input_val[22], 0.01) 
	tmp.BatteryDischaringCurrent = int16_to_int(input_val[23], 0.01)
	tmp.BatteryStatus = fmt.Sprintf("0x%02x", input_val[24])
	
	tmp.InputVoltageDC1st = int16_to_int(input_val[26], 0.1)
	tmp.InputCurrentDC1st = int16_to_int(input_val[27], 0.01)
	tmp.InputPower1st = int32_to_int(input_val[29], input_val[28], 0.1)

	tmp.InputVoltageDC2nd = int16_to_int(input_val[30], 0.1)
	tmp.InputCurrentDC2nd = int16_to_int(input_val[31], 0.01)
	tmp.InputPower2nd = int32_to_int(input_val[33], input_val[32], 0.1)
	
	tmp.ErrorCode1Low = fmt.Sprintf("0x%02x", input_val[33])
	tmp.ErrorCode1High = fmt.Sprintf("0x%02x", input_val[34])
	
	tmp.ErrorCode2Low = fmt.Sprintf("0x%02x", input_val[35])
	tmp.ErrorCode2High = fmt.Sprintf("0x%02x", input_val[36])
	
	tmp.LoadLoadTotal = int32_to_int(input_val[38], input_val[37], 1)
	tmp.GridPowerTotal = int32_to_int(input_val[40], input_val[39], 1)
	tmp.PVPowerTotal = int32_to_int(input_val[42], input_val[41], 1)
	
	return tmp
}
















