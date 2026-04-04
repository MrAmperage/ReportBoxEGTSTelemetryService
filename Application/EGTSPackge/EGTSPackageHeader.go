package EGTSPackge

import (
	"ReportBoxEGTSTelemetryService/Application/BasePackage"
	"bytes"
)

/*Структура для заголовка пакета*/
type EGTSPackgeHeader struct {
	/*Базовые инструменты для манипуляции с данными*/
	BasePackage BasePackage.BasePackage
	/*Версия*/
	Version int
}

/*Декодер заголовка*/
func (EGTSPackgeHeader *EGTSPackgeHeader) Decode(Reader *bytes.Reader) {
	EGTSPackgeHeader.DecodeVersion(Reader)

}

/*Декодер версии*/
func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeVersion(Reader *bytes.Reader) (Error error) {
	ByteVersion, Error := EGTSPackgeHeader.BasePackage.ReadNext(Reader, 1)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.Version = int(ByteVersion[0])
	return Error

}
