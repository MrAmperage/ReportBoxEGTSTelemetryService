package EGTSPackge

import (
	"fmt"
)

/*Структура для заголовка пакета*/
type EGTSPackgeHeader struct {
	/*Версия протокола*/
	ProtocolVersion int
	/*Security Key*/
	SecurityKeyId int
	/*Флаги*/
	EGTSPackgeHeaderFlags EGTSPackgeHeaderFlags
	/*Длинна заголовка*/
	HeaderLength int
	/*Метод кодирования*/
	HeaderEncoding int
}

/*Декодер заголовка*/
func (EGTSPackgeHeader *EGTSPackgeHeader) Decode(Data []byte) (Error error) {
	Version, Error := EGTSPackgeHeader.DecodeProtocolVersion(Data)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.ProtocolVersion = Version
	SecurityKeyId, Error := EGTSPackgeHeader.DecodeSecurityKeyId(Data)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.SecurityKeyId = SecurityKeyId
	Prefix, Route, EncryptionAlgorithm, Compression, Priority, Error := EGTSPackgeHeader.EGTSPackgeHeaderFlags.DecodeFlags(Data)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.EGTSPackgeHeaderFlags.Prefix = Prefix
	EGTSPackgeHeader.EGTSPackgeHeaderFlags.Route = Route
	EGTSPackgeHeader.EGTSPackgeHeaderFlags.EncryptionAlgorithm = EncryptionAlgorithm
	EGTSPackgeHeader.EGTSPackgeHeaderFlags.Compression = Compression
	EGTSPackgeHeader.EGTSPackgeHeaderFlags.Priority = Priority
	HeaderLength, Error := EGTSPackgeHeader.DecodeHeaderLength(Data)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.HeaderLength = HeaderLength
	HeaderEncoding, Error := EGTSPackgeHeader.DecodeHeaderEncoding(Data)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.HeaderEncoding = HeaderEncoding
	return Error
}

/*Декодер версии протокола*/
func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeProtocolVersion(Data []byte) (ProtocolVersion int, Error error) {
	var PackageLength = len(Data)
	if PackageLength < 1 {
		return ProtocolVersion, fmt.Errorf("В сообщении нет Protocol Version")
	}
	ProtocolVersion = int(Data[0])
	if ProtocolVersion != 1 {
		return ProtocolVersion, fmt.Errorf("Версия протокола не поддерживается - %d", ProtocolVersion)
	}
	return ProtocolVersion, Error
}

/*Декодер Security Key Id*/
func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeSecurityKeyId(Data []byte) (SecurityKeyId int, Error error) {
	var PackageLength = len(Data)
	if PackageLength < 2 {
		return SecurityKeyId, fmt.Errorf("В сообщении нет Security Key Id")
	}
	SecurityKeyId = int(Data[1])
	return SecurityKeyId, Error
}

/*Декодер Security Key Id*/
func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeHeaderLength(Data []byte) (HeaderLength int, Error error) {
	var PackageLength = len(Data)
	if PackageLength < 4 {
		return HeaderLength, fmt.Errorf("В сообщении нет Header Length")
	}
	HeaderLength = int(Data[3])
	if HeaderLength < 11 {
		return HeaderLength, fmt.Errorf("Неправильная длина заголовка - %d", HeaderLength)
	}
	return HeaderLength, Error
}

/*Декодер HeaderEncoding*/
func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeHeaderEncoding(Data []byte) (HeaderEncoding int, Error error) {
	var PackageLength = len(Data)
	if PackageLength < 5 {
		return HeaderEncoding, fmt.Errorf("В сообщении нет Header Encoding")
	}
	HeaderEncoding = int(Data[4])
	return HeaderEncoding, Error

}
