package EGTSPackge

import (
	"ReportBoxEGTSTelemetryService/Application/Helpers"
	"bytes"
)

type EGTSPackgeHeaderFlags struct {
	Prefix              uint8
	Route               bool
	EncryptionAlgorithm uint8
	Compression         bool
	Priority            uint8
}

/*Декодер байтовых флагов*/
func (EGTSPackgeHeaderFlags *EGTSPackgeHeaderFlags) DecodeFlags(Reader *bytes.Reader) (Prefix uint8, Route bool, EncryptionAlgorithm uint8, Compression bool, Priority uint8, Error error) {
	Data, Error := Helpers.ReadByte(Reader)
	Prefix = (Data >> 6) & 0x03
	Route = (Data>>5)&0x01 != 0
	EncryptionAlgorithm = (Data >> 3) & 0x03
	Compression = (Data>>2)&0x01 != 0
	Priority = Data & 0x03
	return Prefix, Route, EncryptionAlgorithm, Compression, Priority, Error

}
