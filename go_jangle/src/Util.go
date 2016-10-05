package main

import "time"

//Converts byte array to unsigned int 
func Byte_Converter (data []byte) uint {
	var i uint
	var sum uint
	for i = 0; int(i) < len(data); i++ {
		//Preforms little endian bit shifting and adds int value to sum for each byte
		sum += uint(data[i]) << (8 * i)
	}
	return sum
}

//Converts unsigned int to byte array
func Int_Converter (num uint) []byte {
	data := make([]byte, 4)
	for i := 0; i < 4; i++ {
		mod := num % 256
		data[i] = byte(mod)
		num /= 256
	}
	return data
}

//Returns Current Millisecond time as unsigned int
func Milli_Time() uint {
	return uint(time.Now().UnixNano() / 1000000)
}

//Used for time stamping code type 16 messages
//Takes in a byte array and creates a 4 byte space from byte 13 to 16
//Places 4 byte time stamp in space
func Time_Stamp (data []byte) []byte {
	new_data := make([]byte, len(data) + 4)
	copy(new_data[0:12], data[0:12])
	for i := 13; i < len(data); i++ {
		new_data[i + 4] = data[i]
	}
	copy(new_data[13:16], Int_Converter(Milli_Time()))
	return new_data
}
