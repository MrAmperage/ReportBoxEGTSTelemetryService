package EGTSPackge

import (
	"ReportBoxEGTSTelemetryService/Application/BasePackage"
	"bytes"
	"io"
)

type EGTSPackgeHeaderFlags struct {
	BasePackage         BasePackage.BasePackage
	Prefix              int
	Route               bool
	EncryptionAlgorithm int
	Compression         bool
	Priority            int
}

/*Декодер байтовых флагов*/
func (EGTSPackgeHeaderFlags *EGTSPackgeHeaderFlags) DecodeFlags(Reader *bytes.Reader) (Prefix int, Route bool, EncryptionAlgorithm int, Compression bool, Priority int, Error error) {
	Reader.Seek(2, io.SeekStart)
	BytesFlags, Error := EGTSPackgeHeaderFlags.BasePackage.ReadNext(Reader, 1)

	if Error != nil {
		return Prefix, Route, EncryptionAlgorithm, Compression, Priority, Error
	}
	Prefix = int((BytesFlags[0] >> 6) & 0x03)
	Route = int((BytesFlags[0]>>5)&0x01) != 0
	EncryptionAlgorithm = int((BytesFlags[0] >> 3) & 0x03)
	Compression = int((BytesFlags[0]>>2)&0x01) != 0
	Priority = int(BytesFlags[0] & 0x03)
	return Prefix, Route, EncryptionAlgorithm, Compression, Priority, Error

}
