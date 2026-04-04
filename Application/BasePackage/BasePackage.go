package BasePackage

import "io"

type BasePackage struct {
}

func (BasePackage *BasePackage) ReadNext(Reader io.Reader, Count int) (Buffer []byte, Error error) {
	Buffer = make([]byte, Count)
	_, Error = io.ReadFull(Reader, Buffer)
	return Buffer, Error
}
