package EGTSPackge

import (
	"bytes"
)

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
	ProtocolVersion, Error := EGTSPackgeHeader.DecodeProtocolVersion(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.ProtocolVersion = ProtocolVersion
	SecurityKeyId, Error := EGTSPackgeHeader.DecodeSecurityKeyId(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.SecurityKeyId = SecurityKeyId

	Prefix, Route, EncryptionAlgorithm, Compression, Priority, Error := EGTSPackgeHeader.Flags.DecodeFlags(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.Flags.Prefix = Prefix
	EGTSPackgeHeader.Flags.Route = Route
	EGTSPackgeHeader.Flags.EncryptionAlgorithm = EncryptionAlgorithm
	EGTSPackgeHeader.Flags.Compression = Compression
	EGTSPackgeHeader.Flags.Priority = Priority

	HeaderLength, Error := EGTSPackgeHeader.DecodeHeaderLength(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.HeaderLength = HeaderLength

	HeaderEncoding, Error := EGTSPackgeHeader.DecodeHeaderEncoding(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.HeaderEncoding = HeaderEncoding

	FrameDataLength, Error := EGTSPackgeHeader.DecodeFrameDataLength(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.FrameDataLength = FrameDataLength

	PacketIdentifier, Error := EGTSPackgeHeader.DecodePacketIdentifier(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.PacketIdentifier = PacketIdentifier
	return Error
}

func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeHeaderEncoding(Reader *bytes.Reader) (HeaderEncoding uint8, Error error) {
	HeaderEncoding, Error = ReadByte(Reader)
	return HeaderEncoding, Error

}
func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeFrameDataLength(Reader *bytes.Reader) (FrameDataLength uint16, Error error) {
	FrameDataLength, Error = ReadUshort(Reader)
	return FrameDataLength, Error
}

/*Декодирование версии протокола*/
func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeProtocolVersion(Reader *bytes.Reader) (ProtocolVersion uint8, Error error) {
	ProtocolVersion, Error = ReadByte(Reader)
	return ProtocolVersion, Error
}
func (EGTSPackgeHeader *EGTSPackgeHeader) DecodePacketIdentifier(Reader *bytes.Reader) (PacketIdentifier uint16, Error error) {
	PacketIdentifier, Error = ReadUshort(Reader)
	return PacketIdentifier, Error
}

func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeSecurityKeyId(Reader *bytes.Reader) (SecurityKeyId uint8, Error error) {
	SecurityKeyId, Error = ReadByte(Reader)
	return SecurityKeyId, Error

}

func (EGTSPackgeHeader *EGTSPackgeHeader) DecodeHeaderLength(Reader *bytes.Reader) (HeaderLength uint8, Error error) {
	HeaderLength, Error = ReadByte(Reader)
	return HeaderLength, Error
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
