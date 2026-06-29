package SubRecord

import (
	"ReportBoxEGTSTelemetryService/Application/Helpers"
	"ReportBoxEGTSTelemetryService/Application/SubRecord/EGTS_SR_TERM_IDENTITY"
	"bytes"
)

type SubRecordType = EGTS_SR_TERM_IDENTITY.EGTS_SR_TERM_IDENTITY

/*Подзапись*/
type SubRecord[SybRecordData SubRecordType] struct {
	SubRecordType   uint8
	SubRecordLength uint16
	Data            SybRecordData
}

/*Декодирование подзаписи*/
func DecodeSubRecord(Reader *bytes.Reader) (SubRecord SubRecord[SubRecordType], Error error) {

	SubRecordType, Error := DecodeSubRecordType(Reader)

	if Error != nil {
		return SubRecord, Error
	}
	SubRecord.SubRecordType = SubRecordType

	SubRecordLength, Error := DecodeSubRecordLength(Reader)
	if Error != nil {
		return SubRecord, Error
	}
	SubRecord.SubRecordLength = SubRecordLength
	return SubRecord, Error
}

/*Декодировать SubRecordType*/
func DecodeSubRecordType(Reader *bytes.Reader) (SubRecordType uint8, Error error) {
	return Helpers.ReadByte(Reader)
}

/*Текстовое представление типа подзаписи*/
func SubRecordTypeToString(SubRecordType uint8) string {
	switch SubRecordType {
	case 1:
		return "EGTS_SR_TERM_IDENTITY"
	default:
		return "Unknown"
	}

}

/*Декодировать SubRecordLength*/
func DecodeSubRecordLength(Reader *bytes.Reader) (SubRecordLength uint16, Error error) {
	return Helpers.ReadUshort(Reader)
}
