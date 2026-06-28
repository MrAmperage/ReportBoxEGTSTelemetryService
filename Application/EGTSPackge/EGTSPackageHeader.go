package EGTSPackge

import (
	"ReportBoxEGTSTelemetryService/Application/Helpers"
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
	ProtocolVersion, Error := DecodeProtocolVersion(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.ProtocolVersion = ProtocolVersion
	SecurityKeyId, Error := DecodeSecurityKeyId(Reader)
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

	HeaderLength, Error := DecodeHeaderLength(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.HeaderLength = HeaderLength

	HeaderEncoding, Error := DecodeHeaderEncoding(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.HeaderEncoding = HeaderEncoding

	FrameDataLength, Error := DecodeFrameDataLength(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.FrameDataLength = FrameDataLength

	PacketIdentifier, Error := DecodePacketIdentifier(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.PacketIdentifier = PacketIdentifier

	PacketType, Error := DecodePacketType(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.PacketType = PacketType

	if EGTSPackgeHeader.Flags.Route {
		PeerAddress, Error := DecodePeerAddress(Reader)
		if Error != nil {
			return Error
		}
		EGTSPackgeHeader.PeerAddress = PeerAddress

		RecipientAddress, Error := DecodeRecipientAddress(Reader)
		if Error != nil {
			return Error
		}
		EGTSPackgeHeader.RecipientAddress = RecipientAddress

		TimeToLive, Error := DecodeTimeToLive(Reader)
		if Error != nil {
			return Error
		}
		EGTSPackgeHeader.TimeToLive = TimeToLive
	}

	HeaderCheckSum, Error := DecodeHeaderCheckSum(Reader)
	if Error != nil {
		return Error
	}
	EGTSPackgeHeader.HeaderCheckSum = HeaderCheckSum

	return Error
}
func DecodeTimeToLive(Reader *bytes.Reader) (TimeToLive uint8, Error error) {
	TimeToLive, Error = Helpers.ReadByte(Reader)
	return TimeToLive, Error
}

func DecodeHeaderCheckSum(Reader *bytes.Reader) (HeaderCheckSum uint8, Error error) {
	HeaderCheckSum, Error = Helpers.ReadByte(Reader)
	return HeaderCheckSum, Error
}
func DecodeRecipientAddress(Reader *bytes.Reader) (RecipientAddress uint16, Error error) {
	RecipientAddress, Error = Helpers.ReadUshort(Reader)
	return RecipientAddress, Error
}
func DecodeHeaderEncoding(Reader *bytes.Reader) (HeaderEncoding uint8, Error error) {
	HeaderEncoding, Error = Helpers.ReadByte(Reader)
	return HeaderEncoding, Error

}
func DecodeFrameDataLength(Reader *bytes.Reader) (FrameDataLength uint16, Error error) {
	FrameDataLength, Error = Helpers.ReadUshort(Reader)
	return FrameDataLength, Error
}
func DecodePacketType(Reader *bytes.Reader) (PacketType uint8, Error error) {
	PacketType, Error = Helpers.ReadByte(Reader)
	return PacketType, Error
}

/*Декодирование версии протокола*/
func DecodeProtocolVersion(Reader *bytes.Reader) (ProtocolVersion uint8, Error error) {
	ProtocolVersion, Error = Helpers.ReadByte(Reader)
	return ProtocolVersion, Error
}
func DecodePacketIdentifier(Reader *bytes.Reader) (PacketIdentifier uint16, Error error) {
	PacketIdentifier, Error = Helpers.ReadUshort(Reader)
	return PacketIdentifier, Error
}

func DecodeSecurityKeyId(Reader *bytes.Reader) (SecurityKeyId uint8, Error error) {
	SecurityKeyId, Error = Helpers.ReadByte(Reader)
	return SecurityKeyId, Error

}

func DecodeHeaderLength(Reader *bytes.Reader) (HeaderLength uint8, Error error) {
	HeaderLength, Error = Helpers.ReadByte(Reader)
	return HeaderLength, Error
}

func DecodePeerAddress(Reader *bytes.Reader) (PeerAddress uint16, Error error) {
	PeerAddress, Error = Helpers.ReadUshort(Reader)
	return PeerAddress, Error

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
