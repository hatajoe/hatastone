package console

import "log"

type Writer struct{}

func (w Writer) Write(buf []byte) error {
	log.Println(string(buf))
	return nil
}
