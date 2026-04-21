package EGTSPackge

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
)

/*Структура для ЕГТС пакета*/
type EGTSPackge struct {
	Header  EGTSPackgeHeader /*Заголовок пакета*/
	Records []Record         /*Записи с данными*/

}

/*Декодер пакета*/
func (EGTSPackge *EGTSPackge) Decode(Reader *bytes.Reader) (Error error) {

	Error = EGTSPackge.Header.Decode(Reader)
	if Error != nil {
		return Error
	}
	DecodeRecords(Reader)
	return Error
}
func DecodeRecords(Reader *bytes.Reader) {}

/*Считать одно сообщение*/
func (EGTSPackge *EGTSPackge) ReadMessage(Connection net.Conn) (Buffer []byte, Error error) {

	MinimalHeader := make([]byte, 11)
	_, Error = io.ReadFull(Connection, MinimalHeader)
	if Error != nil {
		return nil, Error
	}
	HeaderLength := int(MinimalHeader[3])
	DataLength := int(binary.LittleEndian.Uint16(MinimalHeader[5:7]))
	TotalLength := HeaderLength + DataLength + 2
	Buffer = make([]byte, TotalLength)
	copy(Buffer, MinimalHeader)
	_, Error = io.ReadFull(Connection, Buffer[11:])
	if Error != nil {
		return nil, Error
	}
	return Buffer, Error
}
