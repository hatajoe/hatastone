package apps

type Reader interface {
	Read() ([]byte, error)
}

type Writer interface {
	Write([]byte) error
}
