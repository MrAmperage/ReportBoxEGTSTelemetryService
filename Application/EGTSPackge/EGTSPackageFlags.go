package EGTSPackge

import "fmt"

type EGTSPackgeHeaderFlags struct {
	Prefix              uint8
	Route               bool
	EncryptionAlgorithm uint8
	Compression         bool
	Priority            uint8
}

/*Декодер байтовых флагов*/
func (EGTSPackgeHeaderFlags *EGTSPackgeHeaderFlags) DecodeFlags(Data []byte) (Prefix int, Route bool, EncryptionAlgorithm int, Compression bool, Priority int, Error error) {
	var PackageLength = len(Data)
	if PackageLength < 3 {
		return Prefix, Route, EncryptionAlgorithm, Compression, Priority, fmt.Errorf("В сообщении нет Flags")
	}
	Prefix = int((Data[2] >> 6) & 0x03)
	Route = int((Data[2]>>5)&0x01) != 0
	EncryptionAlgorithm = int((Data[2] >> 3) & 0x03)
	Compression = int((Data[2]>>2)&0x01) != 0
	Priority = int(Data[2] & 0x03)
	return Prefix, Route, EncryptionAlgorithm, Compression, Priority, Error

}
