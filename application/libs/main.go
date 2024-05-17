package libs

import (
//	"log"
//	"fmt"
//	"github.com/simonvetter/modbus"
)

type MainCell struct {
	ID int32								`json:"ID" sql:"mainsys_id" format:"%d"`
	Chanel int32							`json:"Chanel" sql:"chanel" format:"%d"`
	SystemPower float32						`json:"SystemPower" sql:"system_power" format:"%f"`
	GridPower float32						`json:"GridPower" sql:"grid_power" format:"%f"`
	ChargingPower float32					`json:"ChargingPower" sql:"charging_power" format:"%f"`
	DischargingPower float32				`json:"DischargingPower" sql:"discharging_power" format:"%f"`
	MPPTPower float32						`json:"MPPTPower" sql:"mppt_power" format:"%f"`
	TotalGreenEnergy float32				`json:"TotalGreenEnergy" sql:"total_green_energy" format:"%f"`
	CarbonRedcing float32					`json:"CarbonRedcing" sql:"carbon_redcing" format:"%f"`
	WindMPPTPower float32					`json:"WindMPPTPower" sql:"wind_mppt_power" format:"%f"`
	TodayGreenEnergy float32				`json:"TodayGreenEnergy" sql:"today_green_energy" format:"%f"`
	BatterySOC float32						`json:"BatterySOC" sql:"battery_soc" format:"%f"`
	AmbientTemperature float32				`json:"AmbientTemperature" sql:"ambient_temperature" format:"%f"`
	PVMPPTPower float32						`json:"PVMPPTPower" sql:"pv_mppt_power" format:"%f"`
}

func MainDecode(input_val []uint16, input_id int32, input_chanel int32) MainCell {
	var tmp MainCell
	tmp.ID = input_id
	tmp.Chanel = input_chanel
	tmp.SystemPower = int32_to_int(input_val[1], input_val[0], 0.1)
	tmp.GridPower = int32_to_int(input_val[3], input_val[2], 0.1)
	tmp.ChargingPower = int32_to_int(input_val[5], input_val[4], 0.1)
	tmp.DischargingPower = int32_to_int(input_val[7], input_val[6], 0.1)
	tmp.MPPTPower = int32_to_int(input_val[9], input_val[8], 0.1)
	tmp.TotalGreenEnergy = int32_to_int(input_val[11], input_val[10], 0.001)
	tmp.CarbonRedcing = int32_to_int(input_val[13], input_val[12], 0.01)
	tmp.WindMPPTPower = int32_to_int(input_val[15], input_val[14], 0.1)
	tmp.TodayGreenEnergy = int32_to_int(input_val[17], input_val[16], 0.001)
	tmp.BatterySOC = int32_to_int(input_val[19], input_val[18], 0.1)
	tmp.AmbientTemperature = int16_to_int(input_val[20], 0.01)
	tmp.PVMPPTPower = int32_to_int(input_val[23], input_val[22], 0.1)
	return tmp
}