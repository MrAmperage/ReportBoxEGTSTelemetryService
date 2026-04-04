package EGTSPackge

import "bytes"

/*Структура для ЕГТС пакета*/
type EGTSPackge struct {
	Header EGTSPackgeHeader /*Заголовок пакета*/
}

/*Декодер пакета*/
func (EGTSPackge *EGTSPackge) Decode(Data []byte) {
	Reader := bytes.NewReader(Data)
	EGTSPackge.Header.Decode(Reader)
}
