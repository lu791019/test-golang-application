package libs

import (
	"math"
)

//處理uint數值
func int16_to_int(input_val uint16, input_magni float32) float32 {
	return float32(input_val) * input_magni
}
func int32_to_int(input_high uint16, input_low uint16, input_magni float32) float32 {
	return float32((int32(int(input_high) << 16) + int32(int(input_low)))) * input_magni
}
func int32_to_float(input_high uint16, input_low uint16, input_magni float32) float32 {
	return math.Float32frombits((uint32(input_high) << 16) + uint32(input_low)) * input_magni
}