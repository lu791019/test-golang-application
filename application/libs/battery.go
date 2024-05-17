package libs

import (
//	"log"
	"fmt"
	"github.com/simonvetter/modbus"
)

type BatteryCell struct {
	ID int32								`json:"ID" sql:"battery_id" format:"%d"`
	Chanel int32							`json:"Chanel" sql:"chanel" format:"%d"`
    BatteryStateOfCharge float32			`json:"BatteryStateOfCharge" sql:"charge" format:"%f"`
    PacketWarningFlagLow string				`json:"PacketWarningFlagLow" sql:"flag_low" format:"'%s'"`
    PacketWarningFlagHigh string			`json:"PacketWarningFlagHigh" sql:"flag_high" format:"'%s'"`
	Current float32							`json:"Current" sql:"current" format:"%f"`
	Voltage float32							`json:"Voltage" sql:"voltage" format:"%f"`
	Temperature float32						`json:"Temperature" sql:"temperature" format:"%f"`
	SOH float32								`json:"SOH" sql:"soh" format:"%f"`
}

//Read Battery(Start, End) Max 10 Unit
func Battery_Data(bcs *[]BatteryCell, mod_client *modbus.ModbusClient, cell_start int32, cell_end int32, input_id int32) {
	//var bcs []BatteryCell
	var err_bt	error
	var bt_data	[]uint16
	bt_data, err_bt = mod_client.ReadRegisters(uint16(800 + 10 * cell_start), uint16(10 * (cell_end - cell_start)), modbus.HOLDING_REGISTER)
	if err_bt != nil {
		return
	}
	var n int32 = 0
	for n < (cell_end - cell_start) {
		*bcs = append(*bcs, Battery_Cell_Decode(n, bt_data[n*10:n*10+10], input_id, n + cell_start))
		n++
	}
	//log.Println(bcs)
}

//Battery
func Battery_Cell_Decode(cell_number int32, input_val []uint16, input_id int32, input_chanel int32) BatteryCell {
	var bc BatteryCell
	bc.ID = input_id
	bc.Chanel = input_chanel
	//Battery state of charge
	var state_of_charge float32 = int16_to_int(input_val[0], 0.1)
	bc.BatteryStateOfCharge = state_of_charge
	//Battery Pack Warning Flag(Low Word)
	var pack_warning_flag_low int32 = int32(input_val[1])
	bc.PacketWarningFlagLow = fmt.Sprintf("0x%02x", pack_warning_flag_low)
	//Battery Pack Warning Flag(High Word)
	var pack_warning_flag_high int32 = int32(input_val[2])
	bc.PacketWarningFlagHigh = fmt.Sprintf("0x%02x", pack_warning_flag_high)
	//Battery Current
	var current float32 = int32_to_int(input_val[4], input_val[3], 0.001)
	bc.Current = current
	//Battery Voltage
	var voltage float32 = int32_to_int(input_val[6], input_val[5], 0.001)
	bc.Voltage = voltage
	//High Temperature
	var temperature float32 = int32_to_float(input_val[8], input_val[7], 1)
	bc.Temperature = temperature
	//S.O.H
	var soh float32 = int16_to_int(input_val[9], 1)
	bc.SOH = soh
	return bc
}