package EGTSPackge

import "ReportBoxEGTSTelemetryService/Application/SubRecords"

/*Запись*/
type Record struct {
	RecordHeader RecordHeader
	SubRecords   []SubRecords.SubRecord
}
