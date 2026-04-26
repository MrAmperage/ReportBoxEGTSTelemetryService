package EGTSPackge

import "bytes"

/*Подзапись*/
type SubRecord struct {
	SubRecordType   uint8
	SubRecordLength uint16
}

/*Декодирование подзаписи*/
func DecodeSubRecord(Reader *bytes.Reader) (SubRecord SubRecord, Error error) {

	SubRecordType, Error := DecodeSubRecordType(Reader)

	if Error != nil {
		return SubRecord, Error
	}
	SubRecord.SubRecordType = SubRecordType
	return SubRecord, Error
}

/*Декодировать SubRecordType*/
func DecodeSubRecordType(Reader *bytes.Reader) (SubRecordType uint8, Error error) {
	return ReadByte(Reader)
}

/*Текстовое представление типа подзаписи*/
func (RecordHeader *RecordHeader) SubRecordTypeToString(SubRecordType uint8) string {
	switch SubRecordType {
	case 1:
		return "EGTS_SR_TERM_IDENTITY"
	default:
		return "Unknown"
	}

}
