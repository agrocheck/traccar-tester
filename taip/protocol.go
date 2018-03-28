package taip

import (
	"bytes"
	"encoding/hex"
)

type Device struct {
	id string
	event string
}

func NewDevice(id string, event string) (p *Device) {
	p = &Device{
		id: id,
		event: event,
	}

	return
}

func (p *Device) RandomMessage() string {
	date := "080810"
	time := "000044"
	latitude := "-3460365"
	longitude := "-05838146"
	speed := "000"
	course := "000"
	input := "FF"
	output := "00"
	battery := "131"
	odometer := "00000035"
	gpsPower := "1"
	fixMode := "1"
	pdop := "50"
	satellites := "00"
	dataAge := "FFFF"
	modemPower := "1"
	gsmStatus := "0"
	rssi := "14"
	index := "0000"

	buffer := bytes.Buffer{}
	buffer.WriteString(">RCQ")
	buffer.WriteString(p.event)
	buffer.WriteString(date)
	buffer.WriteString(time)
	buffer.WriteString(latitude)
	buffer.WriteString(longitude)
	buffer.WriteString(speed)
	buffer.WriteString(course)
	buffer.WriteString(input)
	buffer.WriteString(output)
	buffer.WriteString(battery)
	buffer.WriteString(odometer)
	buffer.WriteString(gpsPower)
	buffer.WriteString(fixMode)
	buffer.WriteString(pdop)
	buffer.WriteString(satellites)
	buffer.WriteString(dataAge)
	buffer.WriteString(modemPower)
	buffer.WriteString(gsmStatus)
	buffer.WriteString(rssi)
	buffer.WriteString(";#")
	buffer.WriteString(index)
	buffer.WriteString(";ID=")
	buffer.WriteString(p.id)
	buffer.WriteString(";*")
	buffer.Write(checksum(buffer))
	buffer.WriteString("<")

	return buffer.String()
}

// TODO: cleanup
func checksum(buffer bytes.Buffer) []byte {
	src := buffer.Bytes()

	result := byte(0)
	for _, b := range src {
		result ^= b
	}

	src = []byte{result}
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	return dst
}
