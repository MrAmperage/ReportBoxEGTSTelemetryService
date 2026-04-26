package EGTSPackge

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
