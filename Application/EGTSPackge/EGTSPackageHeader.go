package EGTSPackge

import "bytes"

/*Структура для заголовка пакета*/
type EGTSPackgeHeader struct {
	/*Версия протокола*/
	ProtocolVersion uint8
	/*Security Key*/
	SecurityKeyId uint8
	/*Флаги*/
	Flags EGTSPackgeHeaderFlags
	/*Длинна заголовка*/
	HeaderLength uint8
	/*Метод кодирования*/
	HeaderEncoding uint8
	/*Длина пакета данных*/
	FrameDataLength uint16
	/*Id пакета*/
	PacketIdentifier uint16
	/*Тип пакета*/
	PacketType uint8
	/*Адрес сгенерировавший пакет*/
	PeerAddress uint16
	/*Адрес для которого пакет предназначен*/
	RecipientAddress uint16
	/*Время жизни пакета*/
	TimeToLive uint8
	/*Контрольная сумма заголовка*/
	HeaderCheckSum uint8
}

/*Декодер заголовка*/
func (EGTSPackgeHeader *EGTSPackgeHeader) Decode(Reader *bytes.Reader) (Error error) {

	return Error
}

/*Текстовое представление типа пакета*/
func (EGTSPackgeHeader *EGTSPackgeHeader) PacketTypeToString(PacketType uint8) string {
	switch PacketType {
	case 0:
		return "EGTS_PT_RESPONSE"
	case 1:
		return "EGTS_PT_APPDATA"
	case 2:
		return "EGTS_PT_SIGNED_APPDATA"
	default:
		return "Unknown"
	}

}
