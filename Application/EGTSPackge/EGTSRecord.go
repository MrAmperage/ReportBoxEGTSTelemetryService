package EGTSPackge

/*Запись*/
type Record struct {
	RecordHeader RecordHeader
}

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
