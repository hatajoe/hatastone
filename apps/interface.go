package apps

type Controller interface {
	Read() (interface{}, error)
	Write(interface{}) error
}
