package stream

func Map[T, U any](stream *Receiver[T], f func(T) U) *Receiver[U] {
	us := &Stream[U]{ch: make(chan U, 1)}
	go func(source *Receiver[T], dest *Sender[U], fn func(T) U) {
		for {
			t, ok := source.Receive()
			if !ok {
				dest.Close()
				return
			}
			dest.Send(fn(t))
		}
	}(stream, us.Sender(), f)
	return us.Receiver()
}
