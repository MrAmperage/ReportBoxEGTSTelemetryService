package EGTSPackge

import (
	"ReportBoxEGTSTelemetryService/Application/BasePackage"
	"bytes"
	"fmt"
	"io"
)

/*Структура для заголовка пакета*/
type EGTSPackgeHeader struct {
	/*Базовые инструменты для манипуляции с данными*/
	BasePackage BasePackage.BasePackage
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
func (EGTSPackgeHeader *EGTSPackgeHeader) Decode(Reader *bytes.Reader) (Error error) {
	Version, Error := EGTSPackgeHeader.DecodeProtocolVersion(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.ProtocolVersion = Version
	SecurityKeyId, Error := EGTSPackgeHeader.DecodeSecurityKeyId(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.SecurityKeyId = SecurityKeyId
	Prefix, Route, EncryptionAlgorithm, Compression, Priority, Error := EGTSPackgeHeader.EGTSPackgeHeaderFlags.DecodeFlags(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.EGTSPackgeHeaderFlags.Prefix = Prefix
	EGTSPackgeHeader.EGTSPackgeHeaderFlags.Route = Route
	EGTSPackgeHeader.EGTSPackgeHeaderFlags.EncryptionAlgorithm = EncryptionAlgorithm
	EGTSPackgeHeader.EGTSPackgeHeaderFlags.Compression = Compression
	EGTSPackgeHeader.EGTSPackgeHeaderFlags.Priority = Priority
	HeaderLength, Error := EGTSPackgeHeader.DecodeHeaderLength(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.HeaderLength = HeaderLength
	HeaderEncoding, Error := EGTSPackgeHeader.DecodeHeaderEncoding(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.HeaderEncoding = HeaderEncoding
	return Error
}

/*Декодер версии протокола*/
func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeProtocolVersion(Reader *bytes.Reader) (ProtocolVersion int, Error error) {
	Reader.Seek(0, io.SeekStart)
	BytesVersion, Error := EGTSPackgeHeader.BasePackage.ReadNext(Reader, 1)
	if Error != nil {
		return ProtocolVersion, Error
	}
	ProtocolVersion = int(BytesVersion[0])
	if ProtocolVersion != 1 {
		return ProtocolVersion, fmt.Errorf("Версия протокола не поддерживается - %d", ProtocolVersion)
	}
	return ProtocolVersion, Error
}

/*Декодер Security Key Id*/
func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeSecurityKeyId(Reader *bytes.Reader) (SecurityKeyId int, Error error) {
	Reader.Seek(1, io.SeekStart)
	BytesSecurityKeyId, Error := EGTSPackgeHeader.BasePackage.ReadNext(Reader, 1)
	if Error != nil {
		return SecurityKeyId, Error
	}
	SecurityKeyId = int(BytesSecurityKeyId[0])
	return SecurityKeyId, Error

}

/*Декодер Security Key Id*/
func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeHeaderLength(Reader *bytes.Reader) (HeaderLength int, Error error) {
	Reader.Seek(3, io.SeekStart)
	BytesHeaderLength, Error := EGTSPackgeHeader.BasePackage.ReadNext(Reader, 1)
	if Error != nil {
		return HeaderLength, Error
	}
	HeaderLength = int(BytesHeaderLength[0])
	if HeaderLength < 11 || HeaderLength > 1024 {
		return HeaderLength, fmt.Errorf("Неправильная длина заголовка - %d", HeaderLength)
	}
	return HeaderLength, Error

}

/*Декодер HeaderEncoding*/
func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeHeaderEncoding(Reader *bytes.Reader) (HeaderEncoding int, Error error) {
	Reader.Seek(4, io.SeekStart)
	BytesHeaderEncoding, Error := EGTSPackgeHeader.BasePackage.ReadNext(Reader, 1)
	if Error != nil {
		return HeaderEncoding, Error
	}
	HeaderEncoding = int(BytesHeaderEncoding[0])

	return HeaderEncoding, Error

}
