package main

import "time"

func Byte_Converter (data []byte) uint {
	var i uint
	var sum uint
	for i = 0; int(i) < len(data); i++ {
		sum += uint(data[i]) << (8 * i)
	}
	return sum
}

func Int_Converter (num uint) []byte {
	data := make([]byte, 4)
	for i := 0; i < 4; i++ {
		mod := num % 256
		data[i] = byte(mod)
		num /= 256;
	}
	return data
}

func Time_Stamp (data []byte) []byte {
	new_data := make([]byte, len(data) + 4)
	copy(new_data[0:12], data[0:12])
	for i := 13; i < len(data); i++ {
		new_data[i + 4] = data[i]
	}
	copy(new_data[13:16], Int_Converter(uint(time.Now().UnixNano() / (1000000))))
	return new_data
}