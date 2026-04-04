package EGTSPackge

import (
	"ReportBoxEGTSTelemetryService/Application/BasePackage"
	"bytes"
)

/*Структура для заголовка пакета*/
type EGTSPackgeHeader struct {
	/*Базовые инструменты для манипуляции с данными*/
	BasePackage BasePackage.BasePackage
	/*Версия протокола*/
	ProtocolVersion int
	SecurityKeyId   int
}

/*Декодер заголовка*/
func (EGTSPackgeHeader *EGTSPackgeHeader) Decode(Reader *bytes.Reader) {
	EGTSPackgeHeader.DecodeProtocolVersion(Reader)
	EGTSPackgeHeader.DecodeSecurityKeyId(Reader)

}

/*Декодер версии протокола*/
func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeProtocolVersion(Reader *bytes.Reader) (Error error) {
	BytesVersion, Error := EGTSPackgeHeader.BasePackage.ReadNext(Reader, 1)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.ProtocolVersion = int(BytesVersion[0])
	return Error

}

/*Декодер Security Key Id*/
func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeSecurityKeyId(Reader *bytes.Reader) (Error error) {
	BytesSecurityKeyId, Error := EGTSPackgeHeader.BasePackage.ReadNext(Reader, 1)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.SecurityKeyId = int(BytesSecurityKeyId[0])
	return Error

}
