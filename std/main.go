package main

// Sink is the interface of a data sink
type Sink interface {
	Write(event string) error
	Close()
}

type baseSink struct {
}

func (s *baseSink) Write(event string) error {
	// DO SINK
	return nil
}

func (s *baseSink) Close() {
	// DO CLOSE
}

func main() {
	sink := baseSink{}
	sink.Write("test")
	sink.Close()
}
