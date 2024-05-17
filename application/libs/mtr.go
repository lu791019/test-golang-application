package libs

import (
//	"log"
//	"fmt"
//	"github.com/simonvetter/modbus"
)

type MTRCell struct {
	ID int32								`json:"ID" sql:"meter_id" format:"%d"`
	Chanel int32							`json:"Chanel" sql:"chanel" format:"%d"`
    Frequency float32						`json:"Frequency" sql:"frequency" format:"%f"`
    Phase1Voltage float32					`json:"Phase1Voltage" sql:"phase_1_voltage" format:"%f"`
    Phase2Voltage float32					`json:"Phase2Voltage" sql:"phase_2_voltage" format:"%f"`
    Phase3Voltage float32					`json:"Phase3Voltage" sql:"phase_3_voltage" format:"%f"`
	AveragePhaseVltage float32				`json:"AveragePhaseVltage" sql:"average_phase_vltage" format:"%f"`
	Phase12LineVoltage float32				`json:"Phase12LineVoltage" sql:"phase_12line_voltage" format:"%f"`
	Phase23LineVoltage float32				`json:"Phase23LineVoltage" sql:"phase_23line_voltage" format:"%f"`
	Phase31LineVoltage float32				`json:"Phase31LineVoltage" sql:"phase_31line_voltage" format:"%f"`
	AverageLineVoltage float32				`json:"AverageLineVoltage" sql:"average_line_voltage" format:"%f"`
	I1Current float32						`json:"I1Current" sql:"i1current" format:"%f"`
	I2Current float32						`json:"I2Current" sql:"i2current" format:"%f"`
	I3Current float32						`json:"I3Current" sql:"i3current" format:"%f"`
	AverageCurrent float32					`json:"AverageCurrent" sql:"average_current" format:"%f"`
	NeutralCurrent float32					`json:"NeutralCurrent" sql:"neutral_current" format:"%f"`
	Phase1ActiverPower float32				`json:"Phase1ActiverPower" sql:"phase_1activer_power" format:"%f"`
	Phase2ActiverPower float32				`json:"Phase2ActiverPower" sql:"phase_2activer_power" format:"%f"`
	Phase3ActiverPower float32				`json:"Phase3ActiverPower" sql:"phase_3activer_power" format:"%f"`
	TotalActivePower float32				`json:"TotalActivePower" sql:"total_active_power" format:"%f"`
	Phase1ReactivePower float32				`json:"Phase1ReactivePower" sql:"phase_1reactive_power" format:"%f"`
	Phase2ReactivePower float32				`json:"Phase2ReactivePower" sql:"phase_2reactive_power" format:"%f"`
	Phase3ReactivePower float32				`json:"Phase3ReactivePower" sql:"phase_3reactive_power" format:"%f"`
	TotalApparentPower float32				`json:"TotalApparentPower" sql:"total_apparent_power" format:"%f"`
	Phase1PowerFactor float32				`json:"Phase1PowerFactor" sql:"phase_1power_factor" format:"%f"`
	Phase2PowerFactor float32				`json:"Phase2PowerFactor" sql:"phase_2power_factor" format:"%f"`
	Phase3PowerFactor float32				`json:"Phase3PowerFactor" sql:"phase_3power_factor" format:"%f"`
	AveragePowerFactor float32				`json:"AveragePowerFactor" sql:"average_power_factor" format:"%f"`
	AeImp float32							`json:"AeImp" sql:"aeimp" format:"%f"`
	AeExp float32							`json:"AeExp" sql:"aeexp" format:"%f"`
	AeTotal float32							`json:"AeTotal" sql:"aetotal" format:"%f"`
	AeNet float32							`json:"AeNet" sql:"aenet" format:"%f"`
	ReImp float32							`json:"ReImp" sql:"reimp" format:"%f"`
	ReExp float32							`json:"ReExp" sql:"reexp" format:"%f"`
	ReTotal float32							`json:"ReTotal" sql:"retotal" format:"%f"`
	ReNet float32							`json:"ReNet" sql:"renet" format:"%f"`
	SeTotal float32							`json:"SeTotal" sql:"setotal" format:"%f"`
}

func MTR_Cell_Decode(input_val []uint16, input_id int32, input_chanel int32) MTRCell {
	var tmp MTRCell
	tmp.ID = input_id
	tmp.Chanel = input_chanel
	tmp.Frequency = int16_to_int(input_val[0], 0.01)
	tmp.Phase1Voltage = int32_to_int(input_val[2], input_val[1], 0.1)
	tmp.Phase2Voltage = int32_to_int(input_val[4], input_val[3], 0.1)
	tmp.Phase3Voltage = int32_to_int(input_val[6], input_val[5], 0.1)
	tmp.AveragePhaseVltage = int32_to_int(input_val[8], input_val[7], 0.1)
	tmp.Phase12LineVoltage = int32_to_int(input_val[10], input_val[9], 0.1)
	tmp.Phase23LineVoltage = int32_to_int(input_val[12], input_val[11], 0.1)
	tmp.Phase31LineVoltage = int32_to_int(input_val[14], input_val[13], 0.1)
	tmp.AverageLineVoltage = int32_to_int(input_val[16], input_val[15], 0.1)
	tmp.I1Current = int32_to_int(input_val[18], input_val[17], 0.001)
	tmp.I2Current = int32_to_int(input_val[20], input_val[19], 0.001)
	tmp.I3Current = int32_to_int(input_val[22], input_val[21], 0.001)
	tmp.AverageCurrent = int32_to_int(input_val[24], input_val[23], 0.001)
	tmp.NeutralCurrent = int32_to_int(input_val[26], input_val[25], 0.001)
	tmp.Phase1ActiverPower = int32_to_int(input_val[28], input_val[27], 1)
	tmp.Phase2ActiverPower = int32_to_int(input_val[30], input_val[29], 1)
	tmp.Phase3ActiverPower = int32_to_int(input_val[32], input_val[31], 1)
	tmp.TotalActivePower = int32_to_int(input_val[34], input_val[33], 1)
	tmp.Phase1ReactivePower = int32_to_int(input_val[36], input_val[35], 1)
	tmp.Phase2ReactivePower = int32_to_int(input_val[38], input_val[37], 1)
	tmp.Phase3ReactivePower = int32_to_int(input_val[40], input_val[39], 1)
	tmp.TotalApparentPower = int32_to_int(input_val[42], input_val[41], 1)
	tmp.Phase1PowerFactor = int32_to_int(input_val[44], input_val[43], 1)
	tmp.Phase2PowerFactor = int32_to_int(input_val[46], input_val[45], 1)
	tmp.Phase3PowerFactor = int32_to_int(input_val[48], input_val[47], 1)
	tmp.TotalApparentPower = int32_to_int(input_val[50], input_val[49], 1)
	tmp.Phase1PowerFactor = int16_to_int(input_val[51], 0.001)
	tmp.Phase2PowerFactor = int16_to_int(input_val[52], 0.001)
	tmp.Phase3PowerFactor = int16_to_int(input_val[53], 0.001)
	tmp.AveragePowerFactor = int16_to_int(input_val[54], 0.001)
	tmp.AeImp = int32_to_int(input_val[56], input_val[55], 0.1)
	tmp.AeExp = int32_to_int(input_val[58], input_val[57], 0.1)
	tmp.AeTotal = int32_to_int(input_val[60], input_val[59], 0.1)
	tmp.AeNet = int32_to_int(input_val[62], input_val[61], 0.1)
	tmp.ReImp = int32_to_int(input_val[64], input_val[62], 0.1)
	tmp.ReExp = int32_to_int(input_val[66], input_val[65], 0.1)
	tmp.ReTotal = int32_to_int(input_val[68], input_val[67], 0.1)
	tmp.ReNet = int32_to_int(input_val[70], input_val[69], 0.1)
	tmp.SeTotal = int32_to_int(input_val[72], input_val[71], 0.1)
	return tmp
}















