package EGTSPackge

import "bytes"

/*Флаги для еденичной записи*/
type EGTSRecordHeaderFlags struct {
	SourceServiceOnDevice    bool
	RecipientServiceOnDevice bool
	Group                    bool
	RecordProcessingPriority uint8
	TimeFieldExist           bool
	EventIdFieldExist        bool
	ObjectIdFieldExist       bool
}

func (EGTSRecordHeaderFlags *EGTSRecordHeaderFlags) DecodeFlags(Reader *bytes.Reader) (
	SourceServiceOnDevice bool,
	RecipientServiceOnDevice bool,
	Group bool,
	RecordProcessingPriority uint8,
	TimeFieldExist bool,
	EventIdFieldExist bool,
	ObjectIdFieldExist bool,
	Error error,
) {
	ByteFlags, Error := ReadByte(Reader)
	if Error != nil {
		return SourceServiceOnDevice,
			RecipientServiceOnDevice,
			Group,
			RecordProcessingPriority,
			TimeFieldExist,
			EventIdFieldExist,
			ObjectIdFieldExist,
			Error
	}

	SourceServiceOnDevice = (ByteFlags>>7)&0x01 != 0
	RecipientServiceOnDevice = (ByteFlags>>6)&0x01 != 0
	Group = (ByteFlags>>5)&0x01 != 0
	RecordProcessingPriority = (ByteFlags >> 3) & 0x03
	TimeFieldExist = (ByteFlags>>2)&0x01 != 0
	EventIdFieldExist = (ByteFlags>>1)&0x01 != 0
	ObjectIdFieldExist = ByteFlags&0x01 != 0

	return SourceServiceOnDevice,
		RecipientServiceOnDevice,
		Group,
		RecordProcessingPriority,
		TimeFieldExist,
		EventIdFieldExist,
		ObjectIdFieldExist,
		Error
}
