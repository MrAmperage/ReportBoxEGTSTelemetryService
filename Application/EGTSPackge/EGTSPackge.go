package EGTSPackge

import "bytes"

/*Структура для ЕГТС пакета*/
type EGTSPackge struct {
	Header EGTSPackgeHeader /*Заголовок пакета*/
}

/*Декодер пакета*/
func (EGTSPackge *EGTSPackge) Decode(Data []byte) (Error error) {
	Reader := bytes.NewReader(Data)
	Error = EGTSPackge.Header.Decode(Reader)
	if Error != nil {
		return Error
	}
	return Error
}
