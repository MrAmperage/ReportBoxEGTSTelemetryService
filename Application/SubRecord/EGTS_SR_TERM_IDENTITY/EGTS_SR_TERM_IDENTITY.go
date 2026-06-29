package EGTS_SR_TERM_IDENTITY

/*Подзапись для идентификации терминала*/
type EGTS_SR_TERM_IDENTITY struct {
	TerminalIdentifier                    uint32
	Flags                                 EGTS_SR_TERM_IDENTITYFlags
	HomeDispatcherIdentifier              uint16
	InternationalMobileEquipmentIdentity  string
	InternationalMobileSubscriberIdentity string
	LanguageCode                          string
	//TODO Разобрать эту запись
	NetworkIdentifier                                   []byte
	BufferSize                                          uint16
	MobileStationIntegratedServicesDigitalNetworkNumber string
}
