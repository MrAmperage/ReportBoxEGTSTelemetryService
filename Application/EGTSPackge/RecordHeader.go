package EGTSPackge

import "bytes"

/*Заголовок записи*/
type RecordHeader struct {
	RecordLength         uint16
	RecordNumber         uint16
	RecordFlags          uint8
	ObjectIdentifier     uint32
	EventIdentifier      uint32
	Time                 uint32
	SourceServiceType    uint8
	RecipientServiceType uint8
}

func (RecordHeader *RecordHeader) DecodeRecordHeader(Reader *bytes.Reader) (Error error) {
	RecordLength, Error := DecodeRecordLength(Reader)
	if Error != nil {
		return Error
	}
	RecordHeader.RecordLength = RecordLength

	RecordNumber, Error := DecodeRecordNumber(Reader)
	if Error != nil {
		return Error
	}
	RecordHeader.RecordNumber = RecordNumber
	return Error
}

/*Декодировать длину записи*/
func DecodeRecordLength(Reader *bytes.Reader) (RecordLength uint16, Error error) {
	return ReadUshort(Reader)
}

/*Декодировать номер записи*/
func DecodeRecordNumber(Reader *bytes.Reader) (RecordNumber uint16, Error error) {
	return ReadUshort(Reader)
}
