package EGTSPackge

/*Структура для ЕГТС пакета*/
type EGTSPackge struct {
	Header EGTSPackgeHeader /*Заголовок пакета*/
}

/*Декодер пакета*/
func (EGTSPackge *EGTSPackge) Decode(Data []byte) (Error error) {

	Error = EGTSPackge.Header.Decode(Data)
	if Error != nil {
		return Error
	}
	return Error
}
