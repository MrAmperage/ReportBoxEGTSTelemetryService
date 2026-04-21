package EGTSPackge

import (
	"bytes"
	"encoding/binary"
	"io"
)

/*Считывает из ридера определенное количество байт*/
func ReadSequenceBytes(Reader *bytes.Reader, BytesCount int) (Bytes []byte, Error error) {
	Buffer := make([]byte, BytesCount)
	_, Error = io.ReadFull(Reader, Buffer)
	if Error != nil {
		return Bytes, Error
	}
	Bytes = Buffer
	return Bytes, Error
}

/*Считать из ридера тип данных Byte*/
func ReadByte(Reader *bytes.Reader) (Number uint8, Error error) {
	Bytes, Error := ReadSequenceBytes(Reader, 1)
	if Error != nil {
		return Number, Error
	}
	Number = Bytes[0]
	return Number, Error
}

/*Считать из ридера тип данных Ushort*/
func ReadUshort(Reader *bytes.Reader) (Number uint16, Error error) {
	Bytes, Error := ReadSequenceBytes(Reader, 2)
	if Error != nil {
		return Number, Error
	}
	Number = binary.LittleEndian.Uint16(Bytes)
	return Number, Error
}

/*Считать из ридера тип данных Uint*/
func ReadUint(Reader *bytes.Reader) (Number uint32, Error error) {
	Bytes, Error := ReadSequenceBytes(Reader, 4)
	if Error != nil {
		return Number, Error
	}
	Number = binary.LittleEndian.Uint32(Bytes)
	return Number, Error
}
