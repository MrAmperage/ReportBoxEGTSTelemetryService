package EGTSPackge

import "bytes"

/*Заголовок записи*/
type RecordHeader struct {
	RecordLength         uint16
	RecordNumber         uint16
	RecordFlags          EGTSRecordHeaderFlags
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

	SourceServiceOnDevice, RecipientServiceOnDevice, Group, RecordProcessingPriority, TimeFieldExist, EventIdFieldExist, ObjectIdFieldExist, Error := RecordHeader.RecordFlags.DecodeFlags(Reader)
	if Error != nil {
		return Error
	}
	RecordHeader.RecordFlags.SourceServiceOnDevice = SourceServiceOnDevice
	RecordHeader.RecordFlags.RecipientServiceOnDevice = RecipientServiceOnDevice
	RecordHeader.RecordFlags.Group = Group
	RecordHeader.RecordFlags.RecordProcessingPriority = RecordProcessingPriority
	RecordHeader.RecordFlags.TimeFieldExist = TimeFieldExist
	RecordHeader.RecordFlags.EventIdFieldExist = EventIdFieldExist
	RecordHeader.RecordFlags.ObjectIdFieldExist = ObjectIdFieldExist

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
