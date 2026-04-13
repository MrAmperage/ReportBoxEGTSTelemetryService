package EGTSPackge

import "bytes"

/*Структура для ЕГТС пакета*/
type EGTSPackge struct {
	Header EGTSPackgeHeader /*Заголовок пакета*/
}

/*Декодер пакета*/
func (EGTSPackge *EGTSPackge) Decode(Reader *bytes.Reader) (Error error) {

	Error = EGTSPackge.Header.Decode(Reader)
	if Error != nil {
		return Error
	}
	return Error
}
