package EGTSPackge

import (
	"ReportBoxEGTSTelemetryService/Application/Helpers"
	"bytes"
)

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

	if RecordHeader.RecordFlags.ObjectIdFieldExist {
		ObjectIdentifier, Error := DecodeObjectIdentifier(Reader)
		if Error != nil {
			return Error
		}
		RecordHeader.ObjectIdentifier = ObjectIdentifier
	}
	if RecordHeader.RecordFlags.EventIdFieldExist {
		EventIdentifier, Error := DecodeEventIdentifier(Reader)
		if Error != nil {
			return Error
		}
		RecordHeader.EventIdentifier = EventIdentifier
	}

	if RecordHeader.RecordFlags.TimeFieldExist {
		Time, Error := DecodeTime(Reader)
		if Error != nil {
			return Error
		}
		RecordHeader.Time = Time
	}
	SourceServiceType, Error := DecodeSourceServiceType(Reader)
	if Error != nil {
		return Error
	}
	RecordHeader.SourceServiceType = SourceServiceType

	RecipientServiceType, Error := DecodeRecipientServiceType(Reader)
	if Error != nil {
		return Error
	}
	RecordHeader.RecipientServiceType = RecipientServiceType
	return Error
}

/*Декодировать длину записи*/
func DecodeRecordLength(Reader *bytes.Reader) (RecordLength uint16, Error error) {
	return Helpers.ReadUshort(Reader)
}

/*Декодировать номер записи*/
func DecodeRecordNumber(Reader *bytes.Reader) (RecordNumber uint16, Error error) {
	return Helpers.ReadUshort(Reader)
}

/*Декодировать Id объекта*/
func DecodeObjectIdentifier(Reader *bytes.Reader) (ObjectIdentifier uint32, Error error) {
	return Helpers.ReadUint(Reader)
}

/*Декодировать EventId*/
func DecodeEventIdentifier(Reader *bytes.Reader) (EventIdentifier uint32, Error error) {
	return Helpers.ReadUint(Reader)
}

/*Декодировать Time*/
func DecodeTime(Reader *bytes.Reader) (Time uint32, Error error) {
	return Helpers.ReadUint(Reader)
}

/*Декодировать SourceServiceType*/
func DecodeSourceServiceType(Reader *bytes.Reader) (SourceServiceType uint8, Error error) {
	return Helpers.ReadByte(Reader)
}

/*Декодировать RecipientServiceType*/
func DecodeRecipientServiceType(Reader *bytes.Reader) (RecipientServiceType uint8, Error error) {
	return Helpers.ReadByte(Reader)
}

/*Текстовое представление типа записи*/
func (RecordHeader *RecordHeader) RecordTypeToString(RecordType uint8) string {
	switch RecordType {
	case 1:
		return "EGTS_AUTH_SERVICE"
	default:
		return "Unknown"
	}

}
