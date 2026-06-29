package EGTSPackge

import "ReportBoxEGTSTelemetryService/Application/SubRecord"

/*Запись*/
type Record struct {
	RecordHeader RecordHeader
	SubRecords   []SubRecord.SubRecord[SubRecord.SubRecordType]
}
