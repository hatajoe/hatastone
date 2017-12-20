package match

type InputStream struct {
	id     string
	stream chan []string
}

func (s InputStream) GetStream() chan []string {
	return s.stream
}

func (s InputStream) CloseStream() {
	close(s.stream)
}

type InputStreams []*InputStream

func (s InputStreams) FindByID(id string) *InputStream {
	for _, stream := range s {
		if stream.id == id {
			return stream
		}
	}
	return nil
}
